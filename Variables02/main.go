package main

import (
	"fmt"
	"math"
)

//变量是什么
//变量指定了某存储单元（Memory Location）的名称，该存储单元会存储特定类型的值。在 Go 中，有多种语法用于声明变量。

//声明单个变量
//var name type 是声明单个变量的语法。
func VarDemoOne() {
	//申明一个 int 类型变量,名字为age
	var age int
	//变量未被赋值，Go 会自动地将其初始化，赋值该变量类型的零值
	fmt.Println("my age is", age)
}

//变量可以赋值给本类型的任何值
func VarDemoTwo() {
	//申明变量
	var age int
	fmt.Println("my age is", age)
	//赋值
	age = 29
	fmt.Println("my age is", age)
	//赋值
	age = 54
	fmt.Println("my age is", age)
}

//声明变量并初始化
//声明变量的同时可以给定初始值。 var name type = initialvalue 的语法用于声明变量并初始化。
func VarDemoThree() {
	//申明变量并初始化
	var age int = 10
	fmt.Println("my age is", age)
}

//类型推断（Type Inference）
//如果变量有初始值，那么 Go 能够自动推断具有初始值的变量的类型。因此，如果变量有初始值，就可以在变量声明中省略 type。
//如果变量声明的语法是 var name = initialvalue，Go 能够根据初始值自动推断变量的类型。
func VarDemoFour() {
	//可以推断类型
	var age = 29
	fmt.Println("my age is", age)
}

//声明多个变量
//Go 能够通过一条语句声明多个变量。
//声明多个变量的语法是 var name1, name2 type = initialvalue1, initialvalue2。
func VarDemoFive() {
	//申明多个变量
	var width, height int = 100, 50
	fmt.Println("width is", width, "height is", height)

	//申明多个不同类型的变量
	var (
		name        = "xiaoming"
		age         = 29
		sex  string = "男"
	)
	fmt.Println("my name is", name, ", age is", age, "and sex is", sex)
}

//简短声明
//Go 也支持一种声明变量的简洁形式，称为简短声明（Short Hand Declaration），该声明使用了 := 操作符。
//声明变量的简短语法是 name := initialvalue。
//简短声明要求 := 操作符左边的所有变量都有初始值，否则程序将会抛出错误。
func VarDemoSix() {
	name, age := "xiaoming", 29
	fmt.Println("my name is", name, ", age is", age)

	//简短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的，否则程序将会抛出错误。
	// 声明变量a和b
	a, b := 20, 30
	fmt.Println("a is", a, "b is", b)
	// b已经声明，但c尚未声明
	b, c := 40, 50
	fmt.Println("b is", b, "c is", c)
	// 给已经声明的变量b和c赋新值
	b, c = 80, 90
	fmt.Println("changed b is", b, "c is", c)
}

//变量可以在运行时进行赋值。
func VarDemoSeven() {
	a, b := 145.8, 543.8
	//求 a 和 b 的最小值
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)
}

//由于 Go 是强类型（Strongly Typed）语言，因此不允许某一类型的变量赋值为其他类型的值。
func VarDemoEight() {
	// age是int类型
	age := 29
	// 错误，尝试赋值一个字符串给int类型变量
	//age = "hello"
	fmt.Println("age is", age)
}

func main() {
	VarDemoOne()
	VarDemoTwo()
	VarDemoThree()
	VarDemoFour()
	VarDemoFive()
	VarDemoSix()
	VarDemoSeven()
	VarDemoEight()
}
