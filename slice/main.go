// so far we have learned to do things i already know how to do. how is this better?

package main

import "fmt"

func main() {
	// no need to declare var even! just start with the name and get going!
	slice := make([]string, 3)

	fmt.Println(slice)

	slice[0] = "Brian"
	slice[1] = "Cody"
	slice[2] = "Nicolas"

	fmt.Println(slice)

	slice = append(slice, "Ramsey")

	fmt.Println(slice)

	// slices can be copied
	list := make([]string, len(slice))
	copy(list, slice)

	fmt.Println(list)

	sub := slice[2:5]

	fmt.Println(sub)
}
