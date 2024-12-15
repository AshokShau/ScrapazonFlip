# ScrapazonFlip

ScrapazonFlip is a Go library that allows you to scrape product details from Amazon, Flipkart and Meesho. It provides easy-to-use methods for extracting product information like title, price, rating, availability, images, and more.

## Features

- Scrape product details from Amazon, Flipkart and Meesho.
- Extract product title, price, rating, availability, and images.
- Handle availability and product deals.
- Customizable image quality and dimensions.

## Installation

You can install ScrapazonFlip via Go modules.

```bash
go get github.com/AshokShau/ScrapazonFlip
```

## Usage Example

Below is an example demonstrating how to use the library to scrape data from both Amazon and Flipkart product pages.

[Example](/example/example.go)

```
### Example Output:

Amazon Product Details:
Title: Example Amazon Product Title
Price: ₹2,999
Rating: 4.5
Reviews: 1,234 ratings
Available: true
Images: [image1.jpg image2.jpg]
Has Deal: true Regular Price: ₹3,499

Flipkart Product Details:
Title: Example Flipkart Product Title
Price: ₹3,499
Rating: 4.2
Available: true
Images: [flipkart_image1.jpg flipkart_image2.jpg]

Meesho Product Details:
Title: Example Meesho Product Title
Price: ₹4,999
Rating: 4.05.05.0
Review Count: 20069 Ratings, 6969 Reviews
Images: [meesho_image1.jpg]
```

## Methods

### Amazon Methods: `scraper.NewExtractAmazon(amazonURL)`

- `GetTitle() string`: Returns the product title.
- `GetPrice() string`: Returns the product price.
- `GetRating() string`: Returns the product rating.
- `GetReviewCount() string`: Returns the number of reviews.
- `IsAvailable() bool`: Returns whether the product is available for purchase.
- `GetImages() []string`: Returns a list of product image URLs.
- `HasDeal(getRegularPrice bool) (bool, string)`: Returns whether the product has a deal and the regular price.

### Flipkart Methods: `scraper.NewExtractFlipkart(flipkartURL)`

- `GetTitle() string`: Returns the product title.
- `GetPrice() string`: Returns the product price.
- `GetRating() string`: Returns the product rating.
- `IsAvailable() bool`: Returns whether the product is available for purchase.
- `GetImages(width, height, quality int) []string`: Returns a list of image URLs with customizable size and quality.

### Meesho Methods: `scraper.NewExtractMeesho(meeshoURL)`

- `GetTitle() string`: Returns the product title.
- `GetPrice() string`: Returns the product price.
- `GetRating() string`: Returns the product rating.
- `GetReviewCount() string`: Returns the number of reviews.
- `GetImages() []string`: Returns a list of product image URLs.

## License

This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.
