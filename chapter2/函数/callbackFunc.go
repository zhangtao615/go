// 回调函数

package main

import "fmt"

type callback func(int, int) int // 定义了一个callback函数类型，接收int int类型

// doAdd函数接收3个参数，前两个为int类型，第三个为callback类型，函数将调用传入的f函数并返回该函数的计算结果
func doAdd(a, b int, f callback) int {
	fmt.Println("f callback")
	return f(a, b)
}

// 函数add的函数签名和callback类型一致，add可以作为doAdd函数的第三个参数。
func add(a, b int) int {
	fmt.Println("add running")
	return a + b
}

func main() {
	a, b := 2, 3
	fmt.Println(doAdd(a, b, add))
	// 除了上面的调用方式，还可以直接定义一个匿名函数传入
	fmt.Println(doAdd(a, b, func(a, b int) int {
		fmt.Println("================")
		return a * b
	}))
}
