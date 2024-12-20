package tests

import (
	"testing"

	ok "github.com/AshokShau/ScrapazonFlip"
)

// TestExtractMeesho checks the Meesho scraper functionalities
func TestExtractMeesho(t *testing.T) {
	url := "https://www.meesho.com/s/p/5vdfxi?utm_source=s_cc"
	em, err := ok.NewExtractMeesho(url)
	if err != nil {
		t.Fatalf("Error initializing ExtractMeesho: %v", err)
	}

	if em.GetTitle() == "" {
		t.Errorf("Expected non-empty title, got empty")
	}

	if em.GetPrice() == "" {
		t.Errorf("Expected non-empty price, got empty")
	}

	if em.GetRating() == "" {
		t.Errorf("Expected non-empty rating, got empty")
	}

	if em.GetReviewCount() == "" {
		t.Errorf("Expected non-empty review count, got empty")
	}

	if !em.IsAvailable() {
		t.Errorf("Expected availability, got not available")
	}

	if len(em.GetImages()) == 0 {
		t.Errorf("Expected non-empty images list, got empty")
	}
}
