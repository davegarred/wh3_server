package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davegarred/wh3/dto"
	"github.com/davegarred/wh3/persist"
)

type RequestHandler struct {
	db persist.Persist
}
func (h *RequestHandler) HandleRequest(_ context.Context, _ events.APIGatewayProxyRequest) (*dto.Response, error) {
	wh3Events, hswtfEvents, hamsterEvents, err := h.db.AllCalendarEvents()
	if err != nil {
		return nil, err
	}
	calendarEvents := dto.ConvertCalendarEvents(wh3Events, hswtfEvents, hamsterEvents)

	adminEvents, err := h.db.AllAdminEvents()
	if err != nil {
		return nil, err
	}

	kennels, err := h.db.AllKennels()
	if err != nil {
		return nil, err
	}

	return dto.ProcessAndWrap(calendarEvents, adminEvents, kennels), nil
}

func main() {
	handler := &RequestHandler{persist.NewDynamoClient()}
	lambda.Start(handler)
}
