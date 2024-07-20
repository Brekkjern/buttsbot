# Buttsbot

The builds are currently: [![Go](https://github.com/Brekkjern/buttsbot/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/Brekkjern/buttsbot/actions/workflows/go.yml)

It has been `0` days since the last failed build.

The pinnacle of human engineering in the field of IRC bots.

## Usage

1. Create a configuration file at `/etc/buttsbot/buttsbot.env`

```env
LOGLEVEL=INFO
NICK=mybuttsbot
IRCSERVER=irc.libera.chat:6697
IRCUSESSL=true
CHANNELS="##crustaceans:#testbutt"   # colon separated list of channels
NICKSERVPASS=asdfghj
TWITTERAPITOKEN=asdfghj
```

2. Then run that shit...

```sh
$ go run buttsbot.go
```