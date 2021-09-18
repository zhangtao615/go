package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	content := flag.String("fcontent", "", "-fcontent指定文件写入内容")
	byteLen := flag.Int("flen", 3, "-flen指定读取字节数")
	flag.Parse()

	f, err := os.Create("hello.txt")
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println("文件操作失败:", err)
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	b := ([]byte)(*content)
	pl := *byteLen

	i := 0
	for {
		if i*pl > len(b) {
			break
		}
		if (i+1)*pl > len(b) {
			n, err := f.Write(b[i*pl : len(b)])
			if err != nil {
				fmt.Println(err)
				break
			}
			if n > 0 {
				fmt.Println("写入字节:", n)
			}
		} else {
			n, err := f.Write(b[i*pl : (i+1)*pl])
			if err != nil {
				fmt.Println(err)
				break
			}
			if n > 0 {
				fmt.Println("写入字节:", n)
			}
		}
		i++
	}
	fmt.Println("写入完毕")
}
