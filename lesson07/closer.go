package lesson07

import "fmt"


func Later() func(string) string {
	vat store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("World"))
	fmt.Println(f("Golang"))
}
