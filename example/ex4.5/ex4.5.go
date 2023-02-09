package main

import "fmt"

//type 변환 주의 사항 overflow
func main() {
	var a int16 = 3456   // a는 int16타입 - 2바이트 정수
	var b int8 = int8(a) // int16 -> int8

	fmt.Println(a, b)
}
