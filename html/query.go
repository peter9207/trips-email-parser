package html
import(

	"log"
	"fmt"
	"io"
	"github.com/PuerkitoBio/goquery"
)



func Get(reader io.Reader, tag string){


	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
		return
	}


	// Find the review items
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})





}
