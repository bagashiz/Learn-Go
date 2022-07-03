package main

import "fmt"

//* Slices are arrays that can be resized
//* Slices are passed by reference

func main() {
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
