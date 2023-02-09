package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)  // 3.5 -> 3
	d := float64(a) * b // int float 라 불가능 타입 변환 필요

	var e int64 = 7
	f := a * int(e) // int int64라 불가능 타입 변환 필요

	fmt.Println(a, b, c, d, e, f)
}
