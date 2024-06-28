package Linkpreview

import (
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	hbot "github.com/whyrusleeping/hellabot"

	logger "gopkg.in/inconshreveable/log15.v2"
)

var lgr = logger.Root()
var maxTitleLength = 140

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
			if parsedUrl.Host == "twitter.com" || parsedUrl.Host == "mobile.twitter.com" {
				reply, err := previewTwitterLink(parsedUrl)
				if err == nil {
					b.Reply(m, reply)
				} else {
					lgr.Error("previewTwitterLink failed", "url", r[p], "error", err)
				}
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
			pageData := fetchContents(r[p])
			if len(pageData) == 0 {
				lgr.Debug("No content from url", "url", r[p])
				return false
			}
			title := getTitle(pageData)
			if len(title) >= 1 {
				lgr.Info("Found title for URL", "title", title, "url", r[p])
				b.Reply(m, title)
			} else {
				lgr.Info("Found no title for URL", "url", r[p])
			}
		}
		return false
	},
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func fetchContents(url string) string {
	client := &http.Client{
		Timeout: 7 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		lgr.Error("Failed to create new request", "error", err)
		return ""
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Buttsbot link-previews")

	resp, err := client.Do(request)
	if err != nil {
		lgr.Info("Failed to fetch website", "error", err)
		return ""
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lgr.Info("Failed to read response body", "error", err)
		return ""
	}
	pageContent := string(respBytes)

	return pageContent
}

func getTitle(s string) string {
	titleStartIndex := strings.Index(s, "<title>")
	if titleStartIndex == -1 {
		return ""
	}
	// Skip to end of title declaration
	titleStartIndex += 7

	titleEndIndex := strings.Index(s, "</title>")
	if titleEndIndex == -1 {
		return ""
	}

	title := s[titleStartIndex:titleEndIndex]
	title = html.UnescapeString(title)
	title = strings.Replace(title, "\n", " - ", -1)
	title = strings.TrimSpace(title)

	if len(title) > maxTitleLength {
		title = title[:maxTitleLength] + "..."
	}

	return title
}
