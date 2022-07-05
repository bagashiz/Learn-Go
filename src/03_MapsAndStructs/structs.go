package main

import (
	"fmt"
	"reflect"
)

func main() {
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
		name:   "Bob",
		companions: []string{
			"Alice",
			"Bob",
			"Charlie",
		},
	}
	fmt.Println(doctor1)               // prints {1 Bob [Alice Bob Charlie]}
	fmt.Println(doctor1.number)        // prints 1
	fmt.Println(doctor1.name)          // prints Bob
	fmt.Println(doctor1.companions)    // prints [Alice Bob Charlie]
	fmt.Println(doctor1.companions[1]) // prints Alice

	//* Anonymous struct
	doctor2 := struct {
		name string
		age  int
	}{name: "Andy", age: 20}
	fmt.Println(doctor2) // prints {Andy 20}

	//* Structs are value types (passed by value)
	copyDoctor := doctor1
	copyDoctor.number = 2
	fmt.Println(doctor1)    // prints {1 Bob [Alice Bob Charlie]}
	fmt.Println(copyDoctor) // prints {2 Bob [Alice Bob Charlie]}
	//* Use pointer to struct to modify struct
	copyDoctor2 := &doctor1
	copyDoctor2.name = "Zoe"
	fmt.Println(doctor1) // prints {1 Zoe [Alice Bob Charlie]}

	//* Embedded structs (useful for imitating inheritance)
	type animal struct {
		name   string
		origin string
	}

	type dog struct {
		animal  // embed animal struct to dog struct
		isLoyal bool
	}

	shiba := dog{}
	shiba.name = "Shiba"
	shiba.origin = "Japan"
	shiba.isLoyal = true
	fmt.Println(shiba) // prints {{Shiba Japan} true}

	pitbull := dog{
		animal:  animal{name: "Pitbull", origin: "Brazil"},
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
	Age  int `required_min:"18"` // required_min is a tag
}
