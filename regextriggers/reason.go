package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var reasonRegex = regexp.MustCompile("(?mi)for (a|this|some) reason")

var ReasonTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, reasonRegex, 3)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, "pfft. a Modernist!")
		return false
	},
}
