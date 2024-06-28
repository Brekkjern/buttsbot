package Geminipreview

import (
	"context"
	"io/ioutil"
	"net/url"
	"regexp"

	hbot "github.com/whyrusleeping/hellabot"

	logger "gopkg.in/inconshreveable/log15.v2"

	"git.sr.ht/~adnano/go-gemini"
)

var lgr = logger.Root()
var maxTitleLength = 140

var client = &gemini.Client{}
var ctx = context.Background()

var geminRegex = regexp.MustCompile(`(?mi)gemini://(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,24}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
var titleRegex = regexp.MustCompile(`(?mi)^# (.*)$`)

var GeminiPreviewTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		if m.Command == "PART" || m.Command == "QUIT" {
			return false
		}
		if m.From == b.Nick || m.To == b.Nick {
			return false
		}

		return geminRegex.MatchString(m.Content)
	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		urls := geminRegex.FindAllString(m.Content, -1)
		lgr.Debug("Found links for gemini preview", "url", urls)
		for u := range urls {
			if u > 2 {
				break
			}
			parsedUrl, err := url.Parse(urls[u])
			if err != nil {
				lgr.Warn("Failed to parse gemini url", "error", err, "url", urls[u])
			}

			title := ""
			fetchedTitle := getTitle(*parsedUrl)
			if len(fetchedTitle) == 0 {
				return false
			}
			favicon := getFavicon(*parsedUrl)

			if len(favicon) > 0 {
				title = favicon + " "
			}
			title = title + fetchedTitle
			if len(title) > maxTitleLength {
				title = title[:maxTitleLength] + "..."
			}
			title = title + " - " + parsedUrl.Hostname()

			b.Reply(m, title)
			b.Reply(m, "https://portal.mozz.us/gemini/"+parsedUrl.Hostname()+parsedUrl.Path)
		}
		return false
	},
}

func fetchGemini(url url.URL) string {
	resp, err := client.Get(ctx, url.String())
	if err != nil {
		lgr.Info("Failed to fetch gemini site", "error", err, "url", url.String())
		return ""
	}
	defer resp.Body.Close()
	if resp.Status != 20 {
		lgr.Info("Invalid status for gemini site", "status", resp.Status, "url", url.String())
		return ""
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lgr.Info("Failed to read data from connection", "error", err, "url", url.String())
	}
	if len(respBody) == 0 {
		lgr.Info("Body of response is empty", "url", url.String())
		return ""
	}
	return string(respBody)
}

func getFavicon(url url.URL) string {
	url.Path = "/favicon.txt"
	resp := fetchGemini(url)
	if len(resp) >= 2 {
		return string([]rune(resp)[:1])
	}
	return ""
}

func getTitle(url url.URL) string {
	resp := fetchGemini(url)
	r := titleRegex.FindStringSubmatch(resp)
	if len(r) == 0 {
		lgr.Info("No title found in body", "url", url.String())
		return ""
	}
	return r[1]
}
