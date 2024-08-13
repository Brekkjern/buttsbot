package regextriggers

import (
	"fmt"
	"regexp"

	"buttsbot/buttsbot/utils"

	hbot "github.com/whyrusleeping/hellabot"
)

var ProvocativeTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex:         regexp.MustCompile(`(?mi)what does that mean`),
	Chance:        8,
	ConditionFunc: func(*hbot.Bot, *hbot.Message) bool { return true },
	ActionFunc: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, fmt.Sprintf("No one knows what it means %s, but it's provocative. It gets the people going!", m.Name))
		return false
	},
})
