package lesson08

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "100"

	i, err := strconv.Atoi(s)
	if err := nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}