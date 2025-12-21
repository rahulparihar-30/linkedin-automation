package messaging

import (
	"fmt"
	"linkedin-automation/internal/tracker"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

func ProcessNewConnections(page *rod.Page, template string) {
	fmt.Println("🔍 Checking for new connections...")

	// 1. Go to the 'Recently Added' connections page
	page.MustNavigate("https://www.linkedin.com/mynetwork/invite-connect/connections/")
	page.MustWaitStable()

	// 2. Scrape the visible profiles (Limit to top 10 for safety)
	// The selector finds the profile links in the list
	linkElements := page.MustElements(`a[href^="/in/"]`)

	var newProfiles []string
	for _, el := range linkElements {
		url := el.MustProperty("href").String()
		// Clean URL (remove query params)
		if strings.Contains(url, "?") {
			url = strings.Split(url, "?")[0]
		}

		// check history
		if !tracker.IsMessaged(url) {
			newProfiles = append(newProfiles, url)
		}
	}

	fmt.Printf("found %d potential new connections to message.\n", len(newProfiles))

	// 3. Visit each and send message
	for _, url := range newProfiles {
		sendMessageToProfile(page, url, template)
	}
}

func sendMessageToProfile(page *rod.Page, url, rawTemplate string) {
	fmt.Printf("👉 Visiting: %s\n", url)
	page.MustNavigate(url)
	page.MustWaitStable()

	// A. Get First Name for the template
	// Usually in the top H1 tag
	nameEl, err := page.Timeout(2 * time.Second).Element("h1")
	if err != nil {
		fmt.Println("Could not find name, skipping.")
		return
	}
	fullName := nameEl.MustText()
	firstName := strings.Split(fullName, " ")[0]

	// B. Render the message
	finalMessage := RenderTemplate(rawTemplate, firstName)

	// C. Find and Click "Message" button
	// We use Race because the button might be "Message" or inside "More"
	// (Reuse logic similar to your Connect function, but looking for "Message")
	clicked := false

	// Try finding the primary Message button
	msgBtn, err := page.Timeout(3*time.Second).ElementR("button", "Message")
	if err == nil {
		msgBtn.MustClick()
		clicked = true
	} else {
		fmt.Println("Message button not found on profile (maybe locked).")
	}

	if clicked {
		fmt.Println("📝 Typing message...")

		// Wait for the chat box to appear and focus the text area
		// The selector for the message input area (contenteditable div)
		textBox := page.MustElement(`div[role="textbox"]`)
		textBox.MustClick()

		// Clear existing text if any, then type
		textBox.MustSelectAllText().MustInput(finalMessage)
		time.Sleep(1 * time.Second) // mimic human delay

		// Find the Send button
		sendBtn := page.MustElement("button.msg-form__send-button")
		if sendBtn.MustAttribute("disabled") == nil {
			sendBtn.MustClick()
			fmt.Println("✅ Message Sent!")

			// D. Update History
			tracker.MarkAsMessaged(url)
		} else {
			fmt.Println("⚠️ Send button disabled (empty message?).")
		}

		// Close the chat window to keep UI clean
		tryCloseChat(page)
	}
}

func tryCloseChat(page *rod.Page) {
	closeBtn, err := page.Timeout(1 * time.Second).Element(`button[data-control-name="overlay.close_conversation_window"]`)
	if err == nil {
		closeBtn.MustClick()
	}
}
