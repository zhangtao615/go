// 作用域
package main

import (
	"fmt"
)

// 申明两个全局变量 cname、age
var cname = "Tom"
var age = 20

func printName() {
	// 局部变量 cname
	var cname = "Jerry"
	// 输出局部变量
	fmt.Printf("printName cname ---> %s\n", cname) // printName cname ---> Jerry

	age++
	fmt.Printf("printName age ---> %d\n", age) // printName age ---> 21
}

func main() {
	printName()
	fmt.Printf("global cname ---> %s\n", cname) // global cname ---> Tom
	fmt.Printf("global age ---> %d\a", age)     // global age ---> 21
}

// 从上面可以看出，函数的作用域一般是通过{}来划分，不同作用域中可以定义相同名称的变量，优先调用当前作用域的变量。
