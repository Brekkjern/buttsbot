package linkpreview

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type response struct {
	Data []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

var TwitterAPIToken string
var twitterRegex = regexp.MustCompile(`(?mi)(\w+)/status/(\d+)`)

func previewTwitterLink(loc *url.URL) (string, error) {
	return "Please pay Elon Musk $42,000 to preview this Twatter link.", nil
}

func previewTwitterLinkOld(loc *url.URL) (string, error) {
	var preview = "Twitter - "
	if rand.Intn(4) == 1 {
		preview = "Twatter - "
	}
	if loc.Host != "twitter.com" && loc.Host != "mobile.twitter.com" {
		return "", errors.New("previewTwitterLink() called for non-twitter link")
	}

	m := twitterRegex.FindStringSubmatch(loc.Path)
	fmt.Println("Match 0: " + m[0])
	fmt.Println("Match 1: " + m[1])
	fmt.Println("Match 2: " + m[2])

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	twtr, err := url.Parse("https://api.twitter.com/2/tweets")
	if err != nil {
		return "", errors.New("previewTwitterLink() has a malformed API url")
	}
	vals := url.Values{}
	vals.Add("tweet.fields", "text")
	vals.Add("ids", m[2])
	twtr.RawQuery = vals.Encode()

	req, err := http.NewRequest("GET", twtr.String(), nil)
	req.Header.Add("Authorization", "Bearer "+TwitterAPIToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Buttsbot link-previews")

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("previewTwitterLink() failed to fetch website")
	}

	defer resp.Body.Close()
	pdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("previewTwitterLink() got no valid data from API")
	}

	var responseData response
	json.Unmarshal(pdata, &responseData)

	preview += m[1] + " - "
	if responseData.Data == nil {
		return "", errors.New("previewTwitterLink() got no data from the API")
	}
	preview = html.UnescapeString(preview)
	var n = regexp.MustCompile(`\n`)
	preview += n.ReplaceAllString(responseData.Data[0].Text, " ")

	if len(preview) > 200 {
		preview = string([]rune(preview)[:147]) + "..."
	}

	return preview, nil
}
