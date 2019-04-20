package reader

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"time"
)

var (
	seaH3Cal = "8d65om7lrdq538ksqednh2n648@group.calendar.google.com"
	hswtfH3Cal = "8d65om7lrdq538ksqednh2n648@group.calendar.google.com"
)
type CalendarReader struct {
	svc *calendar.Service
}

func NewCalendarReader(configFile []byte) (*CalendarReader, error){
	config, err :=google.JWTConfigFromJSON(configFile, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}

	svc, err := calendar.New(config.Client(oauth2.NoContext))
	if err != nil {
		return nil, err
	}
	return &CalendarReader{svc},nil
}
func (r *CalendarReader) WH3Events() (*calendar.Events, error) {
	minTime := time.Now().Format(time.RFC3339)
	maxTime := time.Now().Add(3 * 24 * 30 * time.Hour).Format(time.RFC3339)
	events,err := r.svc.Events.List(seaH3Cal).ShowDeleted(false).TimeMin(minTime).TimeMax(maxTime).SingleEvents(true).OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *CalendarReader) HSWTFEvents() (*calendar.Events, error) {
	minTime := time.Now().Format(time.RFC3339)
	maxTime := time.Now().Add(3 * 24 * 30 * time.Hour).Format(time.RFC3339)
	events,err := r.svc.Events.List("hswtfh3@gmail.com").ShowDeleted(false).TimeMin(minTime).TimeMax(maxTime).SingleEvents(true).OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}

	return events, nil
}