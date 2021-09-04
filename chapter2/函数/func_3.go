// 函数有多个返回值

package main

import (
	"fmt"
)

// 函数有两个返回值，一个是两数之和，另一个是两数之积
func sumAndmul(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	s, m := sumAndmul(2, 3)

	fmt.Printf("sumAndmul(2, 3)=%d, %d \n", s, m)

	// 使用空值，只取第一个返回值
	s1, _ := sumAndmul(2, 3)
	fmt.Printf("sumAndmul(2, 3)=%d \n", s1)
}

// 逃逸分析
// .\func_3.go:17:12: ... argument does not escape
// .\func_3.go:17:13: s escapes to heap
// .\func_3.go:17:13: m escapes to heap
// .\func_3.go:21:12: ... argument does not escape
// .\func_3.go:21:13: s1 escapes to heap
//  从上面可以看出变量没有溢出，s, m, s1溢出到堆上了
