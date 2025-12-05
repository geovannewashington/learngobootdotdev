package main

// import (
// 	"fmt"
// )

func splitEmail(email string) (string, string) {
	username, domain := "", ""
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

/* func main() {
	email := "tpose@gmail.com"

	username, domain := splitEmail(email)
	fmt.Printf("username: %s, domain: %v", username, domain)
	fmt.Println()
} */
