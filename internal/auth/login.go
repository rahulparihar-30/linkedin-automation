package auth

import (
	"fmt"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func LogIn(page *rod.Page) {
	if page.MustHas("#username") {
		user := os.Getenv("username")
		pass := os.Getenv("password")

		page.MustElement("#username").MustInput(user)
		page.MustElement("#password").MustInput(pass).MustKeyActions().Press(input.Enter).MustDo()
		page.MustWaitLoad()
	} else {
		fmt.Println("Session fetched successfully.")
	}
}
