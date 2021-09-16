package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "hello.txt", "-fpath指定文件路径读取")
	byteLen := flag.Int("flen", 6, "-flen指定读取字节数")
	flag.Parse()

	file, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	// 关闭文件，释放资源
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件操作失败:", err)
		}
	}()

	// 文件分片读取到内存中，可以处理大文件
	r := bufio.NewReader(file)
	buffer := make([]byte, *byteLen)
	fmt.Println("文件内容为：")

	for {
		n, err := r.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
