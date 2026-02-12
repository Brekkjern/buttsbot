package regextriggers

import (
	"buttsbot/buttsbot/utils"
	"regexp"
)

var sarcastic = []string{
	"Such comment, much wow.",
}

var SarcasticTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex:         regexp.MustCompile(`(?mi).*`),
	Chance:        9001,
	ConditionFunc: utils.BlankConditionFunc,
	ActionFunc:    utils.CreateRandomResponseAction(sarcastic),
})
