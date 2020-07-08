package main

import (
	"fmt"
	"math"
	"time"
	"io"
	"strings"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func do_type_test(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// Stringers
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Errors
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	v := Vertex{3, 4}

	fmt.Println("Abs called from Vertex ", v.Abs())

	// Call through pointer so actual value changes with method reciver as pointer
	v.Scale(10)
	fmt.Println("Abs called from Vertex ", v.Abs())

	// Calling method with interface
	var i I

	i = &T{"hello"}
	i.M()

	i = F(math.Pi)
	i.M()

	// Type Assertion
	var i1 interface{} = "hello"
	t, ok := i1.(string)
	fmt.Println("Type assertion ", t, ok)

	t1, ok := i1.(float64)
	fmt.Println("Type assertion ", t1, ok)

	// This will trigger panic as no second argument is given
	// t2 := i1.(float64)
	// fmt.Println("Type assertion ", t2)

	// Type switch test
	do_type_test(21)
	do_type_test("hello")

	// Stringer Functions
	stringer1 := Person{"Arthur Dent", 42}
	stringer2 := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(stringer1, stringer2)

	// Errors
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Reader
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}
