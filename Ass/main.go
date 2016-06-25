// Ass = assignment
package main

import (
	"flag"
	"fmt"
	"strconv"
	//string conversiion
)

func main() {
	flag.Parse()
	//name := flag.Arg(0)

	//fmt.Println(name)

	// create a program that will greet the user by name, take their drink order, and get their age
	// if the user orders an alcoholic drink and they are under 21 then report them to the police

	// fuck. args are the same as anywhere else. typed right after the name of the program.
	name := flag.Arg(0)
	ageInput := flag.Arg(1)
	drink := flag.Arg(2)

	age, err := strconv.Atoi(ageInput) // converts string from flag.arg() to a number!

	if err != nil {
		fmt.Println("Age was not a number!")
	}

	if age >= 21 {
		fmt.Println("Hello " + name + ", I'll get your " + drink + " right to you.")
	} else {
		fmt.Println(name + "! It's against the law to serve alcohol to anyone under 21!!!")
		Police()
	}
}
``
func Police() {
	fmt.Println("*siren sounds*")
}

/* Useless first draft because Arg(0) doesn't work like that!!! Fuck!

   	fmt.Println("Welcome to Puzzles, what's your name?")
   	bufio.NewReader(os.Stdin).ReadBytes('\n')
   	custName := flag.Arg(0)

   	fmt.Println("Hello " + custName + ", what year were you born?")
   	bufio.NewReader(os.Stdin).ReadBytes('\n')
   	yearBornInput := flag.Arg(0)
   	yearBorn, err := strconv.Atoi(yearBornInput) // converts string from flag.arg(0) to a number!

   	if err != nil {
   		fmt.Println("Age was not a number!")
   	} else {
   		fmt.Println("Next year you will be older!")
   	}

   	var canDrink = false
   	if (time.Now().Year() - yearBorn) > 21 {
   		fmt.Println("I can get you a drink!")
   		canDrink = true
   	} else {
   		fmt.Println("I'm sorry, it's against the law to serve minors.")
   		Police()
   	}

   	if canDrink {
   		fmt.Println("What can I get you to drink?")
   		bufio.NewReader(os.Stdin).ReadBytes('\n')
   		drinkChoice := flag.Arg(0)
   		fmt.Println("one " + drinkChoice + " coming right up!")
   	}
   }
*/
