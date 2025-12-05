package main

import (
	"fmt"
	"errors"
)

func logger() {
	fmt.Println("Exiting the function...")
}

func greetUser(user string) (greetMsg string, err error) {
	defer logger()	
	
	if user == "" {
		// logger will be executed here in case of failure
		return "", errors.New("err: username was not provided")
	}
	
	// logger will be executed here in case of success
	return "Welcome " + user + " hope you have a nice day!", nil
}

func main() {
	greetMsg, err := greetUser("")	

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(greetMsg)
}
