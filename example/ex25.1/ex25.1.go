package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// wait group 생성
	var wg sync.WaitGroup

	// channel 생성
	ch := make(chan int)

	// wait group 1개만
	wg.Add(1)

	//go routine
	go square(&wg, ch)

	// channel에 9를 push (메인 고루틴)
	ch <- 9

	//고루틴이 종료될 때까지 main 고루틴 대기
	wg.Wait()

}

func square(wg *sync.WaitGroup, ch chan int) {
	// channel의 data를 빼온다  메인고루틴에서 서브고루틴으로 데이터 넘겨 줄 때
	n := <-ch

	time.Sleep(time.Second)

	fmt.Println("square:", n*n)
	wg.Done()
}
