package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var trumpRegex = regexp.MustCompile("(?mi)Trump")
var trumpResponses = []string{
	"That's Cheeto Benito to you!",
	"Trump? Agent Orange is more like it.",
	"Call the Avengers! Captain Chaos is loose!",
	"People talking about the Cheeto-In-Chief again?",
	"Cheez Doodle opened his mouth again?",
	"Has the Hair Apparent made his wall yet?",
	"Hair Hitler having a bad hair day today?",
	"Don't say his name! It's He-Who-Must-Not-Be-Named around here!",
	"🍊🤵👎",
}
var trumpTwitterResponses = []string{
	"What's Adolf Twitler been up to now?",
	"Boss Tweet is at it again, huh?",
	"Bumbledore tweeting again?",
	"News from the porcelain office?",
}

var TrumpTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, trumpRegex, 25)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(trumpResponses))
		return false
	},
}
