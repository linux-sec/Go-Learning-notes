//package main - 每一个 Go 文件都应该在开头进行 package name 的声明（译注：只有可执行程序的包名应当为 main）。
//包（Packages）用于代码的封装与重用，这里的包名称是main。
package main

//import "fmt" - 引入了 fmt 包，用于在 main 函数里面打印文本到标准输出。
import (
	"fmt"
)

//func main() - main 是一个特殊的函数。
//整个程序就是从 main 函数开始运行的。main 函数必须放置在 main 包中。{ 和 } 分别表示 main 函数的开始和结束部分。
func main() {
	//fmt.Println("Hello World") - fmt 包中的 Println 函数用于把文本写入标准输出。
	fmt.Println("Hello World")
}
