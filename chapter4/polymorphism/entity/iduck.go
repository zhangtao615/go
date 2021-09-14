package entity

// IDuck类型接口
type IDuck interface {
	Sleep()
	Eat()
	SingGua()
	Type() string
}

// 只需要实现接口中的方法，就是实现了IDuck
