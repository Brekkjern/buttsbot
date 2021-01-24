package linkpreview

import (
	"errors"
	"net/url"
)

var youtubeFrontend = "ytprivate.com"

func isYoutube(loc *url.URL) bool {
	if loc.Host == "youtube.com" {
		return true
	}
	if loc.Host == "www.youtube.com" {
		return true
	}
	if loc.Host == "youtu.be" {
		return true
	}
	return false
}

func previewYoutubeLink(loc *url.URL) (string, error) {
	var preview = "Youtube - "
	if !isYoutube(loc){
		return "", errors.New("previewYoutubeLink() called for non-youtube link")
	}
	altLocation := loc
	altLocation.Host = youtubeFrontend
	pageData := fetchContents(altLocation.String())
	if len(pageData) > 0 {
		title := getTitle(pageData)
		if len(title) > 0 {
			preview = title
		}
	}
	preview += " - " + altLocation.String()
	return preview, nil
}
