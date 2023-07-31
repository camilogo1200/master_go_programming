package main

import (
	"fmt"
	"strconv"
)

func main() {

	//INT TYPE
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65530
	fmt.Printf("%T\n", i2)

	// uint alis for uint32 or uint64
	// int alias for int32 or int64

	//FLOAT TYPE
	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T, %T, %T\n", f1, f2, f3)

	//COMPLEX TYPE
	//complex64, complex128

	//byte alias for uint8
	//rune alias for int32

	//RUNE TYPE
	var r rune = 'f'
	fmt.Printf("%T  %v  %v %v \n", r, r, string(r), strconv.QuoteRune(r))

	//bool type
	var b bool = true
	fmt.Printf("%T\n", b)

	//STRING TYPE
	var s string = "Hi there!"
	fmt.Printf("%T\n", s)

	//ARRAY TYPE
	// has a fixed length, slice has a dynamic length

	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%T \n", numbers)

	//SLICE TYPE
	var cities = []string{"London", "Paris", "New York"}
	fmt.Printf("%T \n", cities)

	//MAP TYPE
	//unordered, indexed by a unique key

	balances := map[string]float64{
		"USD": 4000.2,
		"EUR": 5000.10,
	}
	fmt.Printf("%T\n", balances)

	//STRUCT TYPE
	// can be compared to a class in OOP
	type Car struct {
		brand string
		price int
	}

	var me Car
	fmt.Printf("%T\n", me)

	//POINTER TYPE
	//variable that stores the memory address of another variable
	//value of uninitialized pointer is nil
	var x int = 2
	ptr := &x
	fmt.Printf("ptr of type %T with the value of %v (address) and pointing to %v\n", ptr, ptr, *ptr)
}
