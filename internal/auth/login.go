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

	if page.MustHas("#username") {
		user := os.Getenv("username")
		pass := os.Getenv("password")
		fmt.Println("User:- ", user)
		fmt.Println("Pass:- ", pass)
		page.MustElement("#username").MustInput(user)
		page.MustElement("#password").MustInput(pass).MustKeyActions().Press(input.Enter).MustDo()
		page.MustWaitLoad()
	} else {
		fmt.Println("Session fetched successfully.")
	}
}
