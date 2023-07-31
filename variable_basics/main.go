package main

import "fmt"

func main() {
	// declaring variables

	var x int = 7
	var s1 string

	s1 = "Learning Go"

	fmt.Printf("Hi, %v\n", s1)
	fmt.Printf("Hi, %d\n", x)

	s := "Learning"
	fmt.Println(s)

	//multiple declarations

	car, cost := " Audi", 5000

	fmt.Println(car, cost)

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	sum := 8 + 2.53
	fmt.Println(sum)

	golang := "golang"

	// %q for quoted string
	fmt.Printf("%s %q\n", s, golang)

	//%v replaced by any type
	fmt.Printf("%v %v\n", s, golang)

	// %T shows type
	fmt.Printf("%T %T %T\n", s, golang, sum)

	closed := true
	//%t format boolean to print
	fmt.Printf("%T %t\n", closed, closed)

	//%b base 2
	n := 10
	fmt.Printf("n value %d -> %b\n", n, n)

	fmt.Printf("n value %d -> %b - 8bits -> %08b\n", n, n, n)

	//%x base 16
	fmt.Printf("n value %d -> %x \n", n, n)
}
