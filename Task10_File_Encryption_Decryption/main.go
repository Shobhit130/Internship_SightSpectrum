package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func encrypt(plaintext string, key int) string {
	ciphertext := ""
	for _, char := range plaintext {
		// Encrypt uppercase letters
		if char >= 'A' && char <= 'Z' {
			encryptedChar := 'A' + ((char - 'A' + rune(key)) % 26)
			ciphertext += string(encryptedChar)
		}
		// Encrypt lowercase letters
		if char >= 'a' && char <= 'z' {
			encryptedChar := 'a' + ((char - 'a' + rune(key)) % 26)
			ciphertext += string(encryptedChar)
		}
		// Ignore non-alphabetic characters
		if char < 'A' || (char > 'Z' && char < 'a') || char > 'z' {
			ciphertext += string(char)
		}
	}
	return ciphertext
}

func decrypt(ciphertext string, key int) string {
	plaintext := ""
	for _, char := range ciphertext {
		// Decrypt uppercase letters
		if char >= 'A' && char <= 'Z' {
			decryptedChar := 'A' + ((char - 'A' - rune(key) + 26) % 26)
			plaintext += string(decryptedChar)
		}
		// Decrypt lowercase letters
		if char >= 'a' && char <= 'z' {
			decryptedChar := 'a' + ((char - 'a' - rune(key) + 26) % 26)
			plaintext += string(decryptedChar)
		}
		// Ignore non-alphabetic characters
		if char < 'A' || (char > 'Z' && char < 'a') || char > 'z' {
			plaintext += string(char)
		}
	}
	return plaintext
}

func encryptFile(inputFile, outputFile string, key int) error {
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		plaintext := string(line)
		ciphertext := encrypt(plaintext, key)

		_, err = writer.WriteString(ciphertext + "\n")
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func decryptFile(inputFile, outputFile string, key int) error {
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		ciphertext := string(line)
		plaintext := decrypt(ciphertext, key)

		_, err = writer.WriteString(plaintext + "\n")
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("File Encryption/Decryption Program")
	fmt.Println("==================================")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Encrypt a file")
		fmt.Println("2. Decrypt a file")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		switch choice {
		case 1:
			var inputFile, outputFile string
			var key int

			fmt.Print("Enter the input file path: ")
			fmt.Scanln(&inputFile)

			fmt.Print("Enter the output file path: ")
			fmt.Scanln(&outputFile)

			fmt.Print("Enter the encryption key (a number between 1 and 25): ")
			fmt.Scanln(&key)

			err := encryptFile(inputFile, outputFile, key)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("File encrypted successfully.")
			}

		case 2:
			var inputFile, outputFile string
			var key int

			fmt.Print("Enter the input file path: ")
			fmt.Scanln(&inputFile)

			fmt.Print("Enter the output file path: ")
			fmt.Scanln(&outputFile)

			fmt.Print("Enter the decryption key (a number between 1 and 25): ")
			fmt.Scanln(&key)

			err := decryptFile(inputFile, outputFile, key)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("File decrypted successfully.")
			}

		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
