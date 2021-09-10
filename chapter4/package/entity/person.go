// 通过结构体实现封装

package entity

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

// Set
func (this *Person) SetName(name string) {
	this.name = name
}
func (this *Person) SetAge(age int) {
	this.age = age
}
func (this *Person) SetSex(sex string) {
	if sex == "男" || sex == "女" {
		this.sex = sex
	} else {
		this.sex = "未知"
	}
}

// Get
func (this *Person) GetName() string {
	return this.name
}
func (this *Person) GetAge() int {
	return this.age
}
func (this *Person) GetSex() string {
	return this.sex
}

// 私有方法

func (p *Person) walk() {
	fmt.Println("Person walk")
}

// 公有方法
func (p *Person) eat() {
	fmt.Println("Person eat")
}
