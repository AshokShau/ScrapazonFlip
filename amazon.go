package scraper

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ExtractAmazon struct {
	URL string
	Doc *goquery.Document
}

func NewExtractAmazon(url string) (*ExtractAmazon, error) {
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

	return &ExtractAmazon{URL: url, Doc: doc}, nil
}

func (ea *ExtractAmazon) GetTitle() string {
	title := ea.Doc.Find("#productTitle").Text()
	return strings.TrimSpace(title)
}

func (ea *ExtractAmazon) GetPrice() string {
	price := ea.Doc.Find("#priceblock_ourprice").Text()
	if price == "" {
		price = ea.Doc.Find("#corePriceDisplay_desktop_feature_div span.a-price-whole").Text()
	}
	if price == "" {
		price = ea.Doc.Find("#corePrice_desktop span.a-offscreen").Text()
	}
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(price, "₹", ""), ",", ""))
}

func (ea *ExtractAmazon) GetRating() string {
	rating := ea.Doc.Find("i.a-icon.a-icon-star span.a-icon-alt").Text()
	return strings.TrimSpace(rating)
}

func (ea *ExtractAmazon) GetReviewCount() string {
	reviews := ea.Doc.Find("#acrCustomerReviewText").Text()
	return strings.TrimSpace(reviews)
}

func (ea *ExtractAmazon) IsAvailable() bool {
	available := ea.Doc.Find("#add-to-cart-button").Length() > 0
	return available
}

func (ea *ExtractAmazon) GetImages() []string {
	var images []string
	ea.Doc.Find("#imgTagWrapperId img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			images = append(images, src)
		}
	})
	return images
}

func (ea *ExtractAmazon) HasDeal(getRegularPrice bool) (bool, string) {
	dealBadge := ea.Doc.Find("#dealBadgeSupportingText").Length() > 0
	if getRegularPrice {
		regularPrice := ea.Doc.Find("#corePrice_feature_div .a-price.a-text-normal span.a-price-whole").Text()
		return dealBadge, strings.TrimSpace(regularPrice)
	}
	return dealBadge, ""
}
