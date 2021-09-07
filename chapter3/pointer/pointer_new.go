package main

import "fmt"

func main() {
	a := new(int) // 声明一个int类型的指针变量

	fmt.Println(a) // 输出a的内存地址0xc0000140a0

	fmt.Println(*a) // 返回该类型的零值：0
}
