package main

import "fmt"

type data interface{}
type Car struct {
	Color string
	Brand string
}

func main() {
	slice := make([]data, 3)
	slice[0] = 1
	slice[1] = "Hello"
	slice[2] = Car{"Red", "BMW"}

	for i, v := range slice {
		if value, ok := v.(int); ok {
			fmt.Printf("slice[%d] type id int[%d]\n", i, value) // slice[0] type id int[1]
		} else if value, ok := v.(string); ok {
			fmt.Printf("slice[%d] type id string[%s]\n", i, value) // slice[1] type id string[Hello]
		} else if value, ok := v.(Car); ok {
			fmt.Printf("slice[%d] type id Car[%s]\n", i, value) // slice[2] type id Car[{Red BMW}]
		} else {
		}
	}
}
