package main

import (
	config "buttsbot/buttsbot/Config"
	geminipreview "buttsbot/buttsbot/Geminipreview"
	linkpreview "buttsbot/buttsbot/Linkpreview"
	regextriggers "buttsbot/buttsbot/Regextriggers"
	"flag"

	"strings"

	hbot "github.com/whyrusleeping/hellabot"

	logger "gopkg.in/inconshreveable/log15.v2"
)

func main() {
	configPath := flag.String("c", "/etc/buttsbot/", "Path to configuration file directory")
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic("Loading config failed")
	}

	loglvl, err := logger.LvlFromString(cfg.Loglevel)
	if err != nil {
		panic("Getting loglevel failed")
	}
	log := logger.New()
	logHandler := logger.LvlFilterHandler(loglvl, logger.StdoutHandler)
	log.SetHandler(logHandler)

	log.Info("Initializing buttsbot...")
	log.Info("Config:", "cfg", cfg)

	linkpreview.TwitterAPIToken = cfg.TwitterAPIToken

	channelList := strings.Split(cfg.Channels, ":")
	options := func(bot *hbot.Bot) {
		bot.Channels = channelList
		if cfg.NickservPass != "" {
			bot.SASL = true
			bot.Password = cfg.NickservPass
		}
		bot.SSL = cfg.IrcUseSSL
		bot.HijackSession = false
	}
	log.Info("Initializing bot system...")
	mybot, err := hbot.NewBot(cfg.IrcServer, cfg.Nick, options)
	if err != nil {
		panic(err)
	}

	mybot.Logger.SetHandler(logHandler)
	log.Info("Adding triggers...")

	if loglvl == logger.LvlDebug {
		mybot.AddTrigger(hbot.Trigger{
			Condition: func(b *hbot.Bot, m *hbot.Message) bool {
				return m.To == b.Nick
			},
			Action: func(b *hbot.Bot, m *hbot.Message) bool {
				log.Debug("Message to bot:", "From", m.From, "Content", m.Content)
				return false
			},
		})
	}

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
	mybot.AddTrigger(regextriggers.TrumpTrigger)
	mybot.AddTrigger(linkpreview.LinkPreviewTrigger)
	mybot.AddTrigger(geminipreview.GeminiPreviewTrigger)

	log.Info("Attempting to connect to IRC network...")
	mybot.Run()
}
