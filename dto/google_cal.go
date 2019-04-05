package dto

import "time"

type GoogleCalendar struct {
	Id          string `json:"id"`
	Date        string `json:"date"`
	DateTime    string `json:"dateTime"`
	Summary     string `json:"summary"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

func (c *GoogleCalendar) EventDate() string {
	if c.Date != "" {
		return c.Date
	}
	dateTime, err := time.Parse(time.RFC3339, c.DateTime)
	if err != nil {
		panic(err)
	}
	return dateTime.Format("2006-01-02")
}
