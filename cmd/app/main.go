package main

import (
	"fmt"
	"linkedin-automation/internal/auth"
	"linkedin-automation/internal/browser"
	"linkedin-automation/internal/collector"
	"linkedin-automation/internal/filters"
	"linkedin-automation/internal/search"
)

func main() {
	browser, page := browser.New()
	defer browser.MustClose()

	auth.LogIn(page)
	search.Search(page, "software engineer")

	filters.GoToPeople(page)
	collector.InitCSV()
	defer collector.CloseCSV()
	fmt.Println("Done Creating the csv.")
	fmt.Println("Extracting pages")
	collector.ExtractAllPageNo(page)

	// collector.Parse(page)

	// filters.OpenFilters(page)
	// filters.ByLocation(page, "Nepal")
	// filters.ByKeywords(page, filters.Keywords{
	// 	FirstName: "Rahul",
	// 	Title:     "Software Engineer",
	// 	Company:   "Google",
	// })
	// filters.ApplyFilters(page)
}
