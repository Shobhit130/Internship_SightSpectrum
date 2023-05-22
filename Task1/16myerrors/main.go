package main

import (
	"errors"
	"fmt"
)

func main() {

	//built-in error
	// for i := 0; i < 5; i++ {
	// 	result := 20 / i
	// 	fmt.Println(result)
	// }

	//create our own errors
	//using New() function
	message := "Hello"

	// create error using New() function
	myError := errors.New("WRONG MESSAGE")

	if message != "Hi" {
		fmt.Println(myError)
	}

	//using Errorf() function
	age := -14

	// create an error using Efforf()
	error := fmt.Errorf("%d is negative\nAge can't be negative", age)

	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Printf("Age: %d", age)
	}

	err := divide(4, 0)

	// error found
	if err != nil {
		fmt.Printf("error: %s", err)

		// error not found
	} else {
		fmt.Println("Valid Division")
	}
}

func divide(num1, num2 int) error {

	// returns error
	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero", num1, num2)
	}

	// returns the result of division
	return nil
}
