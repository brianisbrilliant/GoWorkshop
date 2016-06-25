// Brian Foster
// hello world!
// June 25th, 2016

/*
	to run this program, open the greet folder in the terminal and type "go install"
	then type the name of the folder your project is in, for example "greet" and your project will run.
	"go build" will create an exe! but it's 2.5mb! that's good actually because go takes care of putting EVERYTHING that you need in the file. Go also does cross-compiling with "GOOS=windows GOARCH=amd64 go build" (that's Go OS = windows, 64-bit Architecture)
*/

// 	every go file has to be in a package
package main

import "fmt" // for printing stuff out to the console

type User struct {
	name  string
	email string
	age   int
	badge int
}

// accepts no parenthesis
func main() {
	// i don't wanna type this all the time!
	fmt.Println("Hello world!")
	fmt.Println("How do I turn autocomplete off?!!?!?!?")
	fmt.Println("Two commands at once!") // type one ampersand inbetween your commands. Ex. "go instsall & greet"
	// go is built for efficiency, not for academics. there is one way to do things. it's like C# in that it's explicitly styled.
	// gofunct automatically removes white space and other stuff that's not necessary.
}
