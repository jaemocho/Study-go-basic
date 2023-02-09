package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// 함수 type
func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
		// 함수 리터럴
		// return func(a,b int) int {
		// 	return a+b
		// }
	} else {
		return nil
	}
}

//별칭 타입
type OpFn func(int, int) int

func main() {
	var operator OpFn
	operator = getOperator("+")

	var result = operator(3, 4)
	fmt.Println(result)
}

/*
PS C:\Users\조재모\goprojects\example\ex21.3> .\ex21.3.exe
7
*/
