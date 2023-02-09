package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	// 취소기능이 추가된 컨텍스트를 생성
	// context.Backgroup() : 기본 컨텍스트
	// 컨텍스트 생성 시 기본컨텍스트 위에 덮어 쓰게되어있다
	// cancel : function type 이 retrun
	ctx, cancel := context.WithCancel(context.Background())

	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)

	cancel() // ctx 선언 시 return 받았던 cancel function, ctx.Done() 이 발생

	wg.Wait()

}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		// main 함수의 cancel() 이호출되면 동작하는 case
		case <-ctx.Done():
			wg.Done() // 고루틴 종료
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
