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

	// selectors := []string{
	// 	`button[aria-label="Next"]`,
	// 	`button[data-testid="pagination-controls-next-button-visible"]`,
	// 	`//button[.//span[text()='Next']]`,
	// }

	var btn *rod.Element
	// var err error
	// pageWithTimeout := page.Timeout(5 * time.Second)

	fmt.Println("Looking for 'Next' button...")
	/*
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
	*/

	btn = page.MustElementR("span", "Next")
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

/*
	func NextPage(page *rod.Page) bool {
		fmt.Println("Attempting pagination...")

		// ---------- 1️⃣ TRY NEXT BUTTON (SAFE) ----------
		selectors := []string{
			`button[aria-label="Next"]`,
			`button[data-testid="pagination-controls-next-button-visible"]`,
			`//button[.//span[text()='Next']]`,
		}

		for _, sel := range selectors {
			btn, err := func() (*rod.Element, error) {
				if len(sel) > 2 && sel[:2] == "//" {
					return page.Timeout(2 * time.Second).ElementX(sel)
				}
				return page.Timeout(2 * time.Second).Element(sel)
			}()

			if err != nil {
				continue
			}

			fmt.Println("Found Next button, attempting click...")

			// Scroll into view
			btn.MustScrollIntoView()
			time.Sleep(500 * time.Millisecond)

			fmt.Print(btn)
			// Try clicking SAFELY
			if err := btn.Click("left", 1); err != nil {
				fmt.Println("Next button click failed, falling back to scroll.")
				break
			}

			// Wait for new content
			time.Sleep(3 * time.Second)
			return true
		}

		// ---------- 2️⃣ FALLBACK: INFINITE SCROLL ----------
		fmt.Println("Using infinite scroll fallback...")

		prevCount := len(page.MustElements(`div[data-view-name="people-search-result"]`))

		for i := 0; i < 3; i++ {
			page.MustEval(`() => {
				let el = document.querySelector('#workspace') || document.scrollingElement;
				el.scrollTop = el.scrollHeight;
			}`)
			time.Sleep(2 * time.Second)
		}

		newCount := len(page.MustElements(`div[data-view-name="people-search-result"]`))

		fmt.Printf("Profiles before: %d | after: %d\n", prevCount, newCount)

		if newCount <= prevCount {
			fmt.Println("No new profiles loaded. End of results.")
			return false
		}

		return true
	}
*/
func ScrollPage(page *rod.Page) {
	page.MustEval(`() => {
		let el = document.querySelector('#workspace');
		if (!el) el = document.scrollingElement;
		el.scrollTop = el.scrollHeight;
	}`)
}
