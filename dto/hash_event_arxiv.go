package dto

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

type Response_Archive struct {
	Message *string         `json:"message"`
	Events  []*HashEventV2 `json:"events"`
}

func ConvertToV2(events []*GoogleCalendar) *Response_Archive {
	hashEvents := make([]*HashEventV2, 0, len(events))
	for _, event := range events {
		hashEvent, err := ConvertGoogleCal(event)
		if err != nil {
			log.Printf("error converting event '%s' from google calendar: %v", event.Id, err)
			continue
		}
		hashEventV2, err := ConvertGoogleCal_Archiv(hashEvent)
		if err != nil {
			log.Printf("error converting event '%s' from standard event to archive event: %v", event.Id, err)
			continue
		}
		hashEvents = append(hashEvents, hashEventV2)
	}
	sort.Slice(hashEvents, func(i int, j int) bool {
		if hashEvents[i].Date[1] != hashEvents[j].Date[1] {
			return hashEvents[i].Date[1] < hashEvents[j].Date[1]
		}
		return hashEvents[i].Date[2] < hashEvents[j].Date[2]
	})
	return &Response_Archive{nil, hashEvents}
}

type HashEventV2 struct {
	GoogleId    string `json:"id"`
	Date        []int  `json:"date"`
	EventNumber int    `json:"eventNumber"`
	Hare        *string `json:"hare"`
	EventName   string `json:"eventName"`
	Description string `json:"description"`
	Address     *string `json:"address"`
	MapLink     *string `json:"mapLink"`
	Kennel      string `json:"kennel"`
}

func ConvertGoogleCal_Archiv(ev *HashEvent) (*HashEventV2, error) {
	if ev.Date == "" && ev.DateTime == "" {
		return nil, errors.New("no acceptable date found")
	}
	var date []int
	if ev.DateTime != "" {
		rfcTime, err := time.Parse(time.RFC3339, ev.DateTime)
		if err != nil {
			msg := fmt.Sprintf("dateTime '%s' field could not be parsed: %v", ev.DateTime, err)
			return nil, errors.New(msg)
		}
		date = []int{rfcTime.Year(), int(rfcTime.Month()), rfcTime.Day(), rfcTime.Hour(), rfcTime.Minute()}
	} else {
		rfcTime, err := time.Parse("2006-01-02", ev.Date)
		if err != nil {
			msg := fmt.Sprintf("date '%s' field could not be parsed: %v", ev.Date, err)
			return nil, errors.New(msg)
		}
		date = []int{rfcTime.Year(), int(rfcTime.Month()), rfcTime.Day(),0,0}
	}

	eventNumber, err := strconv.Atoi(ev.EventNumber)
	if err != nil {
		eventNumber = 0
	}
	var mapLink *string
	//if ev.Address != "" {
	//	target := "https://maps.google.com/maps?q=" + url.PathEscape(ev.Address)
	//	mapLink = &target
	//}

	hare := &ev.Hare
	if *hare == "" {
		hare = nil
	}
	event := &HashEventV2{
		GoogleId:    ev.GoogleId,
		Date:        date,
		EventNumber: eventNumber,
		Hare:        hare,
		EventName:   ev.EventName,
		Description: ev.Description,
		Address:     nil,
		MapLink:     mapLink,
		Kennel:      "",
	}

	return event, nil
}
