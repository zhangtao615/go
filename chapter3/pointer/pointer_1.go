package main

import "fmt"

func main() {
	var a int = 7 // 声明变量a

	var p *int = &a // 声明一个指针指向a的内存地址

	fmt.Println(&a) // 0xc0000140a0
	fmt.Println(p)  // 0xc0000140a0
}
