package reader

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

const (
	calendarEmail = "account-1@wh3-calendar.iam.gserviceaccount.com"
	seah3Calendar = "8d65om7lrdq538ksqednh2n648@group.calendar.google.com";
)

func TestCal(t *testing.T) {
	dat, err := ioutil.ReadFile("../wh3-calendar-cb8bb1a84750.json")
	if err != nil {
		panic(err)
	}
	client, err := NewCalendarReader(dat)
	if err != nil {
		panic(err)
	}
	events, err := client.HSWTFEvents()
	//events, err := client.WH3Events()
	if err != nil {
		panic(err)
	}

	for _,event := range events.Items {
		dateTime := event.Start.DateTime
		date := event.Start.Date
		if date == "" {
			dt,_ := time.Parse(time.RFC3339, dateTime)
			date = dt.Format("2006-01-02")
		}
		fmt.Printf("%s %v - %s @@ %s\n", event.Id, date, event.Summary, event.Location)
		//event.Description
	}
	fmt.Printf("total events: %d\n", len(events.Items))
}
