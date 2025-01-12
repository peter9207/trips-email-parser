package html

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Get(reader io.Reader, tag string, text string) {

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Find the review items
	doc.Find(tag).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := strings.TrimSpace(s.Text())

		contains := strings.Cont
		ins(title, text)
		// fmt.Printf("contains  %s, %v,  %v", title, text, contains)

		if contains {
			fmt.Printf("%s %d %s: %s\n", tag, i, text, title)
		}

	})

}
