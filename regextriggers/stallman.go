package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var stallmanRegex = regexp.MustCompile(`(?mi)Linux`)
var stallmanResponses = []string{
	"I'd just like to interject for a moment. What you’re referring to as Linux, is in fact, GNU/Linux, or as I’ve recently taken to calling it, GNU plus Linux.",
	// Linux + Windows tweet: https://twitter.com/thatstupiddoll/status/1358789047565565953
	"I'd just like to interject for a moment. What you're referring to as Linux, is in fact, Windows Subsystem for Linux, or as I've recently taken to calling it, WSL.\nLinux is not an operating system unto itself, but rather another free app in the Microsoft Store and part of a fully functioning Windows system made useful by the Windows OS, shell utilities and vital system components comprising a full OS as defined by POSIX.",
	"I'd just like to interject for a moment. What you're referring to as Linux, is in fact, the Linux Subsystem for SystemD, or as I've recently taken to calling it, LSD. Linux is not an operating system unto itself, but rather another free component of fully functioning SystemD.",
}

var StallmanTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, stallmanRegex, 25)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(stallmanResponses))
		return false
	},
}
