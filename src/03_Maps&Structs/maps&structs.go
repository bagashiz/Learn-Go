package main

import (
	"fmt"
	"reflect"
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

	//* structs are collections of fields

	//* Declare structs
	//* type Variable_Name struct { field1 type1 field2 type2 }
	type struct1 struct {
		name string
		age  int
	} // Struct1 is a struct with two fields, name and age, both strings and ints
	fmt.Println(struct1{})                     // prints { 0}
	fmt.Println(struct1{"Bob", 20})            // prints {Bob 20}
	fmt.Println(struct1{name: "Bob", age: 20}) // prints {Bob 20}
	fmt.Println(struct1{name: "Bob"})          // prints {Bob 0}
	fmt.Println(struct1{age: 20})              // prints { 20}

	//* Assign struct to variable
	type doctor struct {
		number     int
		name       string
		companions []string
	}
	doctor1 := doctor{
		number: 1,
		name: "Bob", 
		companions: []string{
			"Alice", 
			"Bob", 
			"Charlie",
		},
	}
	fmt.Println(doctor1) // prints {1 Bob [Alice Bob Charlie]}
	fmt.Println(doctor1.number) // prints 1
	fmt.Println(doctor1.name) // prints Bob
	fmt.Println(doctor1.companions) // prints [Alice Bob Charlie]
	fmt.Println(doctor1.companions[1]) // prints Alice

	//* Anonymous struct
	doctor2 := struct {name string; age int}{name: "Andy", age: 20}
	fmt.Println(doctor2) // prints {Andy 20}

	//* Structs are value types (passed by value)
	copyDoctor := doctor1
	copyDoctor.number = 2
	fmt.Println(doctor1) // prints {1 Bob [Alice Bob Charlie]}
	fmt.Println(copyDoctor) // prints {2 Bob [Alice Bob Charlie]}
	//* Use pointer to struct to modify struct
	copyDoctor2 := &doctor1
	copyDoctor2.name = "Zoe"
	fmt.Println(doctor1) // prints {1 Zoe [Alice Bob Charlie]}

	//* Embedded structs (useful for imitating inheritance)
	type animal struct {
		name string
		origin string
	}

	type dog struct {
		animal // embed animal struct to dog struct
		isLoyal bool
	}

	shiba := dog{}
	shiba.name = "Shiba"
	shiba.origin = "Japan"
	shiba.isLoyal = true
	fmt.Println(shiba) // prints {{Shiba Japan} true}
	
	pitbull := dog{
		animal: animal{name: "Pitbull", origin: "Brazil"},
		isLoyal: false,
	}
	fmt.Println(pitbull) // prints {{Pitbull Brazil} false}
	//* note: embedded structs are not actually inherited,
	//* but they are treated as if they were.

	//* Getting struct tag values
	t := reflect.TypeOf(Person{})
	field, _ := t.FieldByName("Age")
	fmt.Println(field.Tag) // prints required_min="18"
}

//* Package-level struct
//* Tags (used to specify field names)
type Person struct {
	Name string
	Age int `required_min:"18"`      // required_min is a tag
}
