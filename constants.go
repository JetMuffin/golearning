package main

import (
	"fmt"
	"math"
)

const (
	s        string = "string"
	one, two        = 1, 2
)

func main() {
	fmt.Println(s)
	fmt.Println(one)
	fmt.Println(two)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
