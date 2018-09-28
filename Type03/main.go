package main

import (
	"fmt"
	"unsafe"
)

//Go支持的基本类型

//bool
//bool 类型表示一个布尔值，值为 true 或者 false。
func TypeDemoOne() {
	//a 赋值为 true
	a := true
	//b 赋值为 false
	b := false
	fmt.Println("a:", a, "b:", b)
	//c 赋值为 a && b。仅当 a 和 b 都为 true 时，操作符 && 才返回 true。因此，在这里 c 为 false。
	c := a && b
	fmt.Println("c:", c)
	//当 a 或者 b 为 true 时，操作符 || 返回 true。在这里，由于 a 为 true，因此 d 也为 true。
	d := a || b
	fmt.Println("d:", d)
}

func TypeDemoTwo() {
	//无符号整型
	//uint8：表示 8 位无符号整型
	//大小：8 位(0-11111111)
	//范围：0～255
	var a uint8
	fmt.Println("a(uint8):", a)
	a = 255
	fmt.Println("a(uint8):", a)

	//uint16：表示 16 位无符号整型
	//大小：16 位(0-1111111111111111)
	//范围：0～65535
	var b uint16
	fmt.Println("b(uint16):", b)
	b = 65535
	fmt.Println("b(uint16):", b)

	//uint32：表示 32 位无符号整型
	//大小：32 位(0-11111111111111111111111111111111)
	//范围：0～4294967295
	var c uint32
	fmt.Println("c(uint32)", c)
	c = 4294967295
	fmt.Println("c(uint32)", c)

	//uint64：表示 64 位无符号整型
	//大小：64 位(0-1111111111111111111111111111111111111111111111111111111111111111)
	//范围：0～18446744073709551615
	var d uint64
	fmt.Println("d(uint64)", d)
	d = 18446744073709551615
	fmt.Println("d(uint64)", d)

	//有符号整型
	//int8：表示 8 位有符号整型
	//大小：8 位
	//范围：-128～127
	var e int8 = -128
	fmt.Println("e(int8)", e)
	e = 127
	fmt.Println("e(int8)", e)

	//int16：表示 16 位有符号整型
	//大小：16 位
	//范围：-32768～32767
	var f int16 = -32768
	fmt.Println("f(int16)", f)
	f = 32767
	fmt.Println("f(int16)", f)

	//int32：表示 32 位有符号整型
	//大小：32 位
	//范围：-2147483648～2147483647
	var g int32 = -2147483648
	fmt.Println("g(int32)", g)
	g = 2147483647
	fmt.Println("g(int32)", g)

	//int64：表示 64 位有符号整型
	//大小：64 位
	//范围：-9223372036854775808～9223372036854775807
	var h int64 = -9223372036854775808
	fmt.Println("h(int64)", h)
	h = 9223372036854775807
	fmt.Println("h(int64)", h)
}

//uint：根据不同的底层平台（Underlying Platform），表示 32 或 64 位无符号整型。
//大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
//范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。
//int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。
//除非对整型的大小有特定的需求，否则通常应该使用 int 表示整型。
//大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
//范围：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807。
func TypeDemoThree() {
	var a int = 89
	b := 95
	// 输出 a 和 b 的值
	fmt.Println("value of a is", a, "and b is", b)
	//在 Printf 方法中，使用 %T 格式说明符（Format Specifier），可以打印出变量的类型。
	//Go 的 unsafe 包提供了一个 Sizeof 函数，该函数接收变量并返回它的字节大小。
	// a 的类型和大小
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	// b 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b))
	//可以看出 a 和 b 为 int 类型，且大小都是 32 位（4 字节）。
	//如果在 64 位系统上运行上面的代码，会有不同的输出。在 64 位系统下，a 和 b 会占用 64 位（8 字节）的大小。
}

//浮点型
//float32：32 位浮点数
//float64：64 位浮点数
func TypeDemoFour() {
	//a 和 b 的类型根据赋值推断得出
	//在这里，a 和 b 的类型为 float64（float64 是浮点数的默认类型）
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	//我们把 a 和 b 的和赋值给变量 sum
	sum := a + b
	//把 b 和 a 的差赋值给 diff
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	//no1 和 no2 也进行了相同的计算
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

//复数类型这里先不说，以后补上

//其他数字类型(byte 和 rune 在字符串时做详细讲解)
func TypeDemoFive() {
	//byte 是 uint8 的别名。
	var a byte
	fmt.Println("a(byte):", a)
	a = 255
	fmt.Println("a(byte):", a)

	//rune 是 int32 的别名。
	var b rune = -2147483648
	fmt.Println("b(rune):", b)
	b = 2147483647
	fmt.Println("b(rune):", b)
}

//string 类型
//在 Golang 中，字符串是字节的集合。
func TypeDemoSix() {
	//first 赋值为字符串 "Naveen"
	first := "Naveen"
	//last 赋值为字符串 "Ramanathan"
	last := "Ramanathan"
	// "+" 操作符可以用于拼接字符串
	//我们拼接了 first、空格和 last，并将其赋值给 name
	name := first + " " + last
	fmt.Println("My name is", name)
}

//类型转换
//Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换。
func TypeDemoSeven() {
	i := 55
	j := 67.8
	sum := i + int(j)
	fmt.Println("sum:", sum)
}

func main() {
	TypeDemoOne()
	TypeDemoTwo()
	TypeDemoThree()
	TypeDemoFour()
	TypeDemoFive()
	TypeDemoSeven()
}
