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

func ConvertAndWrap(wh3Events []*GoogleCalendar, hswtfEvents []*GoogleCalendar, kennels []*Kennel) *Response {
	hashEvents := make([]*HashEvent, 0, len(wh3Events))
	for _, event := range wh3Events {
		hashEvent, err := ConvertGoogleCal(event)
		if err != nil {
			log.Printf("error converting event '%s' from google calendar: %v", event.Id, err)
			continue
		}
		hashEvents = append(hashEvents, hashEvent)
	}
	for _, event := range hswtfEvents {
		hashEvent, err := ConvertGoogleCalForHSWTF(event)
		if err != nil {
			log.Printf("error converting HSWTF event '%s' from google calendar: %v", event.Id, err)
			continue
		}
		hashEvents = append(hashEvents, hashEvent)
	}
	sort.Slice(hashEvents, func(i int, j int) bool {
		return hashEvents[i].Date < hashEvents[j].Date
	})
	return &Response{"", hashEvents, kennels}
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
