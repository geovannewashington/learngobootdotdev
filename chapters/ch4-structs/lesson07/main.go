package main

import (
	"fmt"
	"unsafe"
)

type contact struct {
	sendingLimit int32
	age          int32
	userID       string // 8 bytes
}

type perms struct {
	canSend         bool // 1 byte
	canReceive      bool
	canManage       bool
	permissionLevel int  // 4 bytes 
}

func main() {
	var b bool
	var i int
	var s string 
	
	fmt.Printf("string: %d, bool: %d, int %d\n", 
		unsafe.Sizeof(s), unsafe.Sizeof(b), unsafe.Sizeof(i))
}
