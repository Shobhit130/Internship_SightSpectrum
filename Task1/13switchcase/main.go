package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1")
		break
	case 2:
		fmt.Println("Dice value is 2")
		break
	case 3:
		fmt.Println("Dice value is 3")
		break
	case 4:
		fmt.Println("Dice value is 4")
		fallthrough //case 5 will also be printed
	case 5:
		fmt.Println("Dice value is 5")
		break
	case 6:
		fmt.Println("Dice value is 6")
		break
	default:
		fmt.Println("Not a valid number")
	}
}
