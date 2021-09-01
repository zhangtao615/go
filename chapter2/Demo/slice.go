package main

import (
	"fmt"
)

func main() {
	var slice_1 []int
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice_1), cap(slice_1), slice_1) // len=0 cap=0 slice=[]
	var slice_2 = []int{1, 2, 3, 4}
	fmt.Printf("slice_2=%v\n", slice_2) // slice_2=[1 2 3 4]
	slice_3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice_3=%d\n", slice_3)           // slice_3=[1 2 3 4 5]
	fmt.Printf("slice_3[0:3]=%d\n", slice_3[0:3]) // slice_3[0:3]=[1 2 3]
	fmt.Printf("slice_3[1:]=%d\n", slice_3[1:])   // slice_3[1:]=[2 3 4 5]
	// 追加
	slice_4 := []int{1, 2, 3}
	slice_4 = append(slice_4, 4, 5)
	fmt.Print(slice_4) // [1 2 3 4 5]
}
