package dto

import (
"fmt"
"github.com/magiconair/properties/assert"
"testing"
)

func TestParseEventName(t *testing.T) {
	expect(t, "HAPPY_HOUR", "Seattle H3 Hashy Hour", "", "Seattle H3 Hashy Hour");
	expect(t, "SEAMON", "TBD", "", "SeaMon H3 #? (TBD)");
	expect(t, "SEATTLE", "SH3 RDR 2016 Planning Meeting", "", "SH3 RDR 2016 Planning Meeting");
	expect(t, "HSWTF", "3rd Anal Staph Infection run", "115", "HSWTFH3 #115 (3rd Anal Staph Infection run)");
	expect(t, "SEATTLE", "Hashmas", "756", "SH3 #756 (Hashmas)");
	expect(t, "SOUTH_SOUND", "XXX-mas", "144", "SSH3 #144 (XXX-mas)");
	expect(t, "PUGET_SOUND", "Christmas Run #35, with a touch of absinthe", "941", "PSH3 #941 (Christmas Run #35, with a touch of absinthe)");
	expect(t, "OTHER", "H3 Founding Day - 1938", "", "H3 Founding Day - 1938")
}

func expect(t *testing.T, expectedKennel string, expectedName string, expectedRun string, calendarSummary string) {
	assert.Equal(t, guessKennel(calendarSummary), expectedKennel)
	assert.Equal(t, parseEventName(calendarSummary), expectedName)
	assert.Equal(t, guessEventNumber(calendarSummary), expectedRun)
}

var (
	hareA = "Puget Sound Hash House Harriers Proudly Announces Run #890, \"The City Dump Run.\"\n\nWhen: Saturday, February 15th, 2014\nWhat Time: 10:00 AM PST\nWhere: Sunset Park, SeaTac, WA\nLocation: Des Moines Memorial Way S and S 136th\n(map the directions on Google, Apple, Bing or whatever interweb map app you have)\nHares: Pus Sucker & Uncle Fucker. \n(Pus' Phone/Text #206-915-1926)\n\nCost: $5.00 USD\nTavern Piss Up, buy your own food (yes they will be open, they have biscuits & gravy, and other food)\nShiggy alert: Bring extra pair of dry shoes.\n\nSave the date: Friday July 4th, 2014 PSH3 Run #900! + Who Needs Ten Fingers Anyways Party.";
	hareB = "Hashy Hour on Valentine’s has been hijacked for a Full Moon Run!\n\nHashy Hour:\n\n5:00PM-Whenever, Targy’s on Queen Anne 600 W. Crockett St. Seattle, Wa 98119\n\nHappy Hour is from 4PM-7PM $2.75 bottles, $3.75 wells and other drink specials.  They don’t serve food but you can bring with, or order in from Zeeks, Pagliacci, or your choice of random Asian delivery (Wangover?)  For those of you too lazy (or still hurting from RDR weekend) to do the trail, at 9:30, what says \"we love dive bars\" more than live DJs spinning 80’s and 90’s Hip-Hop/House/Trap along with more drink specials?\n\nTrail: Hares away at 7PM\n\nWhat? Valentine’s isn’t about beer, it’s about inappropriate decisions after too many shots! So is this trail!\n\nHares: Hood Whornament, I’ll Take Your Cherry, and Steve BlowJobs\n\nBring: $5 hash cash, cranium lamps, your hot brother, new shoes, a cranium lamp, virgins, your on-again-off-again ex-girlfriend who you can’t remember if you already have plans with that night… ";
	hareC = "If you're in the hangover of love the day after St. Valentine's Day, or just need to run off all that chocolate, join us for the running of the Age of Aquarius. That’s right, groovy wankers and bimbos. Break out the hemp, the roach clips and your inner love child. It’s the first anal running of the Age of Aquarius in celebration of all Aquarians (especially Assfault) who want to let their hippie freak flag fly! There will be prizes for most groovetastically dressed harrier and harriette.\n\nWhat: South Sound H3 Run #111, Running of the Age of Aquarius\n\nHares: Assfault Moneyshot QC, Portabella Areola, Just Al, and Just John\n\nWhen: Saturday, Feb 15th at 2:30 p.m. Hares way at 3 p.m. sharp (that Aquarian sun sets so fast!!)\n\nWhere: Katie Downs Waterfront Tavern, 3211 Ruston Way Tacoma, WA 98402 (look for the hippies in the parking lot across the street or nearby.) \n\nWhy: To drink and be merry. You will be drunk but surrounded by glorious food options nearby.\n\nTrail: A to A. Moderate shiggy level. Bring IDs on trail. Dog-friendly, but not child-friendly.\n\nHash cash: $10 gets you beer and some hab to remember the run by. (LSD, not STDs!)\n\nAlso Bring: love (or anal) beads, favorite love child clothes and paraphernalia (more hair!), your favorite roach clip and your pregnant old lady, flowers for your hair, dance shoes and drums for a love circle in the park, and a blanket for love-making. Also a cranium lamp for circle in the park since it’ll be close to dark. Also very warm clothes for circle or the ability to use the flesh of others to warm your pole (I mean soul.) Make love, not war (unless you want to dress like a soldier and in that case you will be a converted--I mean perverted wanker!!)\n\nInfo: Assfault, 253-678-5218";
	hareD = "Hare(s): Just Evan and Just Mike";
	hareE = "HARE(s): Mexican Rimjob";
	hareF = "HARES: Snow White / Stinky in the Pink\n\nIt'll be fun. Somewhere around Greenlake. Bring a bunch of bright lights. Yada yada. More to come...\n\nhttps://www.facebook.com/events/737114099640966/";
	hareG = "Hare(S): SlapSlapFartSlapSlap & Nudie No Name";
	hareH = "Hares: Ass Hopper, Just Daryl, CCBB, NFH卍, JCGIU\nAnd other stuff here";
)

func TestHares(t *testing.T) {
	verifyHare(t, hareA)
	verifyHare(t, hareB)
	verifyHare(t, hareC)
	verifyHare(t, hareD)
	verifyHare(t, hareE)
	verifyHare(t, hareF)
	verifyHare(t, hareG)
	verifyHare(t, hareH)
}

func verifyHare(t *testing.T, s string) {
	hare := parseHare(s)
	if hare == "" {
		t.Error()
	}
	fmt.Println(hare)
}
