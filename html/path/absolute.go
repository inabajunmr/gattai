package path

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ModifyToAbsoluteURL from relative path in html string
func ModifyToAbsoluteURL(html io.Reader, url string) string {
	// load html
	// get elements with url
	// modify all
	doc, err := goquery.NewDocumentFromReader(html)

	if err != nil {
		fmt.Println("Failed to parse html from ", url)
		os.Exit(1)
	}

	// from relative to absolute
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		_, exist := s.Attr("href")
		if exist {

		}

	})

	return ""
}

func modifyToAbsoluteURL(targeturl string, sourceurl string) string {
	tu, err := url.Parse(targeturl)
	if err != nil {
		// TODO
	}
	if len(tu.Host) != 0 {
		// this is already absolute path
		return targeturl
	}

	su, err := url.Parse(sourceurl)
	if err != nil {
		// TODO
	}

	if !strings.HasPrefix(tu.Path, "/") {
		p := path.Join(su.Path, tu.Path)
		tu.Path = p
	}

	tu.Host = su.Host
	tu.Scheme = su.Scheme

	return tu.String()
}
