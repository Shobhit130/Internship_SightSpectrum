package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello")

	content := "This needs to go in a file"

	file, err := os.Create("./myfirstfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Printf("Length is : %v\n", length)
	defer file.Close()
	readFile("./myfirstfile.txt")
}

func readFile(fileName string) {
	dataBytes, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	//panic will stop the execution of the program and shows the error that we are facing
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("Data : \n", string(dataBytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
