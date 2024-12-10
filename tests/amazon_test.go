package tests

import (
	"testing"

	ok "github.com/AshokShau/ScrapazonFlip"
)

// TestExtractAmazon checks the Amazon scraper functionalities
func TestExtractAmazon(t *testing.T) {
	url := "https://amzn.in/d/eJkEB5o"
	ea, err := ok.NewExtractAmazon(url)
	if err != nil {
		t.Fatalf("Error initializing ExtractAmazon: %v", err)
	}

	if ea.GetTitle() == "" {
		t.Errorf("Expected non-empty title, got empty")
	}

	if ea.GetPrice() == "" {
		t.Errorf("Expected non-empty price, got empty")
	}

	if ea.GetRating() == "" {
		t.Errorf("Expected non-empty rating, got empty")
	}

	if len(ea.GetImages()) == 0 {
		t.Errorf("Expected non-empty images list, got empty")
	}
}
