package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var developersRegex = regexp.MustCompile("(?mi)devs|developers|Ballmer")

var DevelopersTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, developersRegex, 7)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, "👏 developers 👏 developers 👏 developers 👏")
		return false
	},
}
