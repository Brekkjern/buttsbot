package regextriggers

import (
	"buttsbot/buttsbot/utils"
	"regexp"
)

var wankpanzers = []string{
	"That car is such a clustertruck...",
	"Is the Wank Panzer falling apart again?",
	"If this Deplorean hits 88 mph you're gonna see some serious shit!",
	"Isn't it called a TP Cruiser? Like 🧻?",
	"It's the Tesla Model Why",
}

var WankpanzerTrigger = utils.CreateTrigger(utils.TriggerOptions{
	Regex:         regexp.MustCompile(`(?mi)cyber ?(truck|stuck)`),
	Chance:        2,
	ConditionFunc: utils.BlankConditionFunc,
	ActionFunc:    utils.CreateRandomResponseAction(wankpanzers),
})
