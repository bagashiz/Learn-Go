package main

import (
	"fmt"
)

func main() {
	//* Naming convention
	const myConst int = 1 //* camelCase: not exported
	const MyConst int = 1 //* TitleCase: exported

	//! const wrongConst float64 = math.Sin(60)
	//! cannot set constant to value of a function

	//* Typed Constants
	const typedConst int = 69
	fmt.Printf("%v, %T\n", typedConst, typedConst)

	//* Untyped Constants
	const untypedConst = 420
	fmt.Printf("%v, %T\n", untypedConst, untypedConst) // int data type
	var a float64 = 69.420
	fmt.Printf("%v, %T\n", untypedConst+a, untypedConst+a) // float64 data type
	//* Compiler sees untypedConst as literal 420

	//* Enumerated Constants
	fmt.Printf("%v, %T\n", i, i) // print 0
	fmt.Printf("%v, %T\n", k, k) // print 1
	fmt.Printf("%v, %T\n", j, j) // print 2

	fmt.Printf("%v, %T\n", l, l) // print 0 (reset)
}

//* Package-level constant
const typedConst float64 = 69.420 // constant data type can be override in function-level

//* Declared in block
const (
	//* Enumerated Constants
	i = iota // iota = counter, default 0 then incremented
	j
	k
	// compiler can see constant block pattern, no need to assign value
)

const (
	//* iota value resetted because different const-block
	l = iota
)
