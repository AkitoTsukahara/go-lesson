package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	//暗黙的な定義
	i4 := 400
	fmt.Println(i4)

	// i4 = 450 //ここで明示的な定義をしているので、この後の暗黙的な定義でエラー
	// i4 := 500
}
