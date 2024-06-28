package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var cmdRegex = regexp.MustCompile("(?mi)\bcmd\b")

var cmdResponses = []string{
	"I'd just like to interject for a moment. What you just referred to as cmd, is in fact conhost.exe, or as I've recently taken to calling it, NT/Console Host. cmd is not a terminal unto itself, but rather another proprietary component of a fully functioning Windows console, comprised of terrible windowing, bad utilities, and other vital components comprising a full console as defined by Steve Ballmer.",
}

var CmdTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, cmdRegex, 10)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(cmdResponses))
		return false
	},
}
