package Linkpreview

import (
	"errors"
	"html"
	"net/url"
)

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
	if !isYoutube(loc) {
		return "", errors.New("previewYoutubeLink() called for non-youtube link")
	}
	pageData := fetchContents(loc.String())
	if len(pageData) > 0 {
		title := getTitle(pageData)
		if len(title) > 0 {
			preview = title
		}
	}
	preview = html.UnescapeString(preview)
	preview += " - " + loc.String()
	return preview, nil
}
