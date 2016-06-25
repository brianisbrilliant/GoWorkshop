package main

import (
	"fmt"
	"time" // real useful to know what time it is!
)

func main() {
	i := 2
	fmt.Println(i)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend!")
	default:
		fmt.Println("it's a weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}
