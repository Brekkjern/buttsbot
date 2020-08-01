package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

func standardizedRegexTrigger(
	b *hbot.Bot,
	m *hbot.Message,
	rgx *regexp.Regexp,
	chance int) bool {
	if m.From == b.Nick {
		return false
	}
	match := rgx.MatchString(m.Content)

	if !match && !randomizeChance(chance) {
		return false
	}
	return true
}
