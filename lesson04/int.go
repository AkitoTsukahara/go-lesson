package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println(i)
	fmt.Println(i + 50)

	var i2 int64 = 200
	// fmt.Println(i + i2) Error

	fmt.Printf("5T\n", i2)

	fmt.Printf("5T\n", int32(i2))

}
