package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	var p *int

	fmt.Println("Pointers in Go")
	fmt.Println("Value of uninitialized pointer", p)
	i := 1
	p = &i
	fmt.Println("Value of pointer (Address of i)", p)
	fmt.Println("Address of i (Must be same as value of p)", &i)
	fmt.Println("Value of pointer p pointinng to", *p)
	fmt.Println("Address of pointer p", &p)

	// Structs
	s := Point{4, 2}
	fmt.Println("Value of Struct's X ", s.X)

	s1 := &s
	fmt.Println("Value of Struct's X thorugh Pointer ", s1.X)

	var s2 = Point{X: 1}
	fmt.Println("Value of Struct S2 ", s2.X, s2.Y)

	var s3 = &Point{X: 4}
	fmt.Println("Value of Struct S2 ", s3.X, s3.Y, s3)

	// Arrays
	var a [10]int
	fmt.Println("Uninitialized Array", a)

	a[1] = 1
	a[2] = 2
	fmt.Println("Uninitialized Array", a[1], a)

	odd := [3]int{1, 3, 5}
	fmt.Println("Pre Initialized Array (Odd)", odd)

	// Array Slicing
	fmt.Println("Pre Initialized Array (Odd)", odd[1:3])

	// Array of Structs
	arr_structs := []Point{{1,2}, {2,3}, {3,4}}
	arr_structs1 := []struct{X,Y int}{{2,3}, {3,4}, {4,5}}
	fmt.Println("Array of pre defined structs", arr_structs)
	fmt.Println("Array of structs", arr_structs1)

	// Range Function
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i,v := range pow{
		fmt.Println("Range Function ", i, v)
	}

	// Maps
	var m map[string]Point
	m = make(map[string]Point)
	m["Bell Labs"] = Point{
		40, -74,
	}
	fmt.Println("Maps ", m)
	fmt.Println("Maps ", m["Bell Labs"])

	var m1 = map[string]Point{
		"Bell Labs": Point{
			40, -74,
		},
		"Google": Point{
			37, -122,
		},
	}
	fmt.Println("Maps ", m1)

	var m2= map[string]Point{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
	}
	fmt.Println("Maps ", m2)

	fmt.Println("The value m2 before delete", m2["Google"])

	delete(m2, "Google")
	fmt.Println("The value m2 ", m2["Google"])

	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// Function pass as value to another function
	// Just like python first class functions
	fmt.Println("Function pass to another function ", compute(math.Pow))

	// Closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			"closures",
			pos(i),
			neg(-2*i),
		)
	}
}