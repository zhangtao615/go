package main

import "fmt"

func main() {
	var a int = 1
	fmt.Println(&a) // 获取变量a的虚拟内存地址： 0xc0000b2008
}
