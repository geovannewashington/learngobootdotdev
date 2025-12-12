package main

import "fmt"

func (e email) cost() int {	
	price := 2           // Default subscribed
	if !e.isSubscribed { // Not subscribed
		price = 5	
	}
	return price * len(e.body)
}

func (e email) format() string {
	msg := "Not Subscribed"

	if e.isSubscribed {
		msg = "Subscribed"
	}
	return fmt.Sprintf("'%v' | %s", e.body, msg);
}

// -- Interfaces
type expense interface {
	cost() int
}

type formatter interface {
	format() string
}
// -- 

type email struct {
	isSubscribed bool
	body         string
}
