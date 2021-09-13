package entity

import "fmt"

type Student struct {
	Person
	school string
}

func (this *Student) Gotoschool() {
	fmt.Println(this.name, "go to", this.school)
}

func (this *Student) Walk() {
	fmt.Println(this.name, " Walk")
}

func (this *Student) New(name string, age int, sex string, school string) {
	// 继承的属性
	this.name = name
	this.age = age
	this.sex = sex
	// 自己的属性
	this.school = school
}
