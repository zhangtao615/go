// 变长函数：可以接受任意多个参数，在接受参数时使用...
package main

import "fmt"

func sum(ns ...int) int {
	ret := 0
	for _, n := range ns {
		ret += n
	}
	return ret
}

func main() {
	fmt.Println(sum())        // 0
	fmt.Println(sum(1))       // 1
	fmt.Println(sum(1, 2))    // 3
	fmt.Println(sum(1, 2, 3)) // 6
}
