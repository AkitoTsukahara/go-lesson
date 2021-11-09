package main

import "fmt"

func main() {
	var s string = "Hello Golang!"
	fmt.Println(s)

	var si string = "300"
	fmt.Println(si)

	fmt.Println(`
		test
			test
				test
	`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(s[0])
}
