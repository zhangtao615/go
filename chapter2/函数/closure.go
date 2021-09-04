package main

import (
	"fmt"
)

func sum(a int, b int) func(c int, d int) int {
	return func(c, d int) int {
		return a + b + c + d
	}
}

func main() {
	c := sum(1, 2)(3, 4)
	fmt.Printf("sum(1, 2)(3, 4)=%d \n", c)
}

// colsure 在函数中返回函数，返回的函数签名必须和函数返回值类型一样。
// 进行变量逃逸分析 go build -gcflags "-l -m" 文件名
