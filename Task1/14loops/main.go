package main

import "fmt"

func main() {
	fmt.Println("Hello")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i, day := range days {
		fmt.Println(i, " ", day)
	}

	var val int = 1

	for val < 10 {
		if val == 5 {
			break
		}
		if val == 2 {
			val += 1
			continue
		}
		if val == 3{
			goto sho
		}
		fmt.Println("Value is: ", val)
		val += 1
	}

	//using goto
	sho:
		fmt.Println("Jumped to the sho label")

	
}
