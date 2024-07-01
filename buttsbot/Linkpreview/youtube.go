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
	title, err := getTitle(pageData)
	if err != nil {
		lgr.Error("Failed to get title from youtube", "url", loc, "error", err)
		return "", err
	}
	if len(title) > 0 {
		preview = title
	}
	return html.UnescapeString(preview), nil
}
