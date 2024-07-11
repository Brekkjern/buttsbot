package regextriggers

import (
	"buttsbot/buttsbot/utils"
	"regexp"
)

var bikesheds = []string{
	"I'd like the bikeshed to be blue!",
	"I think yellow would be nice for the shed",
	"The shed should be red!",
	"I think the bikeshed looks nice how it is...",
}

var BikeshedTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex:         regexp.MustCompile(`(?mi)bikeshed`),
	Chance:        5,
	ConditionFunc: utils.BlankConditionFunc,
	ActionFunc:    utils.CreateRandomResponseAction(bikesheds),
})
