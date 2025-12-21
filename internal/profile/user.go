package profile

import (
	"encoding/csv"
	"fmt"
	"io"
	"linkedin-automation/internal/limiter"
	"os"
	"time"

	"github.com/go-rod/rod"
)

var links []string

func LoadData(fileP string) {
	file, err := os.Open(fileP)
	if err != nil {
		return // file may not exist on first run
	}
	defer file.Close()

	reader := csv.NewReader(file)

	isHeader := false

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		// skip header row
		if isHeader {
			isHeader = false
			continue
		}

		if len(row) == 0 || row[0] == "" {
			continue
		}

		links = append(links, row[0])
	}
}

func ConnectAll(page *rod.Page) {
	for link := range links {
		GoToProfile(page, links[link])
	}
}

func GoToProfile(page *rod.Page, link string) {
	fmt.Println("Opening profile:", link)
	if !limiter.CanSendRequest() {
		return
	}
	page.MustNavigate(link)
	fmt.Println("Load")
	page.MustWaitLoad()
	Connect(page, "Hey there I hove youre doing great nice to connect.")
}

func Connect(page *rod.Page, message string) {
	fmt.Println("Trying to connect")
	actionType := 0
	_, err := page.Timeout(5*time.Second).Race().
		ElementR("button", "[Cc]onnect").MustHandle(func(e *rod.Element) {
		actionType = 1
	}).
		ElementR("span", "More").MustHandle(func(e *rod.Element) {
		actionType = 2
	}).
		Do()

	if err != nil {
		fmt.Println("Connect button not found (or already connected)")
		return
	}

	if actionType == 1 {
		page.MustElementR("button", "[Cc]onnect").MustClick()
		fmt.Println("Clicked Connect button")
	} else if actionType == 2 {
		fmt.Println("Connect button not found, clicking More")

		page.MustElementR("span", "More").MustClick()

		connectBtn := page.MustElementR(`div[role="button"]`, `(Add|Connect)`)
		connectBtn.MustClick()
		fmt.Println("Clicked Add from dropdown")
	}

	SendMessage(page, message)
	fmt.Println("Successfull")
	limiter.IncrementRequest()
}

func SendMessage(page *rod.Page, message string) {
	addNoteBtn := page.MustElementR(
		`button`,
		`Add a note`,
	)
	addNoteBtn.MustClick()

	textarea := page.MustElement("#custom-message")
	textarea.MustFocus()

	textarea.MustInput(message)

	sendBtn := page.MustElementR("span", "Send")
	sendBtn.MustClick()
}
