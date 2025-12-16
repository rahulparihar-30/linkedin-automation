package filters

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func ByCompany(page *rod.Page, company string) {
	page.MustElementR("button", "Current companies").MustClick()

	inputBtn := page.MustElement(`input[placeholder="Add a company"]`)
	inputBtn.MustInput(company)
	inputBtn.MustKeyActions().Press(input.Enter).MustDo()
}
