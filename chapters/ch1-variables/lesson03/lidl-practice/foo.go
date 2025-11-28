package main

import (
	"fmt"
	"os"
)

func main() {
	var user string	
	
	if len(os.Args) > 1 {
		user = os.Args[1]
	} else {
		user = "Joao"
	}
	
	fmt.Printf("Hello, %s\n", user);
}
