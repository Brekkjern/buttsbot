package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var freedomRegex = regexp.MustCompile("(?m)FREED(OM|UMB)")

var FreedomTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, freedomRegex, 3)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Action(m.To, "SCREECHES!")
		return false
	},
}
