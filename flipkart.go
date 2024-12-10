package scraper

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ExtractFlipkart struct {
	soup *goquery.Document
}

// Headers for the HTTP request
var headers = map[string]string{
	"dnt":                       "1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-user":            "?1",
	"sec-fetch-dest":            "document",
	"referer":                   "https://www.amazon.com/",
	"accept-language":           "en-GB,en-US;q=0.9,en;q=0.8",
}

// NewExtractFlipkart initializes the scraper for the given URL
func NewExtractFlipkart(url string) (*ExtractFlipkart, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return &ExtractFlipkart{soup: doc}, nil
}

// GetTitle extracts the product title
func (e *ExtractFlipkart) GetTitle() string {
	title := e.soup.Find("h1 span").Text()
	re := regexp.MustCompile(`\s{2,}`)
	return re.ReplaceAllString(strings.TrimSpace(title), " ")
}

// GetPrice extracts the product price
func (e *ExtractFlipkart) GetPrice() string {
	price := e.soup.Find("#container > div > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > div:nth-child(2) > div > div:nth-child(3) > div:nth-child(1) > div > div:nth-child(1)").Text()
	if price == "" {
		price = e.soup.Find("#container > div > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > div:nth-child(2) > div > div:nth-child(4) > div > div > div:nth-child(1)").Text()
	}
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(price), "₹", ""), ",", "")
}

// GetRating extracts the product rating
func (e *ExtractFlipkart) GetRating() string {
	rating := e.soup.Find("#container > div > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > div:nth-child(2) > div > div:nth-child(2) > div > div > span > div").Text()
	return strings.TrimSpace(rating)
}

// IsAvailable checks if the product is available
func (e *ExtractFlipkart) IsAvailable() bool {
	soldOut := e.soup.Find("#container > div > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > div:nth-child(3) > div:nth-child(1)").Text()
	return strings.TrimSpace(soldOut) != "Sold Out"
}

// GetImages extracts product images
func (e *ExtractFlipkart) GetImages(width, height, quality int) []string {
	var images []string
	e.soup.Find("ul > li > div > div > img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			src = strings.Replace(src, "?q=70", fmt.Sprintf("?q=%d", quality), 1)
			src = strings.Replace(src, "image/128/128", fmt.Sprintf("image/%d/%d", width, height), 1)
			images = append(images, src)
		}
	})
	return images
}
