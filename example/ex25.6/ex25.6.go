package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {

	// 1초에 한번씩
	tick := time.Tick(time.Second)

	//10초에 딱 한번만
	terminate := time.After(10 * time.Second)

	// 어떤 것을 먼저 실행할지는 random
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return // 10초 지나면 종료함을 의미
		case n := <-ch:
			fmt.Println("Square:", n*n)
			time.Sleep(time.Second)

		}
	}
}

/*

PS C:\Users\조재모\goprojects\example\ex25.6> .\ex25.6.exe
Square: 0
Square: 4
Square: 16
Tick
Square: 36
Square: 64
Square: 100
Square: 144
Tick
Square: 196
Square: 256
Tick
Square: 324
Tick
Square: 0
Tick
Terminated
PS C:\Users\조재모\goprojects\example\ex25.6>

*/
