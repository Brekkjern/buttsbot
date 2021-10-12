package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var riirRegex = regexp.MustCompile(`(?mi)(\b(riir|rust(aceans)?)\b)`)
var riirResponses = []string{
	">The rusting will continue until morale improves or memory stops leaking",
	">#rust is being boycotted because the ircd has yet to be rewritten in rust",
	">Can bitcoin, if rewritten in rust, cure loneliness?",
	"That's nice, but have you considered writing it in HTML instead of Rust?",
	">I told a coworker about lobste.rs, and futilely tried to convince him that it isn't HN for Rust",
	">unsafe {} sex can be dangerous",
	">And maybe, just maybe; in those brief, wonderful moments of ecstasy — you might gain additional pleasure knowing in your heart, mind and body that your device is using a safe and fearlessly concurrent language.",
	">Then I found Rust, and now I'm going to either die of starvation or make my next job writing Rust. I just can't stand to work in this industry anymore unless I use Rust. I don't understand how we got here.",
	"🦀",
	">When I started learning Rust in earnest in 2018, I thought this was a fluke. It is just the butterflies you get when you think you fall in love, I told myself.",
	">Rust is mentioned in passing in NYT article",
	"They say if you lock the bathroom, turn off the light, face the mirror, and say, \"C++! C++! C++!\" a Rust developer will reach through to mirror to grab you.",
}

var RiirTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, riirRegex, 15)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(riirResponses))
		return false
	},
}
