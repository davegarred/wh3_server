package persist

import (
	"fmt"
	"github.com/davegarred/wh3/dto"
	"testing"
)

func _TestGet(t *testing.T) {
	client := NewDynamoClient()
	data,err := client.GetKennel("PUGET_SOUND")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func _TestSearch(t *testing.T) {
	client := NewDynamoClient()
	wh3Events, _, _,err := client.AllCalendarEvents()
	if err != nil {
		panic(err)
	}
	for _,event := range wh3Events {
		fmt.Println(event.Date)
	}
}

func _TestPut(t *testing.T) {
	client := NewDynamoClient()
	err := client.Put("test", []*dto.GoogleCalendar{
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
