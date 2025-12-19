package collector

import (
	"encoding/csv"
	"fmt"
	"linkedin-automation/internal/pagination"
	"os"
	"time"

	"github.com/go-rod/rod"
)

var csvFile *os.File
var csvWriter *csv.Writer

func InitCSV() {
	var err error
	csvFile, err = os.Create("profiles.csv")
	if err != nil {
		panic(err)
	}
	csvWriter = csv.NewWriter(csvFile)
	csvWriter.Write([]string{"Name", "Link", "Headline"}) // Added Headline column
	csvWriter.Flush()
	fmt.Println("CSV File Initialized.")
}

func CloseCSV() {
	if csvFile != nil {
		csvFile.Close()
		fmt.Println("CSV File Closed.")
	}
}

func ExtractAllPageNo(page *rod.Page) {
	pageNo := 1
	for {
		fmt.Printf("\n--- Processing Page %d ---\n", pageNo)
		ExtractProfiles(page)

		if !pagination.NextPage(page) {
			fmt.Println("No More Pages.")
			break
		}
		pageNo++
	}
}

func ExtractProfiles(page *rod.Page) {
	fmt.Println("   Scanning for profiles...")

	// 1. Selector for the list items
	selector := `li`

	// Wait for elements to appear
	if _, err := page.Timeout(10 * time.Second).Element(selector); err != nil {
		fmt.Println("Timeout: No list items found.")
		return
	}

	elements, _ := page.Elements(selector)
	fmt.Printf("✅ Found %d items. Filtering valid profiles...\n", len(elements))

	count := 0
	for _, el := range elements {
		info := GetProfile(el)

		// FILTERING LOGIC:
		// 1. Must have a Name
		// 2. Name must NOT be "LinkedIn Member (Restricted)"
		// 3. Link must NOT be "N/A"
		if info.Name != "" &&
			info.Name != "LinkedIn Member (Restricted)" &&
			info.Link != "N/A" {

			count++
			fmt.Printf("   -> Saving: %s\n", info.Name)

			if csvWriter != nil {
				// Use 3 columns now (Name, Link, Headline)
				csvWriter.Write([]string{info.Name, info.Link, info.Headline})
				csvWriter.Flush()
			}
		}
	}

	if count == 0 {
		fmt.Println("All profiles on this page were restricted 'LinkedIn Members'.")
	} else {
		fmt.Printf("Saved %d valid profiles.\n", count)
	}
}

type UserProfile struct {
	Name     string
	Link     string
	Headline string
}
