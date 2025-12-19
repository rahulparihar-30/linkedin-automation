package pagination

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func NextPage(page *rod.Page) bool {
	fmt.Println("Preparing to scroll via Mouse Wheel...")

	for i := 0; i < 5; i++ {
		fmt.Printf("  Scrolling step %d/5...\n", i+1)

		page.Mouse.Scroll(0, 1000, 1)

		time.Sleep(1 * time.Second)
	}

	fmt.Println("Waiting for footer to stabilize...")
	time.Sleep(2 * time.Second)

	selectors := []string{
		`button[aria-label="Next"]`,
		`button[data-testid="pagination-controls-next-button-visible"]`,
		`//button[.//span[text()='Next']]`,
	}

	var btn *rod.Element
	var err error
	pageWithTimeout := page.Timeout(5 * time.Second)

	fmt.Println("Looking for 'Next' button...")

	for _, sel := range selectors {
		if len(sel) > 2 && sel[:2] == "//" {
			btn, err = pageWithTimeout.ElementX(sel)
		} else {
			btn, err = pageWithTimeout.Element(sel)
		}

		if err == nil {
			fmt.Printf("Found button using: %s\n", sel)
			break
		}
	}

	if err != nil {
		fmt.Println("'Next' button truly not found.")
		return false
	}

	btn.MustScrollIntoView()

	if disabled, _ := btn.Attribute("disabled"); disabled != nil {
		fmt.Println("Button is disabled. End of pages.")
		return false
	}

	fmt.Println("Clicking Next...")
	btn.MustClick()

	// Wait for the new results to load
	fmt.Println("Waiting for page transition...")
	time.Sleep(3 * time.Second)

	return true
}
