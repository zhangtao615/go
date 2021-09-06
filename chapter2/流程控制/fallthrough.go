// 在switch语句中，一旦匹配case就会执行break，不会执行其他的case，如果想执行其他的case可以使用fallthrough
package main

import "fmt"

func getSex(code int) string {
	ret := ""
	switch code {
	case 1:
		ret += "男"
		fallthrough // fallthrough只能放在case分支处理逻辑的最后
	case 0:
		ret += "女"
		fallthrough
	default:
		ret += "未知"
	}
	return ret
}

func main() {
	code := 1
	ret := getSex(code)
	fmt.Printf("%d -> %s \n", code, ret) // 1 -> 男女未知
	code = 0
	ret = getSex(code)
	fmt.Printf("%d -> %s \n", code, ret) // 2 -> 女未知
}
