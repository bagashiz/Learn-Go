package main

import "fmt"

func main() {
	//* Pointers are used to store memory addresses

	//* Passed by value
	a := 69
	b := a
	fmt.Println(a, b) // prints 69 69
	a = 420
	fmt.Println(a, b) // prints 420 69

	//* Passed by reference using pointers
	var c int = 69
	//* Use the & operator to get the address of a variable
	var d *int = &c //  / *int is a pointer to an int data type
	fmt.Println(d)  // prints 0xc420018000, because it is the address of c
	//* Use the * operator to dereference a pointer so we can access the value
	fmt.Println(c, *d) // prints 69 69
	*d = 420
	fmt.Println(c, *d) // prints 420 420

	//! Go can not do pointer arithmetic
	//x := [3]int{1, 2, 3}
	//y := &x[0]
	//z := &x[1]
	//! fmt.Println(y + z)
	//! invalid operation: y + z (mismatched types *int and *int)

	//* Pointer objects are created when you assign a value to a variable
	type myStrunct struct {
		foo int
	}
	
	var ms *myStrunct // ms is a pointer to a myStrunct
	ms = &myStrunct{foo: 420} // ms points to a myStrunct with a foo value of 420
	fmt.Println(ms) // prints &{420} because ms is holding the address of a myStrunct object with a foo value of 420

	//* new() function allocates memory for a new object and returns a pointer to it
	var ms2 *myStrunct = new(myStrunct) // ms2 is a pointer to a myStrunct
	fmt.Println(ms2) // prints &{0} because ms2 is holding the address of a myStrunct object with a initial foo value of 0

	//* nil value is a pointer to nothing
	var ms3 *myStrunct
	fmt.Println(ms3) // prints <nil> because ms3 is holding the address of nothing
	//* derefencing a nil pointer will cause a runtime error
	//! fmt.Println(*ms3) //! runtime error: invalid memory address or nil pointer dereference

	//* Assign a value to a pointer
	var ms4 *myStrunct = new(myStrunct)
	fmt.Println(ms4) // prints &{0} because ms4 is holding the address of a myStrunct object with a initial foo value of 0
	(*ms4).foo = 420 // (*ms4) to dereference the pointer and assign the value of foo to 420
	fmt.Println(ms4) // prints &{420} because ms4 is holding the address of a myStrunct object with a foo value of 420
	fmt.Println((*ms4).foo) // prints 420 because ms4 is dereferenced and the foo value is returned
	
	//* But you don't have to dereference a pointer to assign a value to it
	var ms5 *myStrunct = new(myStrunct)
	fmt.Println(ms5) // prints &{0} because ms5 is holding the address of a myStrunct object with a initial foo value of 0
	ms5.foo = 69 // ms5 is dereferenced and the foo value is assigned to 420
	fmt.Println(ms5) // prints &{69} because ms5 is holding the address of a myStrunct object with a foo value of 420
	fmt.Println(ms5.foo) // prints 69 because ms5 is dereferenced and the foo value is returned

	//* NOTE: maps and slices are passed by reference
	//* NOTE: primitives, arrays, and structs are passed by value
}