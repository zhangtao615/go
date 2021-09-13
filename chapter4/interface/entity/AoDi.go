package entity

import "fmt"

type AoDi struct {
}

// Drive实现IDrive接口
func (*AoDi) Drive(name string) {
	fmt.Println("Drive AoDi" + name + " Car")
}
