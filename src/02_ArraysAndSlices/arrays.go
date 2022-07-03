package main

import "fmt"

//* Arrays are a fixed-size list of elements of the same type
//* Arrays are zero-based, meaning the first element is 0, the second is 1, etc.
//* Arrays are homogeneous, meaning all elements in the array are of the same type
//* Arrays are initialized with the zero value for the type of element they contain
//* Arrays are accessed by using the index of the element, starting from 0
//* Arrays are mutable, meaning elements can be added or removed from them
//* Arrays are not comparable, meaning they cannot be compared with == or !=
//* Arrays are not addressable, meaning they cannot be passed by reference
//* Arrays are not iterable, meaning they cannot be used in for loops or range loops
//* Arrays are not resizable, meaning they cannot be resized
//* Arrays are not sliceable, meaning they cannot be used in a slice
//* Arrays are not sortable, meaning they cannot be sorted
//* Arrays are not hashable, meaning they cannot be used as keys in maps
//* Arrays are not dynamically allocated, meaning they cannot be allocated on the heap
//* Arrays are not thread-safe, meaning they cannot be used in concurrent programs
//* Arrays are not garbage collected, meaning they cannot be freed by the garbage collector

func main() {
	//* Declaring  arrays
	//* fixed array with elements
	fixedInt := [3]int{1, 2, 3}

	//* fixed array without element
	var fixedStr [2]string
	//* calling array index of n
	fixedStr[0] = "male"
	fixedStr[1] = "female"

	//* relative array size depends on initiated elements
	relativeSizeArr := [...]float64{69.420, 123.456, 3.14, 420.69}

	//* fixed empty array
	var emptyArr [3]string

	//! invalid use of ...
	//! var emptyArr2 [...]string

	fmt.Printf("number: %v, %T\n", fixedInt, fixedInt)
	fmt.Printf("gender: %v, %T\n", fixedStr, fixedStr)
	fmt.Printf("gender index 0: %v, %T\n", fixedStr[0], fixedStr[0])
	fmt.Printf("float: %v, %T\n", relativeSizeArr, relativeSizeArr)
	fmt.Printf("%v, %T\n", emptyArr, emptyArr)

	//* len(arr) function to return the size of an array
	fmt.Printf("The size of relativeSizeArr is %v\n", len(relativeSizeArr))

	//* 2 dimension array
	var matrixArr [2][2]int
	matrixArr[0] = [2]int{1, 0}
	matrixArr[1] = [2]int{0, 1}
	fmt.Printf("Matrix array: %v\n", matrixArr) // print [[1 0] [01]]

	//* Copying array in go is passed by value, not reference
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Printf("value of a: %v\n", a) // print [1 2 3] (did not altered by b)
	fmt.Printf("value of b: %v\n", b) // print [1 5 3]

	//* to mitigate this, go has a pointer feature
	c := [...]int{4, 5, 6}
	d := &c // value of d pointing at value of c (d and c share the same value)
	d[1] = 2
	fmt.Printf("value of c: %v\n", c) // print [4 2 6] (altered by d)
	fmt.Printf("value of d: %v\n", d) // print &[4 2 6]
}
