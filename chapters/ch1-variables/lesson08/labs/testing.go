package main

// note that the unsafe module provides a Sizeof function, similar to sizeof operator in C/C++
// it returns, in bytes, the size of an object (e.g.: a variable, etc...)
import (
	"fmt"
	"unsafe"
)

func main() {
	// var num uint8 // an unsigned integer of 8 bits
				     // can hold up to 2^8 - 1, 256
    // num = 256 -> error: overflows
	var x int
	fmt.Printf("size of x: %d bytes\n", unsafe.Sizeof(x)); // 8 bytes
}
