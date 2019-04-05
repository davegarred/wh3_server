package persist

import (
	"fmt"
	"github.com/davegarred/wh3/dto"
	"testing"
)

func TestGet(t *testing.T) {
	data,err := Get("0lifalmd3dh70cnibsg90ha2or")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func TestSearch(t *testing.T) {
	events,err := Search()
	if err != nil {
		panic(err)
	}
	for _,event := range events {
		fmt.Println(event.Date)
	}
}

func TestPut(t *testing.T) {
	err := Put([]*dto.GoogleCalendar{
		{
			Id:          "test-id",
			Date:        "2019-04-05",
			DateTime:    "",
			Summary:     "",
			Location:    "",
			Description: "",
		},
	})
	if err != nil {
		panic(err)
	}
}
