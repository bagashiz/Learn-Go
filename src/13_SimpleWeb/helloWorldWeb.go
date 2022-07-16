package main

import (
	"fmt"
	"net/http" //* net/http package is used to create a server that responds to HTTP requests
)

func main() {
	//* http.HandleFunc() is a function that takes in a path and a function that takes in a response writer and a request.
	//* "/" means that the function will be called when the user goes to the root of the website.
	//* http.ResponseWriter is a type that represents the response that will be sent back to the user.
	//* http.Request is a type that represents the request that the user has made.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//* fmt.Fprintf() is a function that takes in a response writer and a string.
		//* The string is the text that will be sent back to the user by the server.
		n, err := fmt.Fprintf(w, "Hello World!")
		//* Check if there was an error.
		if err != nil {
			fmt.Println(err)
		}

		//* fmt.Sprintf() is a function that takes in a string and whatever arguments are passed in into a string and returns a string.
		fmt.Println(fmt.Sprintf("%d bytes written to the connection", n))
	})

	//* http.ListenAndServe() is a function that takes in a string and a function that takes in a response writer and a request.
	//* The function is the function that will be called when the user makes a request.
	//* The string is the address of the server.
	//* ":8080" is the port that the server will be listening on.
	//* nil is the error that will be returned if there is an error.
	http.ListenAndServe(":8080", nil)

	//* The server will now be listening on port 8080.
	//* The user can now go to http://localhost:8080/ to see the text "Hello World!".
	//* It will also print out the number of bytes written to the connection on the terminal.
}
