package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}

}

// main 함수 go routine
func main() {
	// PrintHangul go routine 0.3 초 간격 출력
	go PrintHangul()

	// PrintNumbers go routine 0.4 초 간격 출력
	go PrintNumbers()

	// 위의 go routine 들이 종료될 때까지 main 함수 go routine 이 대기
	time.Sleep(3 * time.Second)
}

/*
PS C:\Users\조재모\goprojects\example\ex24.1> .\ex24.1.exe
가1나2다3라마4바5사
PS C:\Users\조재모\goprojects\example\ex24.1>

*/
