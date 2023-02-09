package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}

/*
PS C:\Users\조재모\goprojects\example\ex23.5> .\ex23.5.exe
9 / 3 = 3
panic: b는 0일 수 없습니다

goroutine 1 [running]:
main.divide(0x9?, 0x3?)
        C:/Users/조재모/goprojects/example/ex23.5/ex23.5.go:7 +0x105
main.main()
        C:/Users/조재모/goprojects/example/ex23.5/ex23.5.go:14 +0x31

*/
