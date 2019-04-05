package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davegarred/wh3/dto"
	"github.com/davegarred/wh3/persist"
	"log"
	"sort"
)

func HandleRequest(_ context.Context, _ events.APIGatewayProxyRequest) (*dto.Response, error) {
	events,err := persist.Search()
	if err != nil {
		return nil,err
	}

	hashEvents := make([]*dto.HashEventV2,0,len(events))
	for _,event := range events {
		hashEvent,err := dto.ConvertGoogleCal(event)
		if err != nil {
			log.Printf("error converting event '%s' from google calendar: %v", event.Id, err)
			continue
		}
		hashEvents = append(hashEvents, hashEvent)
	}
	sort.Slice(hashEvents, func(i int, j int) bool {
		return hashEvents[i].Date < hashEvents[j].Date
	})

	return &dto.Response{"",hashEvents}, nil
}


func main() {
	lambda.Start(HandleRequest)
}
