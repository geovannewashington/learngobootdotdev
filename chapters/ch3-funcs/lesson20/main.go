package main

func adder() func(int) int {
	sum := 0

	// A closure
	return func(x int) int {
		sum += x
		return sum
	}
}
