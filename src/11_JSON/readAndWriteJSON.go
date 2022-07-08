package main

import (
	"encoding/json"
	"log"
)

//* JSON is a subset of JavaScript Object Notation.
//* It is a text file that is human readable and can be parsed by a computer.

//* Create a struct to contain information from a JSON file.
type Person struct {
	// the `json:"..."` tag is used to map the field to the JSON key.
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	//* JSON file
	myJson := `[{
	"firstName": "Clark",
	"lastName": "Kent",
	"hair_color": "Brown",
	"has_dog": true
	},
	
	{
	"firstName": "Bruce",
	"lastName": "Wayne",
	"hair_color": "Black",
	"has_dog": false
	}]`

	//* Read the JSON file.
	//* Create a variable of slice of Person structs to store the JSON data.
	var unmarshalledPeople []Person

	//* Unmarshal the JSON data into the slice of Person structs.
	//* json.Unmarshal() takes a byte slice of JSON data and a pointer to a struct.
	//* The first argument is the JSON data and the second argument is a pointer to the struct.
	err := json.Unmarshal([]byte(myJson), &unmarshalledPeople)

	//* Check for errors.
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
	}

	//* Print the unmarshalled data.
	for _, person := range unmarshalledPeople {
		log.Println(person) // prints {Clark Kent Brown true} and {Bruce Wayne Black false}
	}

	//* Write a JSON file from a struct.
	//* Create a variable of type Person to store the data.
	var mySlc []Person

	//* Create new Person structs.
	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "Red"
	m1.HasDog = false

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "Black"
	m2.HasDog = true

	//* Add the Person structs to the slice.
	mySlc = append(mySlc, m1)
	mySlc = append(mySlc, m2)

	//* Marshal the data into JSON.
	//* json.Marshal() is used to convert a struct into a JSON byte slice.
	//* json.MarshalIndent() is used to format the JSON data with indentation.
	//* The first argument is the data to be marshalled.
	//* The second argument is the prefix for each line.
	//* The third argument is the indentation for each level.
	jsonData, err := json.MarshalIndent(mySlc, "", " ")
	//* Check for errors.
	if err != nil {
		log.Println("Error marshalling JSON:", err)
	}

	//* Print the JSON data.
	log.Println(string(jsonData)) // prints [{"firstName":"Wally","lastName":"West","hair_color":"Red","has_dog":false},{"firstName":"Diana","lastName":"Prince","hair_color":"Black","has_dog":true}]
}
