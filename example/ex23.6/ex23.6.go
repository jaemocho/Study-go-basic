package main

import (
	"fmt"
)

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 - ", r)
		}
	}()

	g()

	fmt.Println("f() 함수 끝")

}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨 ")
}

/*
PS C:\Users\조재모\goprojects\example\ex23.6> .\ex23.6.exe
f() 함수 시작
9/3 = 3
panic 복구 -  제수는 0일 수 없습니다.
프로그램이 계속 실행됨
PS C:\Users\조재모\goprojects\

*/
