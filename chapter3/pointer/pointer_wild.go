// 野指针
package main

import "fmt"

func main() {
	// 指针未初始化，野指针
	var ptr *int

	// 打印指针指向的地址
	fmt.Printf("%p\n", ptr) // 0x0

	var a int = 7
	// 重新定向
	ptr = &a
	*ptr = 3

	fmt.Println(a) // 3
}
