package entity

import "fmt"

// 定义Goose类型结构体并实现IDuck接口
type Goose struct {
	Color string
}

func (this *Goose) Sleep() {
	fmt.Println("Goose Sleep")
}

func (this *Goose) Eat() {
	fmt.Println("Goose Eat")
}

func (this *Goose) SingGua() {
	fmt.Println("Goose SingGua")
}

func (this *Goose) Type() string {
	return "Goose, Like Duck"
}
