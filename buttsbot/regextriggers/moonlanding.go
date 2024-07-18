package regextriggers

import (
	"buttsbot/buttsbot/utils"
	"regexp"
)

var moonlanders = []string {
	"Why would you break your keyboard in two? Anger issues much?",
	"The moonlander is fake! ZSA faked it!",
}

var MoonlanderTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex: regexp.MustCompile(`(?mi)moonlander`),
	Chance: 7,
	ConditionFunc: utils.BlankConditionFunc,
	ActionFunc: utils.CreateRandomResponseAction(moonlanders),
})