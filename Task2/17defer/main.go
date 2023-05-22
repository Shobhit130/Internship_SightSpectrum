package main

import "fmt"

func main() {

	//this will execute at last, it follows LIFO
	defer fmt.Println("Hello World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}