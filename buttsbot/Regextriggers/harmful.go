package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var harmfulRegex = regexp.MustCompile("(?mi)considered harmful")

var HarmfulTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, harmfulRegex, 4)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, "Your FACE is considered harmful!")
		return false
	},
}
