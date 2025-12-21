package filters

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func ByLocation(page *rod.Page, location string) {
	btn, err := page.Timeout(5*time.Second).ElementR("span", "Add a location")
	if err != nil {
		return // location filter not available
	}
	btn.MustClick()

	inputBox, err := page.Timeout(5 * time.Second).
		Element(`input[placeholder*="location"]`)
	if err != nil {
		return
	}

	inputBox.MustInput(location)
	time.Sleep(800 * time.Millisecond)

	inputBox.
		MustKeyActions().
		Press(input.ArrowDown).
		Press(input.Enter).
		MustDo()
}
