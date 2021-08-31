package main

import (
	"fmt"
)

func main() {
	var number int = 200
	if number > 100 {
		fmt.Printf("number=%v", number)
	} else {
		fmt.Printf("number太小了")
	}
}
