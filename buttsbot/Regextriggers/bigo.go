package regextriggers

import (
	"regexp"

	"buttsbot/buttsbot/utils"

	hbot "github.com/whyrusleeping/hellabot"
)

var BigOTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex:         regexp.MustCompile(`(?mi)(big-o|O\(([1n]|log n|n\^2)\))`),
	Chance:        10,
	ConditionFunc: func(*hbot.Bot, *hbot.Message) bool { return true },
	ActionFunc: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, "Your mamma's so fat, when she sat on a binary tree she flattened it in O(1) time!")
		return false
	},
})
