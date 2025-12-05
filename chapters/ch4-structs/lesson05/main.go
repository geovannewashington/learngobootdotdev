package main

type sender struct {
	user // embedded
	rateLimit int
}

type user struct {
	name   string
	number int
}
