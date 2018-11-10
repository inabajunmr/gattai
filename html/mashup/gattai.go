package mashup

import (
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/inabajunmr/gattai/html/path"
)

func Gattai(url1 string, url2 string) string {

	resp1, err := http.Get(url1)
	if err != nil {
		print("Can not access to " + url1)
		print(err)
		os.Exit(1)
	}

	resp2, err := http.Get(url2)
	if err != nil {
		print("Can not access to " + url2)
		print(err)
		os.Exit(1)
	}

	absHTML1 := path.ModifyToAbsoluteURLInHTML(resp1.Body, url1)
	absHTML2 := path.ModifyToAbsoluteURLInHTML(resp2.Body, url2)

	doc1, err := goquery.NewDocumentFromReader(strings.NewReader(absHTML1))
	if err != nil {
		print("Something wrong")
		print(err)
		os.Exit(1)
	}

	doc2, err := goquery.NewDocumentFromReader(strings.NewReader(absHTML2))
	if err != nil {
		print("Something wrong")
		print(err)
		os.Exit(1)
	}

	result, _ := goquery.NewDocumentFromReader(strings.NewReader("<!DOCTYPE html><html><head></head><body></body></html>"))
	// head
	rand.Seed(time.Now().UnixNano())

	head1 := doc1.Find("head").Children()
	head2 := doc2.Find("head").Children()
	result.Find("head").AppendSelection(head1)
	result.Find("head").AppendSelection(head2)

	// body
	body1 := doc1.Find("body")
	body2 := doc2.Find("body")

	// TODO Using body1 and put into body2 lowermost element

	rbody := result.Find("body")

	for {
		s1 := extract(body1.Children())
		s2 := extract(body2.Children())

		if s1 == nil && s2 == nil {
			break
		}

		if s1 != nil {
			rbody = rbody.AppendSelection(s1)
		}

		if s2 != nil {
			rbody = rbody.AppendSelection(s2)
		}

	}

	val, err := result.Html()

	return val

}

func extract(s *goquery.Selection) *goquery.Selection {
	if len(s.Nodes) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	current := s
	for {
		index := rand.Intn(len(current.Nodes))
		node := current.Nodes[index]
		current = goquery.NewDocumentFromNode(node).Children()

		// TODO random hierarchy
		v := rand.Intn(20)
		if len(current.Nodes) == 0 || v == 0 {
			result := goquery.NewDocumentFromNode(node).Clone()
			goquery.NewDocumentFromNode(node).Remove()
			return result
		}
	}
}
