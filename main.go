package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
)

func login(page *rod.Page) {
	user := os.Getenv("username")
	pass := os.Getenv("password")
	username := page.MustElement("#username")
	password := page.MustElement("#password")
	// signIn := page.MustElement("#organic-div > form > div.login__form_action_container > button")

	username.MustInput(user)
	password.MustInput(pass)
	password.MustKeyActions().Press(input.Enter).MustDo()
	// signIn.MustClick()
}

func showResults(page *rod.Page) {
	page.MustElementR("button", "Show results").MustClick()
	page.MustWaitLoad()
}

func filterByLocation(page *rod.Page, location string) {
	page.MustElementR("span", "Add a location").MustClick()
	inputBtn := page.MustElement(`input[placeholder="Add a location"]`)
	inputBtn.MustInput(location)
	inputBtn.MustKeyActions().Press(input.ArrowDown).Press(input.Enter).MustDo()
}

func filterByCompany(page *rod.Page, company string) {
	page.MustElementR("button", "Current companies").MustClick()

	inputBtn := page.MustElement(`input[placeholder="Add a company"]`)
	inputBtn.MustInput(company)
	inputBtn.MustKeyActions().Press(input.Enter).MustDo()

	page.MustElementR("button", "Apply").MustClick()
	page.MustWaitLoad()
}

func filterByKeywords(page *rod.Page, firstname, lastname, title, company, school string) {
	keywordsSection := page.MustElementR("h3", "Keywords").
		MustParent().
		MustElement("ul")

	inputs := keywordsSection.MustElements("input.mt1")

	if len(inputs) < 5 {
		panic("Keyword inputs not fully loaded")
	}

	if firstname != "" {
		inputs[0].MustInput(firstname)

	}
	if lastname != "" {
		inputs[1].MustInput(lastname)

	}

	if title != "" {
		inputs[2].MustInput(title)

	}

	if company != "" {
		inputs[3].MustInput(company)

	}

	if school != "" {
		inputs[4].MustInput(school)
	}
}

func openPeopleTab(page *rod.Page) {
	page.MustElementR("button", "People").MustClick()
	page.MustWaitLoad()
}

func openFiltersTab(page *rod.Page) {
	page.MustElementR("button", "All filters").MustClick()
	page.MustWaitLoad()
}

func search(page *rod.Page, query string) {
	searchBox := page.MustElement(`input[aria-label="Search"]`)
	searchBox.MustClick()
	searchBox.MustInput(query)
	searchBox.MustKeyActions().Press(input.Enter).MustDo()

	page.MustWaitLoad()
}

func main() {
	userDataDir := "profile_data"
	url := launcher.New().UserDataDir(userDataDir).Leakless(false).MustLaunch()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	browser := rod.New().
		ControlURL(url).
		NoDefaultDevice().
		MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://linkedin.com/feed/")
	page.MustWindowFullscreen()
	page.MustWaitStable()

	if page.MustHas("#username") {
		fmt.Println("User is logged out")
		login(page)
		page.MustWaitLoad()
	} else {
		fmt.Println("Session stored successfully.")
	}

	time.Sleep(3 * time.Second)

	search(page, "Software Engineer")

	time.Sleep(2 * time.Second)

	openPeopleTab(page)

	time.Sleep(2 * time.Second)

	openFiltersTab(page)

	time.Sleep(2 * time.Second)

	filterByLocation(page, "Nepal")

	fmt.Println("Scrolling to Keywords section...")

	page.MustElementR("h3", "Keywords")

	time.Sleep(1 * time.Second)

	filterByKeywords(page, "Rahul", "Parihar", "", "", "")

	// time.Sleep(2 * time.Second)

	showResults(page)

	// filterByCompany(page, "Google")

	time.Sleep(10 * time.Second)

}
