package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var riirRegex = regexp.MustCompile(`(?mi)(riir|rust)`)
var riirResponses = []string{
	">The rusting will continue until morale improves or memory stops leaking",
	">#rust is being boycotted because the ircd has yet to be rewritten in rust",
	">Can bitcoin, if rewritten in rust, cure loneliness?",
	"That's nice, but have you considered writing it in HTML instead of Rust?",
	">I told a coworker about lobste.rs, and futilely tried to convince him that it isn't HN for Rust",
	">unsafe {} sex can be dangerous",
}

var RiirTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, riirRegex, 10)
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(riirResponses))
		return false
	},
}
