package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"runtime"
	// "math/cmplx"
)

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y  int) {
	x = sum - 4
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

var i bool
var k int = 1

func main() {

	// Basics
	var j int
	var l = "test"

	m := "das"

	defer fmt.Println("Last Line")
	fmt.Println("Hello, World!")
	fmt.Println("Current Time is", time.Now())
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Math Library Rand Int ", rand.Intn(10))
	fmt.Println("Math Library Sqrt ", math.Sqrt(49))
	fmt.Println("Math Library Pi ", math.Pi)
	fmt.Println("Call Add Function ", add(2,2))
	a,b := swap("hello", "world")
	fmt.Println("Call Swap Function ", a,b)
	a1, a2 := split(10)
	fmt.Println("Call Split Function ", a1, a2)
	fmt.Println("Variables Check ", i, j)
	fmt.Println("Variables Check With Values ", k, l, m)
	fmt.Printf("Print Type Variables:  %T %T %T %T %T \n", i, j, k, l, m)

	// Control flow
	sum := 0
	for loop_count := 0; loop_count < 10; loop_count ++ {
		sum += 1
	}

	for sum < 100 {
		sum += sum
	}
	fmt.Println("For loop sum ", sum)

	fmt.Println("Square root function ", sqrt(4), sqrt(-4.54))

	fmt.Println(
		"Calculate math power ",
		pow(2, 3, 20),
		pow(2, 4, 10),
	)

	// Switch Case
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println(time.Now().Weekday())
}