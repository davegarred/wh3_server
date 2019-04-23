package dto

import (
	"log"
	"sort"
)

type KennelId string

var (
	SEAMON      = KennelId("SEAMON")      //, "Seattle Monday Hash", ""}
	SEATTLE     = KennelId("SEATTLE")     //, "Seattle H3", ""}
	PUGET_SOUND = KennelId("PUGET_SOUND") //, "Puget Sound", ""}
	NO_BALLS    = KennelId("NO_BALLS")    //, "No Balls H3", ""}
	TACOMA      = KennelId("TACOMA")      //, "Tacoma H3", ""}
	SOUTH_SOUND = KennelId("SOUTH_SOUND") //, "South Sound", ""}
	SS_SHITSHOW = KennelId("SS_SHITSHOW") //, "South Sound Shitshow", ""}
	RAIN_CITY   = KennelId("RAIN_CITY")   //, "Rain City H3", ""}
	HSWTF       = KennelId("HSWTF")       //, "Holy Shit WTF H3", ""}
	GIGGITY     = KennelId("GIGGITY")     //, "Renton Happy Hour", ""}
	HANK        = KennelId("HANK")        //, "Renton Happy Hour", ""}

	RENTON_HAPPY_HOUR = KennelId("RENTON_HAPPY_HOUR") //, "Renton Happy Hour", ""}
	FULL_MOON         = KennelId("FULL_MOON")         //, "Full Moon", ""}
	BASH              = KennelId("BASH")              //, "Bike Hash", ""}
	HAPPY_HOUR        = KennelId("HAPPY_HOUR")        //, "Hashy Hour", ""}
	UNKNOWN           = KennelId("UNKNOWN")           //, "", ""}
)

type Response struct {
	Message string       `json:"message,omitempty"`
	Events  []*HashEvent `json:"events"`
	Kennels []*Kennel    `json:"kennels"`
}

func ConvertCalendarEvents(wh3Events []*GoogleCalendar, hswtfEvents []*GoogleCalendar) map[string]*HashEvent {
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
	return hashEventMap
}

func ProcessAndWrap(calendarEvents map[string]*HashEvent, adminEvents map[string]*HashEvent, kennels []*Kennel) *Response {
	for _,e := range adminEvents {
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
			event.Description= e.Description
		}
		if len(e.MapLink) > 0 {
			event.MapLink= e.MapLink
		}
		if len(e.Kennel) > 0 {
			event.Kennel= e.Kennel
		}
		if len(e.Hare) > 0 {
			event.Hare = e.Hare
		}
		calendarEvents[event.GoogleId] = event
	}

	sortedEvents := make([]*HashEvent, 0, len(calendarEvents))
	for _,e := range calendarEvents {
		sortedEvents = append(sortedEvents, e)
	}
	sort.Slice(sortedEvents, func(i int, j int) bool {
		return sortedEvents[i].Date < sortedEvents[j].Date
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
