package main

import "fmt"

func main() {

	fmt.Println("if 7 modulus 2 == 0")
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	name := "Brian"
	if name != "" { // if always needs parenthesis, even with single lines
		fmt.Println("your name is ", name)
	} else if name == "tanner" {
		fmt.Println("Hello Tanner!")
	} else {
		fmt.Println("Your name is blank")
	}

	//-------------------
	if age := 21 + 13; age < 44 && name != "jacob" {
		fmt.Println("")

	}
}
