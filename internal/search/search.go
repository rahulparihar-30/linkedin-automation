package search

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func Search(page *rod.Page, query string) {
	searchBox := page.MustElement(`input[aria-label="Search"]`)
	searchBox.MustClick()
	searchBox.MustInput(query)
	searchBox.MustKeyActions().Press(input.Enter).MustDo()
	page.MustWaitLoad()
}
