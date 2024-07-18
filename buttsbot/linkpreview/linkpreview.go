package linkpreview

import (
	"buttsbot/buttsbot/utils"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"

	hbot "github.com/whyrusleeping/hellabot"

	logger "gopkg.in/inconshreveable/log15.v2"
)

var lgr = logger.Root()
var maxTitleLength = 140

var twitterDomains = []string{
	"fxtwitter.com",
	"twitter.com",
	"mobile.twitter.com",
	"www.x.com",
	"x.com",
	"mobile.x.com",
}

var linkPreviewRegex = regexp.MustCompile(`(?mi)https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
var LinkPreviewTrigger = hbot.Trigger{
	Condition: func(b *hbot.Bot, m *hbot.Message) bool {
		if m.Command == "PART" || m.Command == "QUIT" {
			return false
		}
		if m.From == b.Nick || m.To == b.Nick {
			return false
		}

		return linkPreviewRegex.MatchString(m.Content)

	},
	Action: func(b *hbot.Bot, m *hbot.Message) bool {
		r := linkPreviewRegex.FindAllString(m.Content, -1)
		lgr.Debug("Found links for linkpreview", "url", r)
		for p := range r {
			if p > 2 {
				break
			}
			parsedUrl, _ := url.Parse(r[p])
			if slices.Contains(twitterDomains, parsedUrl.Host) {
				reply, err := previewTwitterLink(parsedUrl)
				if err != nil {
					lgr.Error("previewTwitterLink failed", "url", r[p], "error", err)
				}
				b.Reply(m, reply)
				return false
			}
			if isYoutube(parsedUrl) {
				reply, err := previewYoutubeLink(parsedUrl)
				if err == nil {
					b.Reply(m, reply)
				} else {
					lgr.Error("previewYoutubeLink failed", "url", r[p], "error", err)
				}
				return false
			}
			pageData, err := fetchContents(r[p])
			if err != nil {
				lgr.Info("Failed to fetch website contents.", "url", r[p], "error", err)
				return false
			}
			title, err := getTitle(pageData)
			if err != nil {
				lgr.Error("Error with page body", "url", r[p], "error", err)
				return false
			}
			if len(title) >= 1 {
				t = strings.TrimSpace(title)
				lgr.Info("Found title for URL", "title", t, "url", r[p])
				b.Reply(m, t)
			} else {
				lgr.Info("Found no title for URL", "url", r[p])
			}
		}
		return false
	},
}

func fetchContents(url string) (io.ReadCloser, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   2 * time.Second,
				KeepAlive: 2 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   2 * time.Second,
			ResponseHeaderTimeout: 2 * time.Second,
		},
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		lgr.Error("Failed to create new request", "error", err)
		return nil, err
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Buttsbot link-previews")
	request.Header.Set("Range", "bytes=0-12000")

	resp, err := client.Do(request)
	if err != nil {
		lgr.Info("Failed to fetch website", "error", err)
		return nil, err
	}
	return resp.Body, nil
}

func getTitle(r io.ReadCloser) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	defer r.Close()
	if err != nil {
		return "", err
	}

	t := doc.Find("title").First().Contents().Text()
	return utils.EllipticalTruncate(t, maxTitleLength), nil
}
