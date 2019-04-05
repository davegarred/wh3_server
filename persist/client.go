package persist

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davegarred/wh3/dto"
	"time"
)

const (
	dynamoTable = "wh3_google_calendar"
	primaryKey  = "googleId"
	dateIndex   = "eventDate"
	payload     = "payload"
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
		TableName: aws.String(dynamoTable),
	})
	if err != nil {
		return nil, err
	}

	event := result.Item[payload].S
	fmt.Println(*event)

	return event, nil
}

func Search() ([]*dto.GoogleCalendar, error) {
	svc, err := client()
	if err != nil {
		return nil, err
	}

	scanOutput, err := svc.Scan(eventsAfterToday())
	if err != nil {
		return nil, err
	}

	result := make([]*dto.GoogleCalendar, 0, len(scanOutput.Items))
	for _, item := range scanOutput.Items {
		serEvent := item[payload].S
		event := &dto.GoogleCalendar{}
		err := json.Unmarshal([]byte(*serEvent), event)
		if err != nil {
			return nil, err
		}
		result = append(result, event)
	}

	return result, nil
}

func eventsAfterToday() *dynamodb.ScanInput {
	start := time.Now()
	return &dynamodb.ScanInput{
		ExpressionAttributeNames: map[string]*string{
			"#d": aws.String(dateIndex),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":start": {
				S: aws.String(start.Format("2006-01-02")),
			},
		},
		FilterExpression:     aws.String("#d >= :start"),
		ProjectionExpression: aws.String(payload),
		TableName:            aws.String(dynamoTable),
	}
}

func Put(events []*dto.GoogleCalendar) error {
	svc, err := client()
	if err != nil {
		return err
	}

	for _, event := range events {
		ser, _ := json.Marshal(event)
		_, err := svc.PutItem(&dynamodb.PutItemInput{
			ExpressionAttributeNames:  nil,
			ExpressionAttributeValues: nil,
			Item: map[string]*dynamodb.AttributeValue{
				primaryKey: &dynamodb.AttributeValue{
					S: aws.String(event.Id),
				},
				dateIndex: &dynamodb.AttributeValue{
					S: aws.String(event.EventDate()),
				},
				payload: &dynamodb.AttributeValue{
					S: aws.String(string(ser)),
				}},
			TableName: aws.String(dynamoTable),
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
