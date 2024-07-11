package regextriggers

import (
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"
)

var posixRegex = regexp.MustCompile(`(?mi)(posix me harder)`)
var posixResponses = []string{
	//https://corrupt.tech/txt/posix-me-harder.txt
	"It's easy... you let your body tell you what to do. You know what you want... you know you want to make your body free as in freedom... libre love..",
	"\"I'll be gentle... I know all about project management,\" Eric promised.",
	"\"So... how do we do this? Do I just... open the repository and commit?\" Richard asked.",
	"\"Give me your code! I want ALL your code in my repository!\" Richard exclaimed.",
	"\"Ohhhh! I'm going to tar this up later!\" Richard exclaimed. Eric pulled out, still dripping with his source.",
	"Richard swallowed, trying not to cry. \"I'm not willing to keep love like this proprietary.\"",
	"\"Join us now... Eric... and share the hardware... my repository is open...\" Richard panted, looking at Eric's throbbing phallus.",

	// https://corrupt.tech/txt/the-anti-patent-clause.txt
	"Richard wasn't planning on having sex today, or ever again, frankly. He was surprised to learn Eric had planned a threesome.",
	"Richard was in shock. \"Proprietary love! I don't believe this!\" he shouted.",
	"\"This passion isn't libre... but...\" for a moment, Richard suspended his principles of freedom, for a chance to make love with Eric once more.",
	"Linus, being an experienced project manager, established the pecking order. \"Here is how this is gonna work,\" he said with authority. \"I'm gonna top Eric. Eric, you're gonna top Dick. Got it? Good. These are the terms of my license. It's my show now. This will not be libre. This will be proprietary between us, under NDA. A gag order.\"",
	"\"Lube him, then lube yourself,\" Linus demanded. \"But not too much. I like a bit of interpersonal friction. Keeps people in check.\"",
	"Linus then said to Richard, \"that's my patented love-making technique, I have to ask you not do it yourself without asking me.\"",
	"Linus sighed. \"We can't do this again, Eric. I have a family. I have a wife. Imagine if she knew... this has to remain proprietary between us.\"",

	// https://corrupt.tech/txt/gnus-not-windows.txt
	"\"Go away,\" Richard said. \"You're way too early. I'm busy writing a piece aboutthe virtues of voluntary celibacy.\"",
	"Despite all the jokes, Paul was neither micro, nor soft. He ungagged Richard and then shoved all nine inches into Richard's throat. Richard gagged and tried to cough, but still found himself in the throes of utmost pleasure. He was accustomed to Microsoft dominating the industry, and now Microsoft was dominating him.",
	"It was an open-source dildo, but Paul didn't mind. It was what was on hand. The fact it wasn't made with AutoDesk CAD on Windows on a commercial 3D printer was of no consequence in the moment.",
	"Eric unpacked his thick tarball into Linus's home directory. Eric moaned loudly and gasped.",
}

var PosixTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		return standardizedRegexTrigger(b, m, posixRegex, 1)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		b.Reply(m, selectRandomResponse(posixResponses))
		return false
	},
}
