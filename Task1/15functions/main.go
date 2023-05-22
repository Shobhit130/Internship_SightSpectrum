package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	greeter()
	var i, j int

	//Scan takes input even after new line, new line is treated as space
	// fmt.Print("Type two numbers: ")
	// fmt.Scan(&i, &j)

	//Scanln and Scanf does not take input after newline
	// fmt.Print("Type two numbers: ")
	// fmt.Scanln(&i, &j)

	// fmt.Print("Type two numbers: ")
	// fmt.Scanf("%d %d", &i, &j)

	//for taking input through Scanf after newline also 
	fmt.Print("Type two numbers: ")
	fmt.Scanf("%v\n%v", &i, &j)

	result := adder(i, j)
	fmt.Println(result)

	var n int
	fmt.Print("Enter the number of values: ")
	fmt.Scan(&n)

	nums := []int{}

	for i:=0;i<n;i++{
		var ele int
		fmt.Scan(&ele)
		nums = append(nums, ele)
	}

	result1 := adder2(nums)
	fmt.Println(result1)

	result2,resString := adder3(2,5,3,4)
	fmt.Println(resString,result2)

}

func greeter() {
	fmt.Println("Good Morning")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

//when number of values passed are not known
func adder3(values ...int) (int,string){
	total := 0
	for _,val := range values{
		total += val
	}
	return total,"The result of adder 3 is : "
}

func adder2(values []int) int{
	total := 0
	for _,val := range values{
		total += val
	}
	return total
}