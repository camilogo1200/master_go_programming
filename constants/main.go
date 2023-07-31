package main

import (
	"fmt"
)

func main() {
	const days = 7
	const pi float64 = 3.1415
	const secondsInHour = 3600

	const n, m int = 4, 5

	const (
		maxN = 500
		minY = 3
	)

	//constant rule
	// 1. you cannot change  a constant

	//2. you cannot initialize  a constant at runtime
	//const power = math.Pow(3, 4)

	//3. you cannot use a variable to initialize a constant
	//t := 5
	//const tc = t

	//IOTA

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(c1, c2, c3)
	fmt.Println(a, b, c)
}
