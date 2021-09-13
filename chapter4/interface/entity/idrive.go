package entity

// IDrive接口
// 只要实现Drive(name string) 方法就是实现了接口

type IDrive interface {
	Drive(name string)
}
