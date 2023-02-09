package main

import (
	"fmt"
	"time"
)

func main() {
	// 공간이 0인 channel
	ch := make(chan int)

	// wait가 없어서 main이 종료되면 같이 종료되어야한다
	// 바로 naver print가 출력되고 종료되어야하지만
	go square()

	// channel에 data를 넣을 수 없어서  square 안에 무한루프가 계속 실행된다
	// ch :=make(chan int, 2) 로 만들면 바로 never print 하고 종료
	ch <- 9

	fmt.Println("never print")
}

func square() {
	for {

		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}

/*
never print 가 출력이 안된다

PS C:\Users\조재모\goprojects\example\ex25.2> .\ex25.2.exe
sleep
sleep
sleep
sleep
*/
