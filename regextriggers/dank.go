package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var dankRegex = regexp.MustCompile(`(?mi)\b(69|420)\b`)
var dankResponses = []string{
	`lol blaze it`,
	"That's the weed number!",
}

var DankTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, dankRegex, 3)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		match := dankRegex.FindString(m.Content)
		if match == "69" {
			b.Reply(m, "nice")
		} else if match == "420" {
			selectRandomResponse(dankResponses)
		}
		return false
	},
}
