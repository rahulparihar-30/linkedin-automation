package filters

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func ByLocation(page *rod.Page, location string) {
	page.MustElementR("span", "Add a location").MustClick()
	inputBtn := page.MustElement(`input[placeholder="Add a location"]`)
	inputBtn.MustInput(location)
	inputBtn.MustKeyActions().Press(input.ArrowDown).Press(input.Enter).MustDo()
}
