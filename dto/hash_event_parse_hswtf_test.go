package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEventName_hswtf(t *testing.T) {
	expectHswtfDetails(t, "HS!WTF?H3 - 204TH - Drinko De Mayo", "Drinko De Mayo", "204")
	expectHswtfDetails(t, "HS!WTF?H3 - 205TH - Alzheimer’s Bday", "Alzheimer's Bday", "205")
	expectHswtfDetails(t, "HS!WTF?H3 - 206TH", "Holy Shit! WTF? H3", "206")
}

func expectHswtfDetails(t *testing.T, given, expectedName, expectedRun string) {
	foundName,foundRun := parseEventName_hswtf(given)
	assert.Equal(t, expectedName, foundName)
	assert.Equal(t, expectedRun, foundRun)
}
