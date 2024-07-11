package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var HlangTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(
			b, m, regexp.MustCompile(`(?m)(^h$)`), 1)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse([]string{
			"h",
		}))
		return false
	},
}
