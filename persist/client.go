package persist

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davegarred/wh3/dto"
	"log"
	"time"
)

const (
	eventTable      = "wh3_google_calendar"
	adminEventTable = "wh3_admin_event"
	kennelTable     = "wh3_kennel"
	primaryKey      = "googleId"
	dateIndex       = "eventDate"
	calendarField   = "calendar"
	payload         = "payload"
)

func Get(key string) (*string, error) {
	svc, err := client()
	if err != nil {
		return nil, err
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(key),
			},
		},
		TableName: aws.String(kennelTable),
	})
	if err != nil {
		return nil, err
	}

	event := result.Item[payload].S
	fmt.Println(*event)

	return event, nil
}

func AllCalendarEvents() ([]*dto.GoogleCalendar, []*dto.GoogleCalendar, error) {
	svc, err := client()
	if err != nil {
		return nil, nil, err
	}

	scanOutput, err := svc.Scan(googleCalendarEventsAfterToday())
	if err != nil {
		return nil, nil, err
	}

	wh3Events := make([]*dto.GoogleCalendar, 0, len(scanOutput.Items))
	hswtfEvents := make([]*dto.GoogleCalendar, 0, len(scanOutput.Items))
	for _, item := range scanOutput.Items {
		calendar := item[calendarField].S
		serEvent := item[payload].S
		event := &dto.GoogleCalendar{}
		err := json.Unmarshal([]byte(*serEvent), event)
		if err != nil {
			return nil, nil, err
		}
		if "wh3" == *calendar {
			wh3Events = append(wh3Events, event)
		} else if "hswtf" == *calendar {
			hswtfEvents = append(hswtfEvents, event)
		}
	}

	return wh3Events, hswtfEvents, nil
}

func AllAdminEvents() (map[string]*dto.HashEvent, error) {
	svc, err := client()
	if err != nil {
		return nil, err
	}

	scanOutput, err := svc.Scan(adminEventsAfterToday())
	if err != nil {
		return nil, err
	}

	adminEvents := make(map[string]*dto.HashEvent)
	for _, item := range scanOutput.Items {
		adminEvent := &dto.HashEvent{}
		googleId := item[primaryKey].S
		serializedEvent := item[payload].S
		err = json.Unmarshal([]byte(*serializedEvent), adminEvent)
		if err != nil {
			log.Printf("error deserializing admin event with ID %s - %v", *googleId, err)
		} else {
			adminEvents[adminEvent.GoogleId] = adminEvent
		}
	}
	return adminEvents, nil
}

func AllKennels() ([]*dto.Kennel, error) {
	svc, err := client()
	if err != nil {
		return nil, err
	}

	scanOutput, err := svc.Scan(allKennels())
	if err != nil {
		return nil, err
	}

	result := make([]*dto.Kennel, 0, len(scanOutput.Items))
	for _, item := range scanOutput.Items {
		serEvent := item[payload].S
		event := &dto.Kennel{}
		err := json.Unmarshal([]byte(*serEvent), event)
		if err != nil {
			return nil, err
		}
		result = append(result, event)
	}

	return result, nil
}

func adminEventsAfterToday() *dynamodb.ScanInput {
	return &dynamodb.ScanInput{
		ExpressionAttributeNames: map[string]*string{
			"#d": aws.String(dateIndex),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":start": {
				S: aws.String(today()),
			},
		},
		FilterExpression: aws.String("#d >= :start"),
		//ProjectionExpression: aws.String(payload),
		TableName: aws.String(adminEventTable),
	}
}

func googleCalendarEventsAfterToday() *dynamodb.ScanInput {
	return &dynamodb.ScanInput{
		ExpressionAttributeNames: map[string]*string{
			"#d": aws.String(dateIndex),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":start": {
				S: aws.String(today()),
			},
		},
		FilterExpression: aws.String("#d >= :start"),
		TableName:        aws.String(eventTable),
	}
}

func today() string {
	return time.Now().UTC().Add(time.Duration(-7) * time.Hour).Format("2006-01-02")
}

func allKennels() *dynamodb.ScanInput {
	return &dynamodb.ScanInput{
		ProjectionExpression: aws.String(payload),
		TableName:            aws.String(kennelTable),
	}
}

func Put(calendar string, events []*dto.GoogleCalendar) error {
	svc, err := client()
	if err != nil {
		return err
	}

	trackingSet := make(map[string]bool)
	for _, event := range events {
		if trackingSet[event.Id] {
			log.Printf("found duplicate events for ID: " + event.Id)
		}
		ser, _ := json.Marshal(event)
		_, err := svc.PutItem(&dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				primaryKey: {
					S: aws.String(event.Id),
				},
				dateIndex: {
					S: aws.String(event.EventDate()),
				},
				calendarField: {
					S: aws.String(calendar),
				},
				payload: {
					S: aws.String(string(ser)),
				}},
			TableName: aws.String(eventTable),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func client() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}
