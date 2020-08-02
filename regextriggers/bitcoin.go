package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var bitcoinRegex = regexp.MustCompile("(?mi)bitcoin")

var ButtcoinTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, bitcoinRegex, 3)
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, "More like buttcoin, am I rite!?")
		return false
	},
}
