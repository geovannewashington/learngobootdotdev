package main

import (
	"fmt"
)

// -- We can build our own custom types that the implement the error interface
// -- This struct implicitily implements the error interface
type userError struct {
	name string
}

func (u userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", u.name)
}

func sendSMS(msg, userName string) (string, error) {
	if !canSendToUser(userName) {
		return "", userError{name: userName}
	}	
	return msg, nil
}

func canSendToUser(userName string) bool {
	if len(userName) >= 5 {
		return true
	}
	return false
}

func main() {
	userName := "John"
	_, err := sendSMS("The quick brown fox jumps over the lazy dog.", userName)
	if err != nil {
		fmt.Println(err)
	}
}
