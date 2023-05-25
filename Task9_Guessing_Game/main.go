package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 10
	target := rand.Intn(10) + 1

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have picked a random number between 1 and 10.")
	fmt.Println("Try to guess it!")

	attempts := 0
	maxAttempts := 3

	// Game loop
	for {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		guessStr := scanner.Text()

		// Convert the user's guess to an integer
		guess, err := strconv.Atoi(guessStr)
		if err != nil {
			log.Println("Invalid input. Please enter a number.")
			continue
		}

		// Compare the guess with the target number
		if guess < target {
			fmt.Println("Higher!")
		} else if guess > target {
			fmt.Println("Lower!")
		} else {
			attempts++
			fmt.Printf("Congratulations! You guessed the number in %d attempts!\n", attempts)
			break
		}

		attempts++

		// Check if maximum attempts reached
		if attempts >= maxAttempts {
			fmt.Println("Sorry, you have reached the maximum number of attempts.")
			fmt.Println("The correct answer was:", target)
			break
		}

		// Ask the player if they want to continue
		fmt.Print("Do you want to continue? (y/n): ")
		scanner.Scan()
		answer := strings.ToLower(scanner.Text())

		if answer != "y" && answer != "yes" {
			fmt.Println("The correct answer was:", target)
			break
		}
	}

	playAgain := askToPlayAgain(scanner)
	if playAgain {
		main()
	} else {
		fmt.Println("Thank you for playing the Guessing Game!")
	}
}

func askToPlayAgain(scanner *bufio.Scanner) bool {
	for {
		fmt.Print("Do you want to play again? (y/n): ")
		scanner.Scan()
		answer := strings.ToLower(scanner.Text())

		if answer == "y" || answer == "yes" {
			return true
		} else if answer == "n" || answer == "no" {
			return false
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
}

