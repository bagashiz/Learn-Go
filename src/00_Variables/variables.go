package main

import (
	"fmt"     // for printing
	"strconv" // for converting string to numeric data types
)

func main() {
	//* Declare variable
	var i int //default to 0
	i = 42
	//* OR
	var j int = 420
	//* OR
	k := 69.42

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	//* Formatted print
	fmt.Printf("%d, %T\n", j, j)
	fmt.Printf("%.2f, %T\n", k, k)
	fmt.Printf("%v\n%s\n%v\n%d\n", actorName, companion, doctorNumber, season)
	// note: %v often gets the job done

	//* Converting variable type
	var m float32
	m = float32(i)
	fmt.Printf("%.2f, %T\n", m, m) // print 42.00 instead 42

	var n int
	n = int(k)
	fmt.Printf("%v, %T\n", n, n) // print 69 instead 69.42

	var o string // default to ""

	//! o = string(i)
	//! fmt.Printf("%v, %T\n", o, o)
	//! print ASCII number 42 (*) if not using strconv package

	o = strconv.Itoa(i)          // Itoa() = int to string
	fmt.Printf("%v, %T\n", o, o) // Print "42" instead 42

	//* Boolean
	var b bool // Default to false
	fmt.Printf("%v, %T\n", b, b)
	var b1 bool = true
	fmt.Printf("%v, %T\n", b1, b1)
	b2 := 1 != 1
	fmt.Printf("%v, %T\n", b2, b2)

	//* Complex
	var c complex64 = 1 + 2i
	//* OR
	var c1 complex64 = complex(2, 5.2)

	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", c1, c1)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))
	fmt.Printf("%v + %v = %v\n", c, c1, c+c1)
}

//* Name convention
var l int = 100 //* camelCase: scoped to the same package
var L int = 100 //* TitleCase: scoped globally

//* Declared in package-level scope
var i int = 12 //! Variable i in inner-most scope gets called

//* Declared in block
var (
	actorName    string = "John Doe"
	companion           = "Budi"
	doctorNumber int    = 3
	season              = 11
)
