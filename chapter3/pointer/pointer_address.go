// 地址传参

package main

import (
	"fmt"
	"time"
)

func change(arr [1024]int) {
	// fmt.Printf("%p\n", &arr)

	for i, v := range arr {
		arr[i] = v * 2
	}
}

func changeByAddress(arr *[1024]int) {
	// fmt.Printf("%p\n", &arr)

	for i, v := range arr {
		arr[i] = v * 2
	}
}

func main() {
	arr := [1024]int{}
	for i := 1; i <= 1024; i++ {
		arr[i-1] = i
	}
	// fmt.Printf("%p\n", &arr)

	start := time.Now()
	sum := 0
	for i := 0; i < 10000000; i++ {
		change(arr)
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("change(arr)执行10000000次耗时：", elapsed) // change(arr)执行10000000次耗时： 6.188661112s
	// fmt.Println(arr)

	start = time.Now()
	sum = 0
	for i := 0; i < 10000000; i++ {
		changeByAddress(&arr)
		sum++
	}
	elapsed = time.Since(start)
	fmt.Println("changeByAddress(&arr)执行10000000次耗时：", elapsed) // changeByAddress(&arr)执行10000000次耗时： 4.435039501s
	// fmt.Println(arr)
}
