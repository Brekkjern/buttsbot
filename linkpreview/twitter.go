package linkpreview

import (
	"errors"
	"math/rand"
	"net/url"
)

var twitterFrontend = "nitter.net"

func previewTwitterLink(loc *url.URL) (string, error) {
	var preview = "Twitter - "
	if rand.Intn(4) == 1 {
		preview = "Twatter - "
	}
	if loc.Host != "twitter.com" {
		return "", errors.New("previewTwitterLink() called for non-twitter link")
	}
	altLocation := loc
	altLocation.Host = twitterFrontend
	preview += altLocation.String()
	pageData := fetchContents(altLocation.String())
	if len(pageData) > 0 {
		title := getTitle(pageData)
		if len(title) > 0 {
			preview += " - " + title
		}
	}
	return preview, nil
}
