package main

import "fmt"

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

	//* Declaring slices
	//* slice is similiar to array, but no need to specify size/length
	slc := []int{6, 7, 8}
	fmt.Printf("value of slice: %v, %T\n", slc, slc)
	fmt.Printf("Slice length: %v\n", len(slc))

	//* cap(slc) function to return the capacity of slice
	fmt.Printf("Slice capacity: %v\n", cap(slc))

	//* Copying slice in go is passed by reference
	e := []int{1, 2, 3}
	f := e
	f[1] = 5
	fmt.Printf("value of e: %v\n", e) // print [1 5 3] (altered by f)
	fmt.Printf("value of f: %v\n", f) // print [1 5 3]

	//* Create slices from a slice (can work from array too)
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	j := i[:]   // slice of all elements a
	k := i[3:]  // slice from index of 3 to end
	l := i[:6]  //slice from index of 0 to index of 5
	m := i[1:6] // slice from index of 1 to index of 5
	i[1] = 69   // passed by reference
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)

	//* make(type, size, capacity) function to make slice
	slc2 := make([]int, 3)
	fmt.Println(slc2)                             // print [0 0 0]
	fmt.Printf("Slice length: %v\n", len(slc2))   // print 3
	fmt.Printf("Slice capacity: %v\n", cap(slc2)) // print 3
	slc3 := make([]int, 3, 100)
	fmt.Println(slc3)                             // print [0 0 0]
	fmt.Printf("Slice length: %v\n", len(slc3))   // print 3
	fmt.Printf("Slice capacity: %v\n", cap(slc3)) // print 100

	//* append(slice, value1, value2, ...) function to add value into slice
	slc4 := []int{}
	fmt.Println(slc4)                                    // print []
	fmt.Printf("Slice length before: %v\n", len(slc4))   // print 0
	fmt.Printf("Slice capacity before: %v\n", cap(slc4)) // print 0
	slc4 = append(slc4, 1, 2, 3, 4)
	fmt.Println(slc4)                                   // print [1 2 3 4]
	fmt.Printf("Slice length after: %v\n", len(slc4))   // print 4
	fmt.Printf("Slice capacity after: %v\n", cap(slc4)) // print 4
	//* go created a new slice with added value and with length 4 and capacity 4
}
