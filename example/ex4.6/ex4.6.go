package main

import "fmt"

//package 전역 변수
var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// s 는 변수의 범위를 넘어서 error
	m = s + 20
}
