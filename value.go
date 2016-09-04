package main

import (
	"fmt"
	"math/cmplx"
)

var (
	flag   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	fmt.Println("go" + "lang")

	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	const f = "%T(%v)\n"
	fmt.Printf(f, flag, flag)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
