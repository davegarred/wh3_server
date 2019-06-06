package dto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplyAdminEvents(t *testing.T) {
	actual := testEvent()
	applyAdminEvents(map[string]*HashEvent{"test": {
		GoogleId:    "",
		Date:        "",
		DateTime:    "",
		EventNumber: "",
		Hare:        "",
		EventName:   "",
		Description: "",
		MapLink:     "",
		Kennel:      "",
	}}, map[string]*HashEvent{"test": actual})
	expected := testEvent()
	assert.Equal(t, expected, actual)
}

func testEvent() *HashEvent {
	return &HashEvent{
		GoogleId:    "test",
		Date:        "2019-06-02",
		DateTime:    "2019-06-02T12:00:00-07:00",
		EventNumber: "206",
		Hare:        "DiaB",
		EventName:   "A Super Exciting Event!",
		Description: "Really just another hash.",
		MapLink:     "https://maps.google.com/maps?q=Greenwood%20Park%2C%208905%20Fremont%20Ave%20N%2C%20Seattle%2C%20WA%2098103",
		Kennel:      "PUGET_SOUND",
	}
}

func _TestTmp(t *testing.T) {
	val := "{\"id\":\"0odi5geepodms3144rv27pr7hp_20190602\",\"date\":\"\",\"dateTime\":\"2019-06-02T12:00:00-07:00\",\"summary\":\"HS!WTF?H3 - 206TH\",\"location\":\"Bremerton, WA, USA\",\"description\":\"-WHERE?\\nTBD\\n\\n-WHEN?\\nSunday 6/2/19 @ 12:00 PM\\nHares away - 12:30 PMish\\n\\n-WHO?\\nManmaker\\n\\n-WHAT?\\nTBD\\n\\n-WHAT YOU NEED ON TRAIL: \\n$7\\nA vessel\\nShiggy Gear\\nID\\nNew shoes \\nDry bag\\nLow expectations\\nSacred fruits\\nVirgins (who run for free)\"}"
	fmt.Println(val)
}
