package utils

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

type TriggerCondition func(b *hbot.Bot, m *hbot.Message) bool
type TriggerAction func(b *hbot.Bot, m *hbot.Message) bool

type TriggerOptions struct {
	Regex         *regexp.Regexp
	Chance        uint
	ConditionFunc TriggerCondition
	ActionFunc    TriggerAction
}

func CreateTriggerCondition(opts TriggerOptions) TriggerCondition {
	return func(b *hbot.Bot, m *hbot.Message) bool {
		if m.From == b.Nick || m.To == b.Nick {
			return false
		}
		if opts.Regex == nil || !opts.Regex.MatchString(m.Content) {
			return false
		}
		if opts.Chance == 0 || !RandomizeChance(int(opts.Chance)) {
			return false
		}
		return opts.ConditionFunc(b, m)
	}
}

func CreateTrigger(opts TriggerOptions) hbot.Trigger {
	return hbot.Trigger{
		Condition: CreateTriggerCondition(opts),
		Action:    opts.ActionFunc,
	}
}

func CreateRandomResponseAction(responses []string) TriggerAction {
	return func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, SelectRandomElement(responses))
		return false
	}
}

func BlankConditionFunc(b *hbot.Bot, m *hbot.Message) bool {
	return true
}
