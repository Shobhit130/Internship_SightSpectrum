package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var ptr *int

	//default value in ptr will be nil
	fmt.Println("Value of pointer is ",ptr)

	myNumber := 23

	var ptr1 *int = &myNumber

	fmt.Println("Value of pointer is ",ptr1)
	fmt.Println("Value pointed by pointer is ",*ptr1)

	*ptr1 = *ptr1*2
	fmt.Println("New value is ",myNumber)

}
