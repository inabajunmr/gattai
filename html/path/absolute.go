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

var targets = []struct {
	tag string
	key string
}{
	{"a", "href"},
	{"img", "src"},
	{"link", "href"},
}

// ModifyToAbsoluteURLInHTML from relative path in html string
func ModifyToAbsoluteURLInHTML(html io.Reader, url string) string {

	doc, err := goquery.NewDocumentFromReader(html)

	if err != nil {
		fmt.Println("Failed to parse html from ", url)
		os.Exit(1)
	}

	// from relative to absolute
	for _, tag := range targets {
		doc.Find(tag.tag).Each(func(i int, s *goquery.Selection) {
			val, exist := s.Attr(tag.key)
			if exist {
				s.SetAttr(tag.key, modifyToAbsoluteURL(val, url))
			}
		})

	}
	result, err := doc.Html()

	if err != nil {
		fmt.Println("Something wrong.")
		os.Exit(1)
	}

	return result
}

func modifyToAbsoluteURL(targeturl string, sourceurl string) string {
	tu, err := url.Parse(targeturl)
	if err != nil {
		fmt.Println("Something wrong.")
		os.Exit(1)
	}

	if len(tu.Host) != 0 {
		// this is already absolute path
		return targeturl
	}

	su, err := url.Parse(sourceurl)
	if err != nil {
		fmt.Println("Something wrong.")
		os.Exit(1)
	}

	if !strings.HasPrefix(tu.Path, "/") {
		p := path.Join(su.Path, tu.Path)
		tu.Path = p
	}

	tu.Host = su.Host
	tu.Scheme = su.Scheme

	return tu.String()
}
