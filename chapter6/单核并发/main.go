package main

import (
	"fmt"
	"runtime"
	"time"

	"./onecore"
)

func main() {
	// GOMAXPROCS 设置同时可执行最大CPU数
	runtime.GOMAXPROCS(1)
	gopher1 := onecore.Gopher{Name: "Gopher1", Id: 1}
	gopher2 := onecore.Gopher{Name: "Gopher2", Id: 2}

	// 非并发
	// gopher1.MakeCoffee("A")
	// gopher2.MakeCoffee("B")

	// 并发
	go gopher1.MakeCoffee("A")
	go gopher2.MakeCoffee("B")

	// 等待防止主协程退出，子协程也将退出
	time.Sleep(20 * time.Second)
	fmt.Println("=================END=================")
}
