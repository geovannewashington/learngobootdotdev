package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for i := 2; i < max + 1; i++ {
		if i == 2 {
			fmt.Println(i) // The only even prime
		}

		if i % 2 == 0 {
			continue // skips all even numbers
		}
		
		res := int(math.Sqrt((float64(i))))
		isPrime := true
		for j := 3; j < res; j += 2 {
			if i % j == 0 {
				isPrime = false
				break
			}
		}
		
		if isPrime {
			fmt.Println(i)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
