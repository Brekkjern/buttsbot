package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var stallmanRegex = regexp.MustCompile(`(?mi)Linux`)
var stallmanResponses = []string{
	"I'd just like to interject for a moment. What you’re referring to as Linux, is in fact, GNU/Linux, or as I’ve recently taken to calling it, GNU plus Linux.",
}

var StallmanTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, stallmanRegex, 75)
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(stallmanResponses))
		return false
	},
}
