package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var npmRegex = regexp.MustCompile(`(?mi)(yarn|npm|node\.js)`)
var npmResponses = []string{
	"1/3rd of the worlds bandwidth is Netflix. The rest is re-downloading node_modules",
	"This fixes all your problems: rm -rf node_modules && npm install",
}

var NPMTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, npmRegex, 60)
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(npmResponses))
		return false
	},
}
