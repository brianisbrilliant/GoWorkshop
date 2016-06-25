package main

import "fmt"

const aConstantString string = "this is a constant"

func main() {
	fmt.Println(aConstantString)

	const c = 500000

	fmt.Println(c)

	const e = 12345 / 9876543

	fmt.Println(e)

	var one int = c
	var two float64 = c
	fmt.Println(one, two)

}
