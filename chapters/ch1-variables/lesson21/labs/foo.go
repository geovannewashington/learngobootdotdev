package main

import (
	"fmt"
	"unsafe"
	"unicode/utf8"
)

func main() {	
	str := "tpose"
	var foo rune	
	fmt.Printf("size of str[0]: %d\n", unsafe.Sizeof(str[0]))
	fmt.Printf("size of foo: %d\n", unsafe.Sizeof(foo))
	fmt.Printf("runes in str: %d\n", utf8.RuneCountInString(str))

	// if unsafe.Sizeof(str[0]) == unsafe.Sizeof(foo) {
	// 	fmt.Println("same")
}
