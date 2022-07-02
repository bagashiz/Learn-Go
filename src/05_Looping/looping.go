package main

import "fmt"

func main() {
	//* For loop
	//* For loop syntax:
	// for condition {
	// 	block of code
	// }
	//* OR
	// for initializer; condition; post {
	// 	block of code
	// }

	for i := 0; i < 10; i++ { // i is incremented by 1 each time and stops at 10
		fmt.Println(i)
	} // prints 0 to 9

	//* For loop with multiple initial variables
	for i, j := 0, 1; i < 10; i, j = i+1, j+1 {
		fmt.Println(i, j)
	} // prints 0 1; 1 2; 2 3; 3 4; 4 5; 5 6; 6 7; 7 8; 8 9; 9 10;

	//* Nested loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Println(i, j)
		}
	} // prints 1 1; 1 2; 2 1; 2 2; 3 1; 3 2;

	//* Continue keyword
	for i := 1; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	} // prints 1; 3;

	//* Break keyword
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			if j == 2 {
				break
			}
			fmt.Println(i, j)
		}
	} // prints 1 1; 1 2; 3 1; 3 2;

	//* Labeled loop
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			if i == 2 {
				break outerLoop // breaks out of the outer loop
			}
			fmt.Println(i, j)
		}
	} // prints 1 1; 1 2; 3 1; 3 2;

	//* For loop with range
	//* Syntax:
	// for range variable {
	// 	block of code
	// }

	//* Range returns two values: index and value
	for i, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i, v)
	} // prints 0 1; 1 2; 2 3; 3 4; 4 5;

	//* If you don't need the index, you can use the blank identifier _
	for _, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(v)
	} // prints 1; 2; 3; 4; 5;
}
