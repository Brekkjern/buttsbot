package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var hateRegex = regexp.MustCompile("(?mi)\bhate\b")

var HateTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, hateRegex, 3)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, "hate is such a strong word")
		return false
	},
}
