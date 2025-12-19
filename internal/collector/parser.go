package collector

import (
	"strings"

	"github.com/go-rod/rod"
)

func GetProfile(profile *rod.Element) UserProfile {
	var p UserProfile

	// 1. Try to find the Name and Link
	linkEl, err := profile.Element(`a[href*="/in/"]`)

	if err == nil {
		// Normal Profile
		href, _ := linkEl.Attribute("href")
		if href != nil {
			p.Link = strings.Split(*href, "?")[0]
		}

		// Find Name inside the link
		if nameSpan, err := linkEl.Element(`span[aria-hidden="true"]`); err == nil {
			p.Name = nameSpan.MustText()
		} else {
			p.Name = linkEl.MustText() // Fallback
		}
	} else {
		// 2. Handle "LinkedIn Member" (Restricted Profile)
		text, _ := profile.Text()
		if strings.Contains(text, "LinkedIn Member") {
			p.Name = "LinkedIn Member (Restricted)"
			p.Link = "N/A"
		}
	}

	// 3. Extract Headline (Job Title)
	if p.Name != "" {
		// Try to find the headline element (usually secondary text)
		if headlineEl, err := profile.Element(`.entity-result__primary-subtitle`); err == nil {
			p.Headline = headlineEl.MustText()
		}
	}

	p.Name = strings.TrimSpace(p.Name)
	return p
}
