package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = 2
	// fmt.Println(x + 2) Ereoe
}
