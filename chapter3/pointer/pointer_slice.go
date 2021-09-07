package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for _, value := range slice {
		value = 10

		fmt.Println(value) // 10

		fmt.Println(&value) // 所有迭代的value的地址都是一样的: 0xc0000b2008
	}
	fmt.Println(slice) // [1, 2, 3, 4, 5]
}
