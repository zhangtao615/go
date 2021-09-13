package main

import (
	"fmt"

	"./entity"
)

func main() {
	stu := entity.Student{}

	stu.New("jack", 33, "男", "NEAU")
	// stu.name
	fmt.Println(stu)
	stu.Gotoschool()
	// 覆盖Person的Walk方法
	stu.Walk()
	// 继承Person的Eat方法
	stu.Eat()
}
