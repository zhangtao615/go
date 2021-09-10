## 包

在Go语言中，使用包来封装和隔离不同的功能，可以更好的复用代码。

包命名规范：
1. 包名应该全部小写，不建议包含大写或下划线。
2. 包名应该简短，且能代表该包的主要问题。
3. 包名不用复数形式，例如net、url， 不是nets、urls。
4. 包名不要和标准库中的包名重复，防止导入时需要重命名。

## Go中声明语法
```
package 包名
```
Go语言的源代码存在一个根目录中，包含一个或多个.go文件，每一个.go文件的package声明前面需要用一段文字对该文件的作用进行描述。
```
// main包
package main

import "fmt"

func main() {
	fmt.Println("main包")
}

// demo包
package demo

import "fmt"

func main() {
	fmt.Println("demo包")
}

```
每个包都应该在单独的目录里，不能将多个包放到同一个目录，也不能将一个中的文件分散到不同目录里。

### 创建一个包


