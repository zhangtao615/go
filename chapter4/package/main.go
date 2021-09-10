package main

import (
	"fmt"

	"go.introduce/chapter4/package/entity"
)

func main() {
	p := entity.Person{}

	p.SetName("Tom")
	p.SetAge(20)
	p.SetSex("男")

	p.Eat()

	fmt.Printf("%+v", p)
}
