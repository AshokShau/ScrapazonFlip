package scraper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ExtractMeesho struct {
	URL string
	Doc *goquery.Document
}

func NewExtractMeesho(url string) (*ExtractMeesho, error) {
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

	defer resp.Body.Close()

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
	found := false
	m.Doc.Find("span[font-size='18px'][font-weight='demi'][color='#ffffff']").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Buy Now" {
			found = true
		}
	})
	return found
}

func (m *ExtractMeesho) GetImages() []string {
	var images []string
	m.Doc.Find("img[alt][fetchpriority='high']").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			images = append(images, src)
		}

	})
	return images
}
