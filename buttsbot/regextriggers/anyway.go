package regextriggers

import (
	"buttsbot/buttsbot/utils"
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var AnywayTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex:         regexp.MustCompile(`(?mi)oh no!?`),
	Chance:        5,
	ConditionFunc: utils.BlankConditionFunc,
	ActionFunc: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, "Anyway...")
		return false
	},
})
