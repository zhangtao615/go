package main

import "fmt"

func main() {
	var a int32 = 7
	p := &a  // a的内存地址赋值给p
	pp := &p // p的内存地址赋值给pp
	**pp = 9 // 反指针 **pp -> a = 9

	fmt.Println(a)     // 9
	fmt.Println(*&a)   // 9
	fmt.Println(*&*&a) // 9
}
