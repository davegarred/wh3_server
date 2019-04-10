package dto

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ConvertGoogleCal(cal *GoogleCalendar) (*HashEvent, error) {
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
	eventName := parseEventName(cal.Summary);
	kennel := guessKennel(cal.Summary);
	eventNumber := guessEventNumber(cal.Summary)
	hare := parseHare(cal.Summary)

	event := &HashEvent{
		GoogleId:    cal.Id,
		Date:        date,
		DateTime:    cal.DateTime,
		EventNumber: eventNumber,
		Hare:        hare,
		EventName:   eventName,
		Description: cal.Description,
		MapLink:     cal.EventLocation(),
		Kennel:      kennel,
	}

	return event, nil
}

const (
	TBD                       = "TBD"
	harePattern               = "((?i)Hare\\(?s?\\)?): .*"
	ANALVERSARY_TITLE_PATTERN = "\\([0-9]{4}\\)"
	unkownTitlePattern        = "\\(TBD\\)"
	titlePattern              = "\\(.*\\)"
)

var (
	unknownTitleRegex = regexp.MustCompile(unkownTitlePattern)
	titleRegex        = regexp.MustCompile(titlePattern)
	hareRegex         = regexp.MustCompile(harePattern)
)

func parseEventName(summary string) string {
	kennel := guessKennel(summary)
	if kennel != UNKNOWN {
		//check for anniversary
	}

	unknownTitle := unknownTitleRegex.MatchString(summary)
	if unknownTitle {
		pos := unknownTitleRegex.FindStringIndex(summary)[0]
		return summary[:pos-1]
	}
	matchesTitlePattern := titleRegex.MatchString(summary)
	if matchesTitlePattern {
		enclosedTitle := titleRegex.FindString(summary)
		return enclosedTitle[1 : len(enclosedTitle)-1]
	}
	return summary
}

func guessKennel(summary string) KennelId {
	kennels := map[string]KennelId{
		"Seattle H3 Hashy Hour":      HAPPY_HOUR,
		"SH3":                        SEATTLE,
		"SeaMon":                     SEAMON,
		"NBH3":                       NO_BALLS,
		"TH3":                        TACOMA,
		"SSH3":                       SOUTH_SOUND,
		"S3H3":                       SS_SHITSHOW,
		"PSH3":                       PUGET_SOUND,
		"RCH3":                       RAIN_CITY,
		"HSWTF":                      HSWTF,
		"Thursday Renton Happy Hour": RENTON_HAPPY_HOUR,
		"South End H3 Happy Hour":    RENTON_HAPPY_HOUR,
		"FMH3":                       FULL_MOON,
		"SH2B":                       BASH,
	}

	for identifier, kennel := range kennels {
		if strings.HasPrefix(summary, identifier) {
			return kennel
		}
	}
	return UNKNOWN
}

func guessEventNumber(summary string) string {
	pos := strings.Index(summary, "#")
	if pos > 0 {
		substr := summary[pos+1:]
		endPos := strings.Index(substr, " ")
		if endPos > 0 {
			eventNumber := substr[:endPos]
			_, err := strconv.Atoi(eventNumber)
			if err == nil {
				return eventNumber
			}
		}
	}
	return ""
}

func parseHare(s string) string {
	return hareRegex.FindString(s)
}
