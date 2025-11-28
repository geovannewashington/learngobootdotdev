package main

import (
	"errors"
	"fmt"
	"strings"
	"os"
)

func get_creator(os string) (string, error) {
	if os == "" {
		return "", errors.New("string cannot be empty")
	}
	
	os = strings.ToLower(os)

	var creator string
	
	switch os {
		case "linux":
			creator = "Linus Torvalds"
		case "windows":
			creator = "Bill Portões"

		// all three of these case wil set creator to 'Steve Jobs'
		case "macos":
			fallthrough
		case "mac os x":
			fallthrough
		case "mac":
			creator = "Steve Jobs"
		
		default:  
			creator = "Unknown"
	}

	return creator, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run ./labs.go <os name>\n")
		return
	}
	
	operating_s := os.Args[1];
	os_creator, err := get_creator(operating_s)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("os: creator: %s\n", os_creator)
}
