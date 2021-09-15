package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("hello.txt")

	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	} else {
		// []byte转为string
		fmt.Println("文件内容是: \n", string(content))
	}
}
