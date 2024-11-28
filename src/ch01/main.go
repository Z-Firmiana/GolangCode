package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		fmt.Println("Hello, World!")
		fmt.Print("Hello, World!") // 无换行
		fmt.Printf("Hello, World!") // 格式化输出

		fmt.Println()

		var a = "aaa"
		fmt.Printf("a = %v, a的类型是%T\n", a, a)
	*/

	// 一次定义多个变量

	/*
		var a, b, c int
		a = 10
		b = 20
		c = 30
		fmt.Println(a, b, c)

		var (
			d int
			e string
		)
		d = 40
		e = "hello"
		fmt.Println(d, e)
	*/

	// 短变量声明，在函数内部，可以使用更简略的：=方式声明并初始化变量；注意：短变量只能用于声明局部变量，不能用于全局变量的声明

	/*
		f, g, h := 90, 100, "hello"
		fmt.Println(f, g, h)
	*/

	// 常量定义

	/*
		const pi = 3.1415926
		fmt.Println(pi)

		const (
			A = "A"
			B = "B"
		)
		fmt.Println(A, B)
	*/

	// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	/*
		const (
			n1 = 100
			n2
			n3
		)
		fmt.Println(n1, n2, n3)
	*/

	// iota 常量计数器
	/*
		const (
			n1 = iota // 0
			n2
			n3
			n4
		)
		fmt.Println(n1, n2, n3, n4)
	*/

	/* const (
		n1 = iota
		n2 = 100
		n3 = iota
		n4
	)
	fmt.Println(n1, n2, n3, n4) */

	// fmt.Spintf()拼接字符串
	/*
		str1 := "hello"
		str2 := "world"
		str3 := fmt.Sprintf("%v!%v.", str1, str2)
		fmt.Println(str3)
	*/

	// strings.Split 分割字符串为切片
	/* str1 := "123-456-789"
	arr := strings.Split(str1, "-")
	fmt.Println(arr) */

	// strings.Join 将切片连接为字符串
	/* str2 := strings.Join(arr, "+")
	fmt.Println(str2) */

	// strings.Index 查找字符串
	str1 := "this is a string"
	str2 := "s"
	index1 := strings.Index(str1, str2)
	index2 := strings.LastIndex(str1, str2)
	fmt.Println(index1)
	fmt.Println(index2)

}
