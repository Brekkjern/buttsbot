package linkpreview

import (
	"net/url"
)

func previewTwitterLink(loc *url.URL) (string, error) {
	xcancel := *loc
	xcancel.Host = "xxcancel.com"
	xcancel.Scheme = "https"
	return xcancel.String(), nil
}
