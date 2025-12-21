package collector

import (
	"encoding/csv"
	"fmt"
	"io"
	"linkedin-automation/internal/pagination"
	"os"
	"time"

	"github.com/go-rod/rod"
)

var csvFile *os.File
var csvWriter *csv.Writer
var seen = make(map[string]bool)

func InitCSV(file_path string) {
	var err error

	csvFile, err = os.OpenFile(
		file_path,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0644,
	)
	if err != nil {
		panic(err)
	}

	csvWriter = csv.NewWriter(csvFile)
	fmt.Println("CSV File Initialized.")
}

func CloseCSV() {
	if csvFile != nil {
		csvFile.Close()
		fmt.Println("CSV File Closed.")
	}
}

func LoadExisting(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if len(row) > 0 {
			seen[row[0]] = true
		}
	}
}

func ExtractAllPageNo(page *rod.Page) {
	pageNo := 1

	for {
		fmt.Printf("\n--- Processing Page %d ---\n", pageNo)

		page.Timeout(5 * time.Second).
			MustElement(`div[data-view-name="people-search-result"]`)

		ExtractProfiles(page)

		if !pagination.NextPage(page) {
			break
		}

		pageNo++
		if pageNo > 3 {
			break
		}
	}
}

func ExtractProfiles(page *rod.Page) {
	fmt.Println("   Scanning for profiles...")

	// Selector for the list items (cards)
	selector := `div[data-view-name="people-search-result"]`

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

			if !seen[info.Link] {
				seen[info.Link] = true
				count++
				fmt.Printf("   -> Saving: %s\n", info.Name)

				if csvWriter != nil {
					// Use 3 columns (Link, Name, Headline)
					// Important: Link must be first for profile.LoadData to work correctly
					csvWriter.Write([]string{info.Link, info.Name})
					csvWriter.Flush()
				}
			}
		}
	}

	if count == 0 {
		fmt.Println("No new valid profiles found on this page.")
	} else {
		fmt.Printf("Saved %d new valid profiles.\n", count)
	}
}

type UserProfile struct {
	Name     string
	Link     string
	Headline string
}
