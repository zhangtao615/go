package main

import "fmt"

func doPanic() {
	err := recover() // recover不会Panic异常
	if err != nil {
		fmt.Println("doPanic()捕获到异常:", err)
	}
}

func main() {
	// 注册捕获异常的处理函数
	defer doPanic()
	n := 0
	ret := 1 / n // 抛出panic异常 doPanic()捕获到异常: runtime error: integer divide by zero

	fmt.Println("main ret = ", ret)
}
