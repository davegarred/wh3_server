package dto

import (
	"log"
	"sort"
)

type KennelId string

var (
	SEAMON      = KennelId("SEAMON")
	SEATTLE     = KennelId("SEATTLE")
	PUGET_SOUND = KennelId("PUGET_SOUND")
	NO_BALLS    = KennelId("NO_BALLS")
	TACOMA      = KennelId("TACOMA")
	SOUTH_SOUND = KennelId("SOUTH_SOUND")
	SS_SHITSHOW = KennelId("SS_SHITSHOW")
	RAIN_CITY   = KennelId("RAIN_CITY")
	HSWTF       = KennelId("HSWTF")
	HAMSTER     = KennelId("HAMSTER")
	GIGGITY     = KennelId("GIGGITY")
	HANK        = KennelId("HANK")

	RENTON_HAPPY_HOUR = KennelId("RENTON_HAPPY_HOUR")
	FULL_MOON         = KennelId("FULL_MOON")
	BASH              = KennelId("BASH")
	HAPPY_HOUR        = KennelId("HAPPY_HOUR")
	UNKNOWN           = KennelId("UNKNOWN")
)

type Response struct {
	Message string       `json:"message,omitempty"`
	Events  []*HashEvent `json:"events"`
	Kennels []*Kennel    `json:"kennels"`
}

func ConvertCalendarEvents(wh3Events []*GoogleCalendar, hswtfEvents []*GoogleCalendar, hamsterEvents []*GoogleCalendar) map[string]*HashEvent {
	hashEventMap := make(map[string]*HashEvent)
	for _, event := range wh3Events {
		hashEvent, err := ConvertGoogleCal(event)
		if err != nil {
			log.Printf("error converting event '%s' from google calendar: %v", event.Id, err)
			continue
		}
		hashEventMap[hashEvent.GoogleId] = hashEvent
	}

	for _, event := range hswtfEvents {
		hashEvent, err := ConvertGoogleCalForHSWTF(event)
		if err != nil {
			log.Printf("error converting HSWTF event '%s' from google calendar: %v", event.Id, err)
			continue
		}
		hashEventMap[hashEvent.GoogleId] = hashEvent
	}

	for _, event := range hamsterEvents {
		hashEvent, err := ConvertGoogleCalForHamster(event)
		if err != nil {
			log.Printf("error converting hamster event '%s' from google calendar: %v", event.Id, err)
			continue
		}
		hashEventMap[hashEvent.GoogleId] = hashEvent
	}
	return hashEventMap
}

func ProcessAndWrap(calendarEvents map[string]*HashEvent, adminEvents map[string]*HashEvent, kennels []*Kennel) *Response {
	for _, e := range adminEvents {
		event := calendarEvents[e.GoogleId]
		if event == nil {
			continue
		}
		if len(e.DateTime) > 0 {
			event.DateTime = e.DateTime
		}
		if len(e.EventName) > 0 {
			event.EventName = e.EventName
		}
		if len(e.EventNumber) > 0 {
			event.EventNumber = e.EventNumber
		}
		if len(e.Description) > 0 {
			event.Description = e.Description
		}
		if len(e.MapLink) > 0 {
			event.MapLink = e.MapLink
		}
		if len(e.Kennel) > 0 {
			event.Kennel = e.Kennel
		}
		if len(e.Hare) > 0 {
			event.Hare = e.Hare
		}
		calendarEvents[event.GoogleId] = event
	}

	sortedEvents := make([]*HashEvent, 0, len(calendarEvents))
	for _, e := range calendarEvents {
		sortedEvents = append(sortedEvents, e)
	}
	sort.Slice(sortedEvents, func(i int, j int) bool {
		if (sortedEvents[i].Date != sortedEvents[j].Date) {
			return sortedEvents[i].Date < sortedEvents[j].Date
		}
		if (sortedEvents[i].DateTime != sortedEvents[j].DateTime) {
			return sortedEvents[i].DateTime < sortedEvents[j].DateTime
		}
		return sortedEvents[i].EventName < sortedEvents[j].EventName
	})
	return &Response{"", sortedEvents, kennels}
}

type HashEvent struct {
	GoogleId    string   `json:"googleId"`
	Date        string   `json:"date"`
	DateTime    string   `json:"dateTime,omitempty"`
	EventNumber string   `json:"eventNumber,omitempty"`
	Hare        string   `json:"hare,omitempty"`
	EventName   string   `json:"eventName"`
	Description string   `json:"description"`
	MapLink     string   `json:"mapLink,omitempty"`
	Kennel      KennelId `json:"kennel"`
}

type Kennel struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	AlternativeImage string   `json:"alternativeImage,omitempty"`
	HareraiserName   string   `json:"hareraiserName"`
	HareraiserEmail  string   `json:"hareraiserEmail"`
	Badges           []string `json:"badges"`
	FirstHash        string   `json:"firstHash"`
	Founders         string   `json:"founders"`
	Lineage          string   `json:"lineage"`
}
