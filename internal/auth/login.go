package auth

import (
	"fmt"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/joho/godotenv"
)

func LogIn(page *rod.Page) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Can't get the environment variables")
		return
	}

	if page.MustHas(".global-nav") {
		fmt.Println("Already logged in.")
		return
	}

	if page.MustHas("#username") {
		user := os.Getenv("username")
		pass := os.Getenv("password")

		fmt.Println("Attempting Login...")
		page.MustElement("#username").MustInput(user)
		page.MustElement("#password").MustInput(pass).MustKeyActions().Press(input.Enter).MustDo()

		handleSecurityChallenges(page)

		page.MustWaitLoad()
		fmt.Println("Login flow complete.")
	}
}

func handleSecurityChallenges(page *rod.Page) {
	err := page.Race().
		Element("#global-nav-search").MustHandle(func(e *rod.Element) {
		fmt.Println("Login Successful!")
	}).
		Element("input[name='pin'], #input__email_verification_pin").MustHandle(func(e *rod.Element) {
		fmt.Println("2FA Detected!")
		fmt.Print("Please enter the code sent to your email/phone: ")

		var code string
		fmt.Scanln(&code)

		e.MustInput(code).MustType(input.Enter)

		page.MustWaitStable()
		fmt.Println("2FA Submitted.")
	}).
		Element(".g-recaptcha, iframe[src*='captcha'], #captcha-internal").MustHandle(func(e *rod.Element) {
		fmt.Println("CAPTCHA Detected!")
		fmt.Println("Please switch to the browser window and solve the CAPTCHA manually.")
		fmt.Println("Press [ENTER] in this terminal once you have solved it and the page loads.")
		fmt.Scanln()
		fmt.Println("Resuming automation...")
	}).
		Element(".error-message, #error-for-password").MustHandle(func(e *rod.Element) {
		fmt.Println("Login Failed: Invalid credentials.")
		os.Exit(1)
	}).
		MustDo()

	if err != nil {
		fmt.Println("Race error (or timeout):", err)
	}
}
