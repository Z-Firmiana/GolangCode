package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
	for index, arg := range os.Args{
		fmt.Println("index:",index, "arg:", arg)
	}
}
