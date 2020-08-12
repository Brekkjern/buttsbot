package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var bigoRegex = regexp.MustCompile(`(?mi)(big-o|O\(([1n]|log n|n\^2)\))`)
var bigoResponses = []string{
	"Your mamma's so fat, when she sat on a binary tree she flattened it in O(1) time!",
	"O(yeah)",
	"O(mg!)",
	"O(no!)",
	"O(fuck)",
}

var BigOTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, bigoRegex, 3)
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(bigoResponses))
		return false
	},
}
