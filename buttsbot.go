package main

import (
	"flag"
	"log"
	"regexp"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
)

func main() {
	nick := flag.String("nick", "buttsbot", "IRC nickname")
	connectionString := flag.String("server", "", "Connection string to IRC network")
	channels := flag.String("channels", "", "IRC channels to connect to.")
	var password string
	flag.StringVar(&password, "password", "", "Password for nickserv")
	flag.Parse()

	var buttTrigger = hbot.Trigger{
		func(b *hbot.Bot, m *hbot.Message) bool {
			if m.From == b.Nick {
				return false
			}
			matched, err := regexp.MatchString("bitcoin", m.Content)
			if err != nil {
				log.Fatal(err)
			}
			return matched
		},
		func(b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, "More like buttcoin, am I rite!?")
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
	}
	mybot, err := hbot.NewBot(*connectionString, *nick, options)
	if err != nil {
		panic(err)
	}

	mybot.AddTrigger(buttTrigger)
	mybot.Run()
}
