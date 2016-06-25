// two types of lists - arrays and slices
//arrays first because they're not often useful, but they're very fast.
//arrays have types and sizes

package main

import "fmt"

func main() {
	var array [5]int

	array[0] = 534 // set value of index in array
	array[3] = 666

	fmt.Println(array)

	array[2] = 13
	array[3] = 777
	fmt.Println(array)

	fmt.Println("len: ", len(array)) // length of array
	secondArray := [5]int{9, 8, 7, 6, 5}
	fmt.Println(secondArray)
}
