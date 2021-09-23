package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	content := flag.String("fcontent", "", "-fcontent指定写入的文件内容")
	flag.Parse()

	// 创建文件
	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 字符串转为[]byte类型
	slice := strings.Split(*content, ";")
	for _, line := range slice {
		// 行写入
		_, err := fmt.Fprintln(f, line)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("写入成功")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
