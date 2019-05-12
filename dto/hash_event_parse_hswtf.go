package dto

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

func ConvertGoogleCalForHSWTF(cal *GoogleCalendar) (*HashEvent, error) {
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
	eventName,eventNumber := parseEventName_hswtf(cal.Summary);
	description := strings.Replace(cal.Description,"’","'",-1)
	kennel := HSWTF;
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
		Kennel:      kennel,
	}

	return event, nil
}

const (
	hswtfSummaryRunNumberPattern = "HS!WTF\\?H3 - [0-9]{3}[a-zA-Z]{2}"
	hswtfSummaryRunTitlePattern = "HS!WTF\\?H3 - [0-9]{3}[a-zA-Z]{2} - .+"
	runSuffix = len("HS!WTF\\?H3 - ")
	titleSuffix = len("HS!WTF\\?H3 - XXXTH - ")
)
var (
	hswtfRunPatternRegex = regexp.MustCompile(hswtfSummaryRunNumberPattern)
	hswtfTitlePatternRegex = regexp.MustCompile(hswtfSummaryRunTitlePattern)
)
func parseEventName_hswtf(summary string) (string,string) {
	eventNumber := ""
	if hswtfRunPatternRegex.MatchString(summary){
		eventNumber = summary[runSuffix-1:runSuffix+2]
	}
	eventName := "Holy Shit! WTF? H3"
	if hswtfTitlePatternRegex.MatchString(summary) {
		eventName = summary[titleSuffix-1:]
		eventName = strings.Replace(eventName, "’", "'", -1)
	}
	return eventName,eventNumber
}