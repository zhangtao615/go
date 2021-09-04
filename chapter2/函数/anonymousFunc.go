// 匿名函数：没有函数名，所以一般只能在定义的时候就调用。当函数只需要调用一次就可以使用匿名函数。

package main

import "fmt"

func main() {
	multiCall()
}

// 匿名函数单次调用
func singleCall() {
	// 定义匿名函数
	s, m := func(a int, b int) (int, int) {
		return a + b, a * b
	}(2, 3)

	fmt.Printf("匿名函数(2, 3)=%d, %d\n", s, m) // 匿名函数(2, 3)=5, 6
}

// 匿名函数多次调用，将匿名函数赋值给一个变量
func multiCall() {
	// 将匿名函数赋值给f
	f := func(a int, b int) (int, int) {
		return a + b, a * b
	}

	s1, m1 := f(2, 3)
	fmt.Printf("第一次调用f(2,3)=%d, %d\n", s1, m1) // 第一次调用f(2,3)=5, 6

	s2, m2 := f(6, 7)
	fmt.Printf("第二次调用f(2,3)=%d, %d\n", s2, m2) // 第二次调用f(2,3)=13, 42

}
