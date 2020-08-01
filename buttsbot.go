package main

import (
	"flag"
	"log"
	"strings"

	"github.com/brekkjern/buttsbot/regextriggers"

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
	mybot.AddTrigger(regextriggers.GetButtTrigger())
	mybot.AddTrigger(regextriggers.GetHarmfulTrigger())
	mybot.AddTrigger(regextriggers.GetTrumpTrigger())
	mybot.AddTrigger(regextriggers.GetClawTrigger())
	log.Println("Attempting to connect to IRC network...")
	mybot.Run()
}
