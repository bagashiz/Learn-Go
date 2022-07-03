package main

import (
	"fmt"
)

func main() {
	//* maps are collections of key-value pairs

	//* Declare maps
	//* map[key_type]value_type
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	} //* map1 is a map of strings to ints
	fmt.Println(map1) // prints map[a:1 b:2 c:3 d:4 e:5]

	//! incomparable maps
	//! errorMap := map[[]int]string {}
	//! fmt.Println(errorMap)
	//! slices, maps, and other functions cannot be a key type in a map

	map2 := map[[3]int]string{
		{1, 2, 3}: "a",
		{2, 3, 4}: "b",
		{3, 4, 5}: "c",
	}
	fmt.Println(map2) // prints map[[1 2 3]:a [2 3 4]:b [3 4 5]:c]

	//* make(map[key_type]value_type) function to make map
	map3 := make(map[string]int)
	map3 = map[string]int{
		"i": 10,
		"j": 20,
		"k": 30,
	}
	fmt.Println(map3) // prints map[i:10 j:20 k:30]

	//* Manipulate value in map
	//* print value of key from a map
	fmt.Println(map1["a"]) // prints 1
	fmt.Println(map1["b"]) // prints 2

	//* Add new key-value pair to map
	map1["F"] = 10
	fmt.Println(map1) // prints map[F:10 a:1 b:2 c:3 d:4 e:5]
	//* Maps are unordered unlike slice

	//* delete(map, key) function to delete key-value pair from map
	delete(map1, "F") // deletes key-value pair with key "F"
	fmt.Println(map1) // prints map[a:1 b:2 c:3 d:4 e:5]

	//* Printing unavaliable key from map returns zero value
	fmt.Println(map1["F"]) // prints 0
	//* Use OK to check if key is in map
	value, ok := map1["F"]
	fmt.Println(value, ok) // prints 0 false
	//* OR
	_, ok = map1["F"]
	fmt.Println(ok) // prints false

	//* len(map) function to get length of map
	fmt.Println(len(map1)) // prints 6
	fmt.Println(len(map2)) // prints 3
	fmt.Println(len(map3)) // prints 3

	//* Maps are reference types (passed by reference)
	copyMap := map1
	copyMap["F"] = 10
	fmt.Println(map1) // prints map[F:10 a:1 b:2 c:3 d:4 e:5]
	delete(copyMap, "F")
	fmt.Println(map1) // prints map[a:1 b:2 c:3 d:4 e:5]
}
