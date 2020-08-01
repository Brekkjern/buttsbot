package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var clawRegex = regexp.MustCompile(`(?mi)v\.v\.v`)
var clawResponses = []string{
	`( ＾◡＾)っ (‿|‿)`,
	`(‿|‿)`,
	`(‿ˠ‿)`,
	`(☞ ͡° ͜ʖ ͡°)☞ (‿ˠ‿)`,
}

var ClawTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, clawRegex, 4)
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(clawResponses))
		return false
	},
}
