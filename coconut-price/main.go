package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type record struct {
	Date int
	Store1Price string
	Store2Price string
}

func main() {

	url := "http://www.oae.go.th/view/1/%E0%B8%A3%E0%B8%B2%E0%B8%84%E0%B8%B2%E0%B8%AA%E0%B8%B4%E0%B8%99%E0%B8%84%E0%B9%89%E0%B8%B2%E0%B8%A3%E0%B8%B2%E0%B8%A2%E0%B8%A7%E0%B8%B1%E0%B8%99/%E0%B8%A1.%E0%B8%84.63/33160/TH-TH"
	//url := "http://www.oae.go.th/view/1/%E0%B8%A3%E0%B8%B2%E0%B8%84%E0%B8%B2%E0%B8%AA%E0%B8%B4%E0%B8%99%E0%B8%84%E0%B9%89%E0%B8%B2%E0%B8%A3%E0%B8%B2%E0%B8%A2%E0%B8%A7%E0%B8%B1%E0%B8%99/%E0%B8%81.%E0%B8%9E.63/33316/TH-TH"
	//url := "http://www.oae.go.th/view/1/%E0%B8%A3%E0%B8%B2%E0%B8%84%E0%B8%B2%E0%B8%AA%E0%B8%B4%E0%B8%99%E0%B8%84%E0%B9%89%E0%B8%B2%E0%B8%A3%E0%B8%B2%E0%B8%A2%E0%B8%A7%E0%B8%B1%E0%B8%99/%E0%B8%A1%E0%B8%B5.%E0%B8%84.63/33512/TH-TH"

	err := getCoconutPrice(url)
	if err != nil {
		log.Fatalln(err)
	}
}

func getCoconutPrice(url string) error {

	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	r := 1
	n := 0

	doc.Find("#table_data").Find(".xl98").Each(func(i int, s *goquery.Selection) {

		line, _ := s.Html()
		line = strings.TrimSpace(line)

		if len(line) != 0 {

			if strings.Index(line, "span") > 0 {
				dl, _ := goquery.NewDocumentFromReader(strings.NewReader(line))
				line = dl.Text()
				//fmt.Println(line)
			}

			fmt.Printf("%d: %s\n", r, line)

			if (i % 2) != 0 {
				n++
				r++
			} else {
				n = 0
			}
		}
	})

	return nil
}
