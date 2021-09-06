package main

import (
	"fmt"
)
func main () {
	arr_1 := [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr_1) // [1, 2, 3, 4, 5]
	var arr_2 [3] int
	fmt.Println(arr_2) // [0, 0, 0]
	arr_4 := [...] int {1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr_4)) // 6
	// 二维数组
	var arr_5 = [3][5] int {{ 1, 2, 3, 4, 5}, {3,4,5,6,7}, {}}
	fmt.Println(arr_5) // [[1 2 3 4 5] [3 4 5 6 7] [0 0 0 0 0]]
	arr_6 :=  [...][5] int {{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0:3, 1:5, 4:6}} // [1 2 3 4 5] [9 8 7 6 5] [3 5 0 0 6]]
	fmt.Println(arr_6)
	// 数组赋值
	var arr_7 = [5] int {1, 2, 3, 4, 5}
	arr_7[5] = 6 // invalid array index 5 (out of bounds for 5-element array),数组的长度一旦声明就不可变

	var arr_8 [5] int
	var arr_9 [6] int
	arr_8 = arr_7 // [1, 2, 3, 4, 5]
	arr_9 = arr_7 // cannot use arr (type [5]int) as type [6]int in assignment
}

