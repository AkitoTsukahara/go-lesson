package lesson07

import "fmt"

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		fmt.Println("ABCDEFG")
	})
}
