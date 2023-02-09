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
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {

	// 대기 하면서 ch에 데이터가 오길 기다린다
	for n := range ch {
		fmt.Println("Square:", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

/*
PS C:\Users\조재모\goprojects\example\ex25.3> .\ex25.3.exe
Square: 0
Square: 4
Square: 16
Square: 36
Square: 64
Square: 100
Square: 144
Square: 196
Square: 256
Square: 324

main 고루틴은 	wg.Wait() 서 기다리고
서브 고루틴은 for n := range ch { 에서 무한이 대기하기 때문에
wg.Done() 을 수행하지 못하고 데드락이 발생
해결을 위해선 close(ch) 를 wg.Wait() 앞 for문이 끝나는 지점 for i := 0; i < 10; i++ { 에 넣어주면된다

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc00007ff58?)
        C:/Program Files/Go/src/runtime/sema.go:62 +0x25
sync.(*WaitGroup).Wait(0x60?)
        C:/Program Files/Go/src/sync/waitgroup.go:139 +0x52
main.main()
        C:/Users/조재모/goprojects/example/ex25.3/ex25.3.go:20 +0xdd

goroutine 6 [chan receive]:
main.square(0x0?, 0x0?)
        C:/Users/조재모/goprojects/example/ex25.3/ex25.3.go:24 +0xac
created by main.main
        C:/Users/조재모/goprojects/example/ex25.3/ex25.3.go:15 +0xa5
PS C:\Users\조재모\goprojects\example\ex25.3>
*/
