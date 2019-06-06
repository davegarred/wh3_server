package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleCalendar_EventLocation(t *testing.T) {
	cal := &GoogleCalendar{
		Location: "Greenwood Park, 8905 Fremont Ave N, Seattle, WA 98103",
	}
	assert.Equal(t, "https://maps.google.com/maps?q=Greenwood%20Park%2C%208905%20Fremont%20Ave%20N%2C%20Seattle%2C%20WA%2098103", cal.EventLocation())
}

func TestGoogleCalendar_EventLocation_none(t *testing.T) {
	cal := &GoogleCalendar{
		Location: "TBD",
	}
	assert.Equal(t, "", cal.EventLocation())
}

func TestGoogleCalendar_EventDate(t *testing.T) {
	assert.Equal(t, "2019-04-05", testCalendarDates("2019-04-05", "").EventDate())
	assert.Equal(t, "2019-04-05", testCalendarDates("", "2019-04-05T23:59:59.000Z").EventDate())
	assert.Equal(t, "2019-04-06", testCalendarDates("", "2019-04-06T05:45:23.000Z").EventDate())
}
func testCalendarDates(date, dateTime string) *GoogleCalendar {
	return &GoogleCalendar{
		Date:     date,
		DateTime: dateTime,
	}
}
