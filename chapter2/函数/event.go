// 回调函数——模拟事件的绑定机制
package main

import "fmt"

// 定义类型
type OnSumBefore func(int) int
type OnSum func(int, int) int
type OnSumEnd func(string)

// 定义变量
var SumBeforeEvent OnSumBefore
var SumEvent OnSum
var SumEndEvent OnSumEnd

// Start Sum 启动计算
func StartSum(a, b int, c string) int {
	t, f := 0, 0
	// 判定释放绑定的事件，并按事件执行的顺序
	if SumBeforeEvent != nil {
		t = SumBeforeEvent(a)
	}
	if SumEvent != nil {
		f = SumEvent(t, b)
	}
	if SumEndEvent != nil {
		SumEndEvent(c)
	}
	return f
}

// RegEvent 注册事件的实现
func RegEvent(f1 OnSumBefore, f2 OnSum, f3 OnSumEnd) {
	SumBeforeEvent = f1
	SumEvent = f2
	SumEndEvent = f3
}

func main() {
	f1 := func(a int) int {
		fmt.Println("====SumBeforeEvent====")
		return a + 1
	}
	f2 := func(a int, b int) int {
		fmt.Println("====SumEvent====")
		return a + b
	}
	f3 := func(c string) {
		fmt.Println("====SumEndEvent====")
		fmt.Println(c)
	}
	RegEvent(f1, f2, f3)
	f := StartSum(3, 7, "End")
	fmt.Println(f)
}

// 函数执行结果
// ====SumBeforeEvent====
// ====SumEvent====
// ====SumEndEvent====
// End
// 11
