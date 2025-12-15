package main 

import (
	"fmt"
	"strconv"
)

// function signature of strconv.Atoi

func main() {
	numStr := "1919"

	// check if strconv.Atoi could convert to integer
	if numDigit, err := strconv.Atoi(numStr); err != nil {
		fmt.Printf("%s\n", err)
		return
	} else {
		fmt.Printf("%d\n", numDigit)
	}
}
