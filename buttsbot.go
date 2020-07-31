package main

import (
	"flag"
	"log"
	"regexp"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
)

func main() {
	log.Println("Initializing buttsbot...")

	nick := flag.String("nick", "buttsbot", "IRC nickname")
	connectionString := flag.String("server", "", "Connection string to IRC network")
	channels := flag.String("channels", "", "IRC channels to connect to.")
	var password string
	flag.StringVar(&password, "password", "", "Password for nickserv")
	flag.Parse()

	bitcoinRegex := regexp.MustCompile("(?mi)bitcoin")
	var buttTrigger = hbot.Trigger{
		func(b *hbot.Bot, m *hbot.Message) bool {
			if m.From == b.Nick {
				return false
			}
			return bitcoinRegex.MatchString(m.Content)
		},
		func(b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, "More like buttcoin, am I rite!?")
			return false
		},
	}

	harmfulRegex := regexp.MustCompile("(?mi)considered harmful")
	var harmfulTrigger = hbot.Trigger{
		func(b *hbot.Bot, m *hbot.Message) bool {
			if m.From == b.Nick {
				return false
			}
			return harmfulRegex.MatchString(m.Content)
		},
		func(b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, "Your FACE is considered harmful!")
			return false
		},
	}

	channelList := strings.Split(*channels, ":")
	options := func(bot *hbot.Bot) {
		bot.Channels = channelList
		if password != "" {
			bot.SASL = true
			bot.Password = password
		}
		bot.HijackSession = true
	}
	log.Println("Initializing bot system...")
	mybot, err := hbot.NewBot(*connectionString, *nick, options)
	if err != nil {
		panic(err)
	}
	log.Println("Adding triggers...")
	mybot.AddTrigger(buttTrigger)
	mybot.AddTrigger(harmfulTrigger)
	log.Println("Attempting to connect to IRC network...")
	mybot.Run()
}
