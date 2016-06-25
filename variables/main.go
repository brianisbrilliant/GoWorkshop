package main

import "fmt"

func main() {
	var myName string = "Brian"
	fmt.Println("My name is " + myName + ".")

	myName = "Voldemort"
	fmt.Println("Just kidding my name is " + myName + "!!!")

	// don't have to say what the variable type is
	// this is scary but i need to learn it.
	var myFavNum, aNum = 13, 400
	fmt.Println(aNum, myFavNum)

	var name string = "Butters"
	name := "Butters" // i don't know what := does.aNum

	// for unassigned variables we have zero-values.
	var adam int
	var cartman string
	fmt.Println(adam)    // this will print "0"
	fmt.Println(cartman) // this will print an empty line.
	dumbledore := "dumbledore"
	fmt.Println(dumbledore)
}
