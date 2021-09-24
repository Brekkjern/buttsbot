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
	twitterAPIToken := flag.String("twittertoken", "", "Twitter API bearer token")
	var password string
	flag.StringVar(&password, "password", "", "Password for nickserv")
	ssl := flag.Bool("ssl", false, "Enable SSL for connection")
	flag.Parse()

	linkpreview.TwitterAPIToken = *twitterAPIToken

	channelList := strings.Split(*channels, ":")
	options := func(bot *hbot.Bot) {
		bot.Channels = channelList
		if password != "" {
			bot.SASL = true
			bot.Password = password
		}
		bot.SSL = *ssl
		bot.HijackSession = false
	}
	log.Println("Initializing bot system...")
	mybot, err := hbot.NewBot(*connectionString, *nick, options)
	if err != nil {
		panic(err)
	}

	mybot.Logger.SetHandler(logHandler)
	log.Println("Adding triggers...")

	var messageLogger = hbot.Trigger{
		Condition: func(b *hbot.Bot, m *hbot.Message) bool {
			return m.To == b.Nick
		},
		Action: func(b *hbot.Bot, m *hbot.Message) bool {
			log.Println("Message to bot:", "From", m.From, "Content", m.Content)
			return false
		},
	}

	mybot.AddTrigger(messageLogger)

	mybot.AddTrigger(regextriggers.BigOTrigger)
	mybot.AddTrigger(regextriggers.ButtcoinTrigger)
	mybot.AddTrigger(regextriggers.ClawTrigger)
	mybot.AddTrigger(regextriggers.CmdTrigger)
	mybot.AddTrigger(regextriggers.DankTrigger)
	mybot.AddTrigger(regextriggers.DevelopersTrigger)
	mybot.AddTrigger(regextriggers.FPVsOOPTrigger)
	mybot.AddTrigger(regextriggers.FreedomTrigger)
	mybot.AddTrigger(regextriggers.FreenodeTrigger)
	mybot.AddTrigger(regextriggers.HarmfulTrigger)
	mybot.AddTrigger(regextriggers.HlangTrigger)
	mybot.AddTrigger(regextriggers.NPMTrigger)
	mybot.AddTrigger(regextriggers.PosixTrigger)
	mybot.AddTrigger(regextriggers.RiirTrigger)
	mybot.AddTrigger(regextriggers.StallmanTrigger)
	mybot.AddTrigger(linkpreview.LinkPreviewTrigger)
	log.Println("Attempting to connect to IRC network...")
	mybot.Run()
}
