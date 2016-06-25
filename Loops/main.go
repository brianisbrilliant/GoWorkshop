// for is the only loop in Go
package main

import "fmt"

func main() {

	fmt.Println("for i lessthan or equal to 3:")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//asdf
	fmt.Println("for j = 0, if j is less than 4, add 1 to j")
	for j := 0; j < 4; j++ {
		fmt.Println(j)
	}

	fmt.Println("for...ever until the break statement!")
	for {
		fmt.Println("Infinite loop detected!!!!")
		break
	}
}
