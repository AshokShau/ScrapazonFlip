package tests

import (
	"testing"

	ok "github.com/AshokShau/ScrapazonFlip"
)

// TestExtractFlipkart checks the Flipkart scraper functionalities
func TestExtractFlipkart(t *testing.T) {
	url := "https://dl.flipkart.com/s/qH9GzpNNNN"
	ef, err := ok.NewExtractFlipkart(url)
	if err != nil {
		t.Fatalf("Error initializing ExtractFlipkart: %v", err)
	}

	if ef.GetTitle() == "" {
		t.Errorf("Expected non-empty title, got empty")
	}

	if ef.GetPrice() == "" {
		t.Errorf("Expected non-empty price, got empty")
	}

	if ef.GetRating() == "" {
		t.Errorf("Expected non-empty rating, got empty")
	}

	if len(ef.GetImages(500, 500, 100)) == 0 {
		t.Errorf("Expected non-empty images list, got empty")
	}
}
