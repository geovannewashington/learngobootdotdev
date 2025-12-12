package main

import "fmt"

func (e email) cost() int {	
	var price int
	price = 2 * len(e.body) // default subscribed
	if !e.isSubscribed {
		// not subscribed
		price = 5 * len(e.body)	
	}
	return price
}

func (e email) format() string {
	var msg string = "Not Subscribed"

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
