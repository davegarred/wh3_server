package main

import (
	"context"
	"github.com/davegarred/wh3/persist"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/calendar/v3"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	handler := &UpdateHandler{persist.NewDynamoClient()}
	_, err := handler.HandleRequest(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}

func TestFormatGoogleId(t *testing.T) {
	assert.Equal(t, "6tqvpheanlsel1siabqa8m5ljt", formatGoogleId(googleEvent("6tqvpheanlsel1siabqa8m5ljt", "2019-05-04", "")))
	assert.Equal(t, "t0qm960i1gh7f7uuvf90uusggo_20190429", formatGoogleId(&calendar.Event{
		Id: "t0qm960i1gh7f7uuvf90uusggo_20190429",
		Start: &calendar.EventDateTime{
			Date:     "2019-04-29",
			DateTime: "",
		},
	}))
	formatted := formatGoogleId(googleEvent("0odi5geepodms3144rv27pr7hp_20190505", "", "2019-05-05T12:00:00-07:00"))
	assert.Equal(t, "0odi5geepodms3144rv27pr7hp_20190505", formatted)

	assert.Equal(t, "0odi5geepodms3144rv27pr7hp_20190505T190000Z", formatGoogleId(&calendar.Event{
		Id:    "0odi5geepodms3144rv27pr7hp_20190505T190000Z",
		Start: nil,
	}))
	formatted = formatGoogleId(googleEvent("sngrdtiil6asch38f5sb9hoquo_20190506T005500Z", "", "2019-05-05T17:55:00-07:00"))
	assert.Equal(t, "sngrdtiil6asch38f5sb9hoquo_20190505", formatted)
}

func googleEvent(id string, date string, dateTime string) *calendar.Event {
	return &calendar.Event{
		Id: id,
		Start: &calendar.EventDateTime{
			Date:     date,
			DateTime: dateTime,
		},
	}
}
