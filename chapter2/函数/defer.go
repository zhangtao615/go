// defer 用于注册延迟调用的一种机制。常用于对资源进行释放的场景，例如：数据库连接、解锁、关闭文件等。
package main

import "fmt"

func print(s string) {
	fmt.Println("run", s)
}

func main() {
	fmt.Println("====start====")
	// defer 执行顺序和调用顺序相反
	defer print("order 1")
	defer print("order 2")
	defer print("order 3")
	fmt.Println("====end====")
}

// 执行顺序
// ====start====
// ====end====
// run order 1
// run order 2
// run order 3
// tips： 当调用os.Exit()时，defer并不会执行
