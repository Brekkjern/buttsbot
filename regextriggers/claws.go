package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var clawRegex = regexp.MustCompile(`(?mi)v\.v\.v`)
var clawResponses = []string{
	`( ＾◡＾)っ (‿|‿)`,
	`(‿|‿)`,
	`(‿ˠ‿)`,
	`(☞ ͡° ͜ʖ ͡°)☞ (‿ˠ‿)`,
	`🍑`,
	"(‿!‿) ԅ(´ڡ`ԅ)",
	`(‿ˠ‿)( ͡⚆ ͜ʖ ͡⚆ )(‿ˠ‿)`,
	"༼ つ ✿◕‿◕✿༽つ (‿ˠ‿)",
	"( ‿|‿ ) ԅ(¯﹃¯ԅ)",
	"∈)✹(∋",
	"(‿ˠ‿) (￣ｍ￣〃)",
	"(‿.ꜟ‿) (￣ε (#￣)",
	"( ㅅ ) ゞƪ(ړײ)‎ƪ​​ゞ",
	"(‿ꜟ‿) ლ(´ڡ`ლ)",
}

var ClawTrigger = hbot.Trigger{
	func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, clawRegex, 1)
	},
	func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(clawResponses))
		return false
	},
}
