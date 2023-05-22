package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello")

	//we don't mention the number of values when creating slice
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("The type of fruit list is %T\n", fruitList)

	//we can add as many values as we like, it automatically expands
	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 2
	highScores[1] = 1
	highScores[2] = 4
	highScores[3] = 3

	// highScores[5] = 5 //gives error

	highScores = append(highScores, 9,1,5,23) //no error, reallocate the memory
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove a value from slice based on index
	var courses = []string{"A","B","C","D","E"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index],courses[index+1:]...)

	fmt.Println(courses)
}
