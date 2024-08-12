package regextriggers

import (
	"regexp"

	"buttsbot/buttsbot/utils"

	hbot "github.com/whyrusleeping/hellabot"
)

var PythonTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex:         regexp.MustCompile(`(?mi)python`),
	Chance:        15,
	ConditionFunc: func(*hbot.Bot, *hbot.Message) bool { return true },
	ActionFunc: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, "Hiss!")
		return false
	},
})
