package linkpreview

import (
	"errors"
	"html"
	"net/url"
	"strings"
)

func formatTitle(t string) string {
	title := html.UnescapeString(t)
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, "\n", "")
	return title
}

func previewDefaultLink(loc *url.URL) (title string, err error) {
	pageData, err := fetchContents(loc.String())
	if err != nil {
		return
	}
	title, err = getTitle(pageData)
	if err != nil {
		return
	}
	if len(title) == 0 {
		return "", errors.New("Title too short")
	}
	return
}
