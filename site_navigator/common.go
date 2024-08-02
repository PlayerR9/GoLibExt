package site_navigator

import (
	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

// WaitFunc is a function that waits for a page to load.
//
// Parameters:
//   - url: The URL of the page to wait for.
//
// Returns:
//   - chromedp.Tasks: The tasks to wait for the page to load.
type WaitFunc func(url string) chromedp.Tasks

// ExtractFunc is a function that extracts data from the HTML.
//
// Parameters:
//   - doc: The HTML node of the page.
//
// Returns:
//   - T: The data extracted from the HTML.
//   - error: The error that occurred while extracting the data.
type ExtractFunc[T any] func(doc *html.Node) (T, error)
