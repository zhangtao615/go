package entity

import "fmt"

// bmw结构体中实现了Drive方法
type BMW struct {
}

// Drive实现了IDrive接口

func (*BMW) Drive(name string) {
	fmt.Println("Drive BMW" + name + " Car")
}
