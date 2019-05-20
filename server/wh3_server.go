package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davegarred/wh3/dto"
	"github.com/davegarred/wh3/persist"
)

func HandleRequest(_ context.Context, _ events.APIGatewayProxyRequest) (*dto.Response, error) {
	wh3Events, hswtfEvents, hamsterEvents, err := persist.AllCalendarEvents()
	if err != nil {
		return nil, err
	}
	calendarEvents := dto.ConvertCalendarEvents(wh3Events, hswtfEvents, hamsterEvents)

	adminEvents, err := persist.AllAdminEvents()
	if err != nil {
		return nil, err
	}

	kennels, err := persist.AllKennels()
	if err != nil {
		return nil, err
	}

	return dto.ProcessAndWrap(calendarEvents, adminEvents, kennels), nil
}

func main() {
	lambda.Start(HandleRequest)
}
