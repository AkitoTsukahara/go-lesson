package main

import "fmt"

const Pi = 3.14

const (
	URL      = "https://test.com"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

func main() {
	fmt.Println(Pi)

	// Pi = 3  Error

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)
}
