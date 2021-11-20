package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)

	var s string = "100"
	is, _ := strconv.Atoi(s)
	fmt.Println(is)

	var ii int = 200
	ss := strconv.Itoa(ii)
	fmt.Println(ss)

}
