package main

import "fmt"

// "strings"

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
	/* str1 := "this is a string"
	str2 := "s"
	index1 := strings.Index(str1, str2)
	index2 := strings.LastIndex(str1, str2)
	fmt.Println(index1)
	fmt.Println(index2) */

	/* age := 30;
	if age > 20 {
		fmt.Println("age is greater than 20")
	}
	fmt.Println(age)
	if age1 := 20; age1 > 10 {
		fmt.Printf("age1 is greater than 10")
	} */

	/* for i := 0; i < 10; i++ {
		fmt.Println(i+1)
	} */

	/* sum := 0
	count := 0
	for i := 1; i <= 100; i++ {
		if i%9==0 {
			sum += i
			count++
		}
	}
	fmt.Println(sum)
	fmt.Println(count) */

	/* for i := 1; i <= 16; i++ {
		fmt.Printf("*")
		if i%4==0 {
			fmt.Println()
		}
	} */

	/* n := 5

	switch n {
	case 1,3,5,7,9:
		fmt.Println("奇数")
	case 2,4,6,8,10:
		fmt.Println("偶数")
	default:
		fmt.Println("data is wrong")
	} */

	/* n := 27
		if n > 24 {
			fmt.Println("Adult")
			goto label2
		}

		fmt.Println("A")
		fmt.Println("B")
	label2:
		fmt.Println("C")
		fmt.Println("D") */

	/* arr := [...]int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Printf("%v", v)
	} */

	var sliceA []int
	var sliceB = []int{2, 4, 5}
	sliceA = append(sliceA, 12)
	fmt.Println(sliceA)
	sliceA = append(sliceA, sliceB...)
	fmt.Println(sliceA)
	fmt.Println(sliceB)
}
