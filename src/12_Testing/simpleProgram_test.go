package main

import "testing" // testing is a package that is used to test code in go

//* Test file for simpleProgram.go
//* The test file must be in the same directory as the source file
//* The test file must be in the same package as the source file
//* The test file must be in the same directory as the source file
//* The test file must have the same name as the source file with the addition of _test.go
//* Use go test -v command to run the test file

//* Manual test function for divide function
//* 1. Test function for divide function
func TestDivide(t *testing.T) { //* t is a pointer to a testing.T that is a framework for testing in go
	//* calling the function that is being tested
	_, err := divide(10.0, 1.0)
	//* if the error is not nil then the test fails
	if err != nil {
		//* t.Error is used to print error message
		t.Error("Error: ", err)
	}
}

//* 2. Test function for divide function that return error
func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err != nil {
		t.Error("Expected an error, Error: ", err)
	}
}

//* Automated test function for divide function
//* Create a variable with a type a slice of structs that contains the test data
var tests = []struct {
	name           string
	dividend       float32
	divisor        float32
	expectedResult float32
	expectedError  bool
}{
	{"valid", 10.0, 1.0, 10.0, false},
	{"invalid", 10.0, 0.0, 0.0, true},
}

//* 1. Test function for divide function that return error
func TestDivision(t *testing.T) {
	//* Loop through the test data
	for _, tt := range tests {
		//* Call the function that is being tested
		got, err := divide(tt.dividend, tt.divisor)

		//* Check the expected result
		if tt.expectedError {

			//* If the expected error is true then check if the error is nil
			if err == nil {
				t.Errorf("%s: expected error but got result %f", tt.name, got)
			}

		} else {

			//* If the expected error is false then check if the error is not nil
			if err != nil {
				t.Errorf("%s: expected result %f but got error %s", tt.name, tt.expectedResult, err)
			}
		}

		//* If the expected result is not the same as the got result then print the error message
		if got != tt.expectedResult {
			t.Errorf("%s: expected %f but got %f", tt.name, tt.expectedResult, got)
		}
	}
}
