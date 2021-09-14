package main

import (
	"fmt"

	"./entity"
)

func main() {
	animal := entity.Factory("duck")
	animal.SingGua()                          // Duck SingGua
	fmt.Printf("animal is %s", animal.Type()) // animal is Duck

	animal = entity.Factory("goose")
	fmt.Println()
	animal.SingGua()                          //Goose SingGua
	fmt.Printf("animal is %s", animal.Type()) // animal is Goose, Like Duck
}
