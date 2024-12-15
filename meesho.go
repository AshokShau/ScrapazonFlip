package scraper

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ExtractMeesho struct {
	URL string
	Doc *goquery.Document
}

func NewExtractMeesho(url string) (*ExtractMeesho, error) {
	headers := map[string]string{
		"dnt":                       "1",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-user":            "?1",
		"sec-fetch-dest":            "document",
		"referer":                   "https://www.google.com/",
		"accept-language":           "en-GB,en-US;q=0.9,en;q=0.8",
	}

	client := &http.Client{Timeout: 20 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
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

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch the URL: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return &ExtractMeesho{URL: url, Doc: doc}, nil
}

func (m *ExtractMeesho) GetTitle() string {
	return m.Doc.Find("span[font-size='18px'][font-weight='demi'][color='greyT2']").Text()
}

func (m *ExtractMeesho) GetPrice() string {
	return m.Doc.Find("h4[font-size='32px'][font-weight='book'][color='greyBase']").Text()
}

func (m *ExtractMeesho) GetRating() string {
	return m.Doc.Find("span[font-size='16px'][font-weight='demi'][color='#ffffff']").Text()
}

func (m *ExtractMeesho) GetReviewCount() string {
	return m.Doc.Find("span[font-size='12px'][font-weight='book'][color='greyT2']").Text()

}

func (m *ExtractMeesho) IsAvailable() bool {
	return true // TODO
}

func (m *ExtractMeesho) GetImages() []string {
	var images []string
	m.Doc.Find("img[alt][fetchpriority='high']").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		images = append(images, src)

	})
	return images
}