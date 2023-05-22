package main

import "fmt"

const LoginToken string = "jhsdgfew" //created variable with first letter as capital 'L', this means that it is now a public variable 

func main() {
	var username string = "Shobhit"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n",username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n",isLoggedIn)

	var smallVal uint8 = 255 //can accept values between 0 to 255 only
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n",smallVal)
	
	var floatValue float32 = 255.928937918982 //only 5 values after decimal, but in case of float64 we get more precise value i.e. more digits after decimal
	fmt.Println(floatValue)
	fmt.Printf("Variable is of type %T \n",floatValue)

	//default values and some aliases
	var anotherVariable int //default value assigned is 0, not any garbage value
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T \n",anotherVariable)

	//implicit type
	var website = "xyz.com" //lexer automatically assigns the data type according to the value given
	fmt.Println(website)

	//no var style
	numverOfUsers := 29212 //cannot use this outside a method like global declaration
	fmt.Println(numverOfUsers)

	fmt.Println(LoginToken)
}
