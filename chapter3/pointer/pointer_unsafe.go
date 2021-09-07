package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int32 = 8
	var b int64 = 12
	ptr := &a // *int32

	fmt.Println(ptr) // 0xc00001409c

	ptr = (*int32)(unsafe.Pointer(&b)) // *int64 -> *int32
	fmt.Println(ptr)                   // 0xc0000140a0

	*ptr = 10

	fmt.Println(a) // 8
	fmt.Println(b) // 10
}
