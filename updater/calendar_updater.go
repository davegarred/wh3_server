package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davegarred/wh3/dto"
	"github.com/davegarred/wh3/persist"
	"github.com/davegarred/wh3/reader"
	"io/ioutil"
	"log"
)

func HandleRequest(_ context.Context, _ interface{}) (interface{}, error) {
	hashEvents, err := pullEvents()
	if err != nil {
		return nil, err
	}

	return nil, persistEvents(hashEvents)
}

func pullEvents() ([]*dto.GoogleCalendar, error) {
	dat, err := ioutil.ReadFile("wh3-calendar-cb8bb1a84750.json")
	if err != nil {
		return nil, err
	}

	client, err := reader.NewCalendarReader(dat)
	if err != nil {
		return nil, err
	}

	hashEvents, err := client.Events()
	if err != nil {
		return nil, err
	}

	calendarEvents := make([]*dto.GoogleCalendar, 0, len(hashEvents.Items))
	for _, hashEvent := range hashEvents.Items {
		calendarEvents = append(calendarEvents, &dto.GoogleCalendar{
			Id:          hashEvent.Id,
			Date:        hashEvent.Start.Date,
			DateTime:    hashEvent.Start.DateTime,
			Summary:     hashEvent.Summary,
			Location:    hashEvent.Location,
			Description: hashEvent.Description,
		})
	}
	return calendarEvents, err
}

func persistEvents(events []*dto.GoogleCalendar) error {
	log.Printf("found and persisting %d events", len(events))
	return persist.Put(events)
}

func main() {
	lambda.Start(HandleRequest)
}
