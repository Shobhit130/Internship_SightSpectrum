package main

import "fmt"

func main() {
	fmt.Println("Hello")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[2] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("List is:", fruitList)
	fmt.Printf("The type of fruit list is %T",fruitList)
	var arr1 = [3]int{1, 2, 3} //completely initialized
	arr2 := [5]int{4, 5, 6}    //partially initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(len(arr2)) //gives 5 as we declared the array length as 5 above

	arr3 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr4 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr3))
	fmt.Println(len(arr4))

	arr5 := [...]int{1: 10, 4: 20}

	fmt.Println(arr5)
}
