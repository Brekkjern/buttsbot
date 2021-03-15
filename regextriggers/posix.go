package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

//https://corrupt.tech/txt/posix-me-harder.txt

var posixRegex = regexp.MustCompile(`(?mi)(posix me harder)`)
var posixResponses = []string{
	"It's easy... you let your body tell you what to do. You know what you want... you know you want to make your body free as in freedom... libre love..",
	"\"I'll be gentle... I know all about project management,\" Eric promised.",
	"\"So... how do we do this? Do I just... open the repository and commit?\" Richard asked.",
	"\"Give me your code! I want ALL your code in my repository!\" Richard exclaimed.",
	"\"Ohhhh! I'm going to tar this up later!\" Richard exclaimed. Eric pulled out, still dripping with his source.",
	"Richard swallowed, trying not to cry. \"I'm not willing to keep love like this proprietary.\"",
	"\"Join us now... Eric... and share the hardware... my repository is open...\" Richard panted, looking at Eric's throbbing phallus.",
}

var PosixTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, posixRegex, 1)
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(posixResponses))
		return false
	},
}
