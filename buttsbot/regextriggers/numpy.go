package regextriggers

import (
	"buttsbot/buttsbot/utils"
	"regexp"
)

var numpys = []string{
	"Num-pee.",
	"#️⃣Num💦pee",
}

var NumpyTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex:         regexp.MustCompile(`(?mi)numpy`),
	Chance:        7,
	ConditionFunc: utils.BlankConditionFunc,
	ActionFunc:    utils.CreateRandomResponseAction(numpys),
})
