package filters

import "github.com/go-rod/rod"

func GoToPeople(page *rod.Page) {
	page.MustElementR("button", "People").MustClick()
	page.MustWaitLoad()
}

func OpenFilters(page *rod.Page) {
	page.MustElementR("button", "All filters").MustClick()
	page.MustWaitLoad()
}

func ApplyFilters(page *rod.Page) {
	page.MustElementR("button", "Show results").MustClick()
	page.MustWaitLoad()
}
