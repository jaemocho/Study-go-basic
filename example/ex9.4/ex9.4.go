package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	var a int
	fmt.Println("Hello world")

	// 이런 경우에 쇼트서킷 때문에 IncreaseAndReturn 이 호출 안되는 경우가 있으니 잘 생각해서..
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}
}
