package dto

import (
	"errors"
	"strings"
	"time"
)

func ConvertGoogleCalForHamster(cal *GoogleCalendar) (*HashEvent, error) {
	if cal.Date == "" && cal.DateTime == "" {
		return nil, errors.New("no acceptable date found")
	}
	date := cal.Date
	if date == "" {
		rfcTime, err := time.Parse(time.RFC3339, cal.DateTime)
		if err != nil {
			return nil, errors.New("no acceptable date found")
		}
		date = rfcTime.Format("2006-01-02")
	}
	eventName,eventNumber := parseEventName_hamster(cal.Summary);
	description := strings.Replace(cal.Description,"’","'",-1)
	hare := ""

	event := &HashEvent{
		GoogleId:    cal.Id,
		Date:        date,
		DateTime:    cal.DateTime,
		EventNumber: eventNumber,
		Hare:        hare,
		EventName:   eventName,
		Description: description,
		MapLink:     cal.EventLocation(),
		Kennel:      HAMSTER,
	}

	return event, nil
}

func parseEventName_hamster(summary string) (string,string) {
	return summary,""
}