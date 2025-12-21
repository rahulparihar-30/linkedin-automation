package main

import (
	"fmt"
	"linkedin-automation/internal/auth"
	"linkedin-automation/internal/browser"
	"linkedin-automation/internal/collector"
	"linkedin-automation/internal/filters"
	"linkedin-automation/internal/messaging"
	"linkedin-automation/internal/profile"
	"linkedin-automation/internal/search"
)

func main() {
	br, page := browser.New()
	defer br.MustClose()

	// 1️⃣ Login
	auth.LogIn(page)

	collector.InitCSV("profiles.csv")
	defer collector.CloseCSV()

	collector.LoadExisting("profiles.csv")

	search.Search(page, "Adobe")
	filters.GoToPeople(page)
	filters.OpenFilters(page)
	// Optional
	filters.ByLocation(page, "Nepal")
	// Optional
	// filters.ByKeywords(page, filters.Keywords{
	// 	FirstName: "Rahul",
	// 	Title:     "Software Developer",
	// })
	filters.ApplyFilters(page)

	collector.ExtractAllPageNo(page)
	fmt.Println("All profiles are collected.")

	profile.LoadData("profiles.csv")
	profile.ConnectAll(page)

	followUpTemplate := `Hi {{firstName}}, thanks for connecting!
Looking forward to staying in touch 🙂`
	messaging.ProcessNewConnections(page, followUpTemplate)

}
