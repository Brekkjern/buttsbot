package main

import (
	"flag"
	"log"
	"strings"

	"github.com/brekkjern/buttsbot/linkpreview"
	"github.com/brekkjern/buttsbot/regextriggers"

	hbot "github.com/whyrusleeping/hellabot"

	logger "gopkg.in/inconshreveable/log15.v2"
)

func main() {
	logHandler := logger.LvlFilterHandler(logger.LvlInfo, logger.StdoutHandler)
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

	mybot.Logger.SetHandler(logHandler)
	log.Println("Adding triggers...")
	mybot.AddTrigger(regextriggers.ButtcoinTrigger)
	mybot.AddTrigger(regextriggers.HarmfulTrigger)
	mybot.AddTrigger(regextriggers.TrumpTrigger)
	mybot.AddTrigger(regextriggers.ClawTrigger)
	mybot.AddTrigger(regextriggers.FPVsOOPTrigger)
	mybot.AddTrigger(linkpreview.LinkPreviewTrigger)
	log.Println("Attempting to connect to IRC network...")
	mybot.Run()
}
