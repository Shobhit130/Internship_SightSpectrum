package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

const (
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits           = "0123456789"
	SpecialChars     = "!@#$%^&*()-_=+,.?/:;{}[]~"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter the password length
	fmt.Print("Enter the password length: ")
	length, err := readInt(reader)
	if err != nil {
		log.Fatal("Invalid input:", err)
	}

	// Prompt the user to choose password complexity
	fmt.Println("Choose password complexity:")
	fmt.Println("1. Lowercase letters")
	fmt.Println("2. Lowercase and uppercase letters")
	fmt.Println("3. Lowercase letters, uppercase letters, and digits")
	fmt.Println("4. Lowercase letters, uppercase letters, digits, and special characters")
	fmt.Print("Enter the complexity option: ")
	complexityOption, err := readInt(reader)
	if err != nil {
		log.Fatal("Invalid input:", err)
	}

	// Generate the allowed characters based on the chosen complexity
	var allowedChars string
	switch complexityOption {
	case 1:
		allowedChars = LowercaseLetters
	case 2:
		allowedChars = LowercaseLetters + UppercaseLetters
	case 3:
		allowedChars = LowercaseLetters + UppercaseLetters + Digits
	case 4:
		allowedChars = LowercaseLetters + UppercaseLetters + Digits + SpecialChars
	default:
		log.Fatal("Invalid complexity option")
	}

	// Generate the password
	password, err := generatePassword(length, allowedChars)
	if err != nil {
		log.Fatal("Failed to generate password:", err)
	}

	fmt.Println("Generated password:", password)
}

func generatePassword(length int, allowedChars string) (string, error) {
	var password strings.Builder

	// Generate random indices to select characters from the allowed characters
	maxIndex := big.NewInt(int64(len(allowedChars)))
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return "", err
		}

		char := allowedChars[index.Int64()]
		password.WriteByte(char)
	}

	return password.String(), nil
}

func readInt(reader *bufio.Reader) (int, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	text = strings.TrimSpace(text)
	var num int
	_, err = fmt.Sscanf(text, "%d", &num)
	if err != nil {
		return 0, fmt.Errorf("invalid input")
	}

	return num, nil
}
