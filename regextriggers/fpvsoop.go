package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var fpRegex = regexp.MustCompile(`(?m)(FP|[Ff]unctional [Pp]rogramming)`)
var oopRegex = regexp.MustCompile(`(?m)(OOP|[Oo]bject [Oo]riented)`)

var FPVsOOPTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		if m.From == b.Nick || m.To == b.Nick {
			return false
		}
		match := fpRegex.MatchString(m.Content) != oopRegex.MatchString(m.Content)

		if match && randomizeChance(6) {
			return true
		}
		return false
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		if fpRegex.MatchString(m.Content) {
			b.Reply(m, "'Foolish pupil - objects are merely a poor man's closures!' - https://wiki.c2.com/?ClosuresAndObjectsAreEquivalent")
		} else {
			b.Reply(m, "'When will you learn? Closures are a poor man's object.' - https://wiki.c2.com/?ClosuresAndObjectsAreEquivalent")
		}
		return false
	},
}
