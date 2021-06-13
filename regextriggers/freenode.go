package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var freenodeRegex = regexp.MustCompile(`(?mi)Freenode`)
var freenodeResponses = []string{
	// Gerikson...
	"I’d just like to interject for a moment. What you’re referring to as Freenode, is in fact Leenode, or as I’ve recently taken to calling it, Andrew Lee plus IRC. Freenode is not an IRC network unto itself, but rather another bastion of Andrew Lee’s princedom.",
}

var FreenodeTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, freenodeRegex, 10)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(freenodeResponses))
		return false
	},
}
