package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit bool
	smsSendingLimit = true
	
	var costPerSMS float64
	costPerSMS = 12.99
	
	var hasPermission bool
	hasPermission = false
	
	var username string 
	username = "urmom"

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
