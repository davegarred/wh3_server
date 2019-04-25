package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davegarred/wh3/dto"
	"github.com/davegarred/wh3/persist"
	"github.com/davegarred/wh3/reader"
	"google.golang.org/api/calendar/v3"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func HandleRequest(_ context.Context, _ interface{}) (interface{}, error) {
	err := updateEvents()
	if err != nil {
		log.Printf("Error encountered %v", err)
	}
	return nil, err
}

func updateEvents() error {
	wh3Events, hswtfEvents, err := pullWH3Events()
	if err != nil {
		return err
	}

	err = persistEvents("wh3", wh3Events)
	if err != nil {
		return err
	}
	return persistEvents("hswtf", hswtfEvents)
}

func formatGoogleId(event *calendar.Event) string {
	idx := strings.Index(event.Id, "_")
	if idx < 0 || event.Start == nil {
		return event.Id
	}
	if len(event.Start.DateTime) > 0{
		truncatedId := event.Id[:idx + 1]
		t,err := time.Parse(time.RFC3339, event.Start.DateTime)
		if err == nil {
			_,offset := time.Now().Zone()
			t.In(time.UTC).Add(time.Duration(offset) * time.Hour)
			formattedDate := t.Format("20060102")
			return truncatedId + formattedDate
		}
	}
	return event.Id;
}
func pullWH3Events() ([]*dto.GoogleCalendar, []*dto.GoogleCalendar, error) {
	dat, err := ioutil.ReadFile("wh3-calendar-cb8bb1a84750.json")
	if err != nil {
		return nil, nil, err
	}

	client, err := reader.NewCalendarReader(dat)
	if err != nil {
		return nil, nil, err
	}

	wh3HashEvents, err := client.WH3Events()
	if err != nil {
		return nil, nil, err
	}

	wh3CalendarEvents := make([]*dto.GoogleCalendar, 0, len(wh3HashEvents.Items))
	for _, hashEvent := range wh3HashEvents.Items {
		formattedId := formatGoogleId(hashEvent)
		wh3CalendarEvents = append(wh3CalendarEvents, &dto.GoogleCalendar{
			Id:          formattedId,
			Date:        hashEvent.Start.Date,
			DateTime:    hashEvent.Start.DateTime,
			Summary:     hashEvent.Summary,
			Location:    hashEvent.Location,
			Description: hashEvent.Description,
		})
	}


	hswtfHashEvents, err := client.HSWTFEvents()
	if err != nil {
		return nil, nil, err
	}

	hswtfCalendarEvents := make([]*dto.GoogleCalendar, 0, len(hswtfHashEvents.Items))
	for _, hashEvent := range hswtfHashEvents.Items {
		formattedId := formatGoogleId(hashEvent)
		hswtfCalendarEvents = append(hswtfCalendarEvents, &dto.GoogleCalendar{
			Id:          formattedId,
			Date:        hashEvent.Start.Date,
			DateTime:    hashEvent.Start.DateTime,
			Summary:     hashEvent.Summary,
			Location:    hashEvent.Location,
			Description: hashEvent.Description,
		})
	}
	return wh3CalendarEvents, hswtfCalendarEvents, err
}

func persistEvents(calendar string, events []*dto.GoogleCalendar) error {
	log.Printf("found and persisting %d events", len(events))
	return persist.Put(calendar, events)
}

func main() {
	lambda.Start(HandleRequest)
}
