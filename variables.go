package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	var g int = 10
	var h byte = 10

	i := 10.0
	j := "10"

	fmt.Printf("g type:%s\n", reflect.TypeOf(g))
	fmt.Printf("h type:%s\n", reflect.TypeOf(h))
	fmt.Printf("i type:%s\n", reflect.TypeOf(i))
	fmt.Printf("j type:%s\n", reflect.TypeOf(j))
}
