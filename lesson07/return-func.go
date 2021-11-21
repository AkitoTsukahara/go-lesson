package lesson07

import "fmt"

func ReturnFunc() func() {
	return func() {
		fmt.Println("ABCDEFG")
	}
}

func main() {
	f := ReturnFunc() // ReturnFuncの戻り値のfuncが代入される
	f()
}
