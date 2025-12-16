package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func New() (*rod.Browser, *rod.Page) {
	url := launcher.New().
		UserDataDir("profile_data").
		Leakless(false).
		MustLaunch()

	browser := rod.New().
		ControlURL(url).
		NoDefaultDevice().
		MustConnect()

	page := browser.MustPage("https://linkedin.com/feed/")
	page.MustWindowFullscreen()

	return browser, page
}
