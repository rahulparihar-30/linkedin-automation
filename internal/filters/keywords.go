package filters

import "github.com/go-rod/rod"

type Keywords struct {
	FirstName string
	LastName  string
	Title     string
	Company   string
	School    string
}

func ByKeywords(page *rod.Page, k Keywords) {
	keywordsSection := page.MustElementR("h3", "Keywords").
		MustParent().
		MustElement("ul")

	inputs := keywordsSection.MustElements("input.mt1")

	if len(inputs) < 5 {
		panic("Keyword inputs not fully loaded")
	}

	if k.FirstName != "" {
		inputs[0].MustInput(k.FirstName)

	}
	if k.LastName != "" {
		inputs[1].MustInput(k.LastName)

	}

	if k.Title != "" {
		inputs[2].MustInput(k.Title)

	}

	if k.Company != "" {
		inputs[3].MustInput(k.Company)

	}

	if k.School != "" {
		inputs[4].MustInput(k.School)
	}
}
