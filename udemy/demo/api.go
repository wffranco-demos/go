package main

import (
	"fmt"
	"log"

	"github.com/wffranco-demos/go/udemy/api"
)

func main() {
	// get example
	data, err := api.Get("users/1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	fmt.Println()

	// get user
	user, err := api.GetUser(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	fmt.Println()

	// create user
	newUser, err := api.CreateUser("Willem", "Developer")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUser)
	fmt.Println()
}
