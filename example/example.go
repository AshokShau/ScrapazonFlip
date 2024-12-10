package main

import (
	"fmt"
	"log"

	scraper "github.com/AshokShau/ScrapazonFlip"
)

func main() {
	// Example for Amazon product scraping
	amazonURL := "https://amzn.in/d/eJkEB5o"
	amazonScraper, err := scraper.NewExtractAmazon(amazonURL)
	if err != nil {
		log.Fatalf("Error initializing Amazon scraper: %v", err)
	}

	fmt.Println("Amazon Product Details:")
	fmt.Println("Title:", amazonScraper.GetTitle())
	fmt.Println("Price:", amazonScraper.GetPrice())
	fmt.Println("Rating:", amazonScraper.GetRating())
	fmt.Println("Reviews:", amazonScraper.GetReviewCount())
	fmt.Println("Available:", amazonScraper.IsAvailable())
	fmt.Println("Images:", amazonScraper.GetImages())
	deal, regularPrice := amazonScraper.HasDeal(true)
	fmt.Println("Has Deal:", deal, "Regular Price:", regularPrice)

	// Example for Flipkart product scraping
	flipkartURL := "https://www.flipkart.com/s/qH9GzpNNNN"
	flipkartScraper, err := scraper.NewExtractFlipkart(flipkartURL)
	if err != nil {
		log.Fatalf("Error initializing Flipkart scraper: %v", err)
	}

	fmt.Println("\nFlipkart Product Details:")
	fmt.Println("Title:", flipkartScraper.GetTitle())
	fmt.Println("Price:", flipkartScraper.GetPrice())
	fmt.Println("Rating:", flipkartScraper.GetRating())
	fmt.Println("Available:", flipkartScraper.IsAvailable())
	fmt.Println("Images:", flipkartScraper.GetImages(500, 500, 100))
}
