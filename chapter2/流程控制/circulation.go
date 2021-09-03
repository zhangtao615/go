package main

import (
	"fmt"
)

func main() {
	// circularArray()
	// circularSlice()
	// circularMap()
	circularControl()
}

// 循环数组的几种方式

func circularArray() {
	person := [3]string{"Tom", "Jerry", "Spike"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person) // len=3 cap=3 array=[Tom Jerry Spike]

	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
		// person[0]: Tom
		// person[1]: Jerry
		// person[2]: Spike
	}

	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
		// person[0]: Tom
		// person[1]: Jerry
		// person[2]: Spike
	}

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
		// person[0]: Tom
		// person[1]: Jerry
		// person[2]: Spike
	}

	for _, name := range person {
		fmt.Printf("name=%s\n", name)
		// name=Tom
		// name=Jerry
		// name=Spike
	}
}

// 循环切片
func circularSlice() {
	person := []string{"Tom", "Jerry", "Spike"}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(person), cap(person), person) // len=3 cap=3 slice=[Tom Jerry Spike]

	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
		// person[0]: Tom
		// person[1]: Jerry
		// person[2]: Spike
	}

	for i := range person {
		fmt.Printf("person[%d]=%s\n", i, person[i])
		// person[0]=Tom
		// person[1]=Jerry
		// person[2]=Spike
	}

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]=%s\n", i, person[i])
		// person[0]=Tom
		// person[1]=Jerry
		// person[2]=Spike
	}

	for _, name := range person {
		fmt.Printf("name: %s\n", name)
		// name: Tom
		// name: Jerry
		// name: Spike
	}
}

// 循环map

func circularMap() {
	person := map[int]string{
		1: "Tom",
		2: "Jerry",
		3: "Spike",
	}
	fmt.Printf("len=%d map=%v\n", len(person), person) // len=3 map=map[1:Tom 2:Jerry 3:Spike]

	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
		// person[1]: Tom
		// person[2]: Jerry
		// person[3]: Spike
	}

	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
		// person[1]: Tom
		// person[2]: Jerry
		// person[3]: Spike
	}

	for i := 1; i <= len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
		// person[1]: Tom
		// person[2]: Jerry
		// person[3]: Spike
	}

	for _, name := range person {
		fmt.Println("name :", name)
		// name : Spike
		// name : Tom
		// name : Jerry
	}

}

// 改变循环 break continue goto

// break

func circularControl() {
	// break
	for i := 1; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Printf("i=%d\n", i) // 1 2 3
	}

	// continue 只能用于for循环
	for i := 1; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("i=%d\n", i) // 1, 2, 4
	}

	// goto 改变函数内代码执行顺序，不能跨函数使用。
	for i := 1; i <= 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println("i = ", i)
	}

END:
	fmt.Println("end")

	// i =  1
	// i =  2
	// i =  3
	// i =  4
	// i =  5
	// end
}
