package linkpreview

import (
	"buttsbot/buttsbot/utils"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"

	hbot "github.com/whyrusleeping/hellabot"

	logger "gopkg.in/inconshreveable/log15.v2"

	"mvdan.cc/xurls/v2"
)

const (
	maxTitleLength  = 140
	maxLinksToFetch = 2
)

var lgr = logger.Root()

var twitterDomains = []string{
	"fxtwitter.com",
	"twitter.com",
	"mobile.twitter.com",
	"www.x.com",
	"x.com",
	"mobile.x.com",
}

var rx = xurls.Strict()

func linkPreviewCondition(b *hbot.Bot, m *hbot.Message) bool {
	if m.Command == "PART" || m.Command == "QUIT" {
		return false
	}
	if m.From == b.Nick || m.To == b.Nick {
		return false
	}

	urls := rx.FindAllString(m.Content, -1)
	if urls == nil {
		return false
	}

	for i := range urls {
		url := urls[i]
		if strings.HasPrefix(url, "http") {
			return true
		}
	}

	return false
}

func linkPreviewAction(b *hbot.Bot, m *hbot.Message) bool {
	r := rx.FindAllString(m.Content, -1)
	lgr.Debug("Found links for linkpreview", "url", r)

	for p := range r {
		if p > maxLinksToFetch {
			break
		}
		pu, _ := url.Parse(r[p])

		var reply = ""
		var err error = nil

		switch site := getSite(pu); site {
		case YouTube:
			reply, err = previewYoutubeLink(pu)
		case Twitter:
			reply, err = previewTwitterLink(pu)
		case DefaultSite:
			reply, err = previewDefaultLink(pu)
		}
		if err != nil {
			lgr.Debug("Failed to fetch title", "url", r[p], "err", err)
		}

		b.Reply(m, reply)
	}

	return false
}

type Site int

const (
	YouTube Site = iota
	Twitter
	DefaultSite
)

func getSite(u *url.URL) Site {
	if isYoutube(u) {
		return YouTube
	}
	if slices.Contains(twitterDomains, u.Host) {
		return Twitter
	}
	return DefaultSite
}

var LinkPreviewTrigger = hbot.Trigger{
	Condition: linkPreviewCondition,
	Action:    linkPreviewAction,
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
	t = formatTitle(t)
	return utils.EllipticalTruncate(t, maxTitleLength), nil
}
