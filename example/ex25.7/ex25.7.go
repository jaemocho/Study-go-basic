package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)  // car channel 생성
	paintCh := make(chan *Car) // paintCh channel 생성

	fmt.Printf("Start Factory\n")

	wg.Add(3)           // 3개의 고루틴 추가
	go MakeBody(tireCh) // Go 루틴 생성
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait() // 3개의 고루틴이 종료될 때까지 대기
	fmt.Println("Close the factory")
}

func MakeBody(tireCh chan *Car) { // 차체 생산
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick: // 1초에 한번씩 body를 생성해서 전달
			// Make a body
			car := &Car{} // car 생성
			car.Body = "Sports car"
			tireCh <- car // tier  channel로 전달
		case <-after: // 10초 뒤 종료
			close(tireCh) // 10초가 지나면 tire channel 종료
			wg.Done()     // 고루틴 종료
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) { // 바퀴 설치
	for car := range tireCh {
		// Make a body
		time.Sleep(time.Second) // 타이어 다는데 1초가 걸린다고 설정
		car.Tire = "Winter tire"
		paintCh <- car // paint channel로 car 전달
	} // tier channel 이 닫히면 for문 종료
	wg.Done()      //고루틴 종료
	close(paintCh) // paint channel 종료
}

func PaintCar(paintCh chan *Car) { // 도색
	for car := range paintCh {
		// Make a body
		time.Sleep(time.Second)
		car.Color = "Red"                     // 도색
		duration := time.Now().Sub(startTime) // 경과 시간 출력
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done() // 고루틴 종료
}

/*
  첫 차는 3초에 생성

PS C:\Users\조재모\goprojects\example\ex25.7> .\ex25.7.exe
Start Factory
3.02 Complete Car: Sports car Winter tire Red
4.03 Complete Car: Sports car Winter tire Red
5.04 Complete Car: Sports car Winter tire Red
6.06 Complete Car: Sports car Winter tire Red
7.07 Complete Car: Sports car Winter tire Red
8.08 Complete Car: Sports car Winter tire Red
9.10 Complete Car: Sports car Winter tire Red
10.11 Complete Car: Sports car Winter tire Red
11.12 Complete Car: Sports car Winter tire Red
12.14 Complete Car: Sports car Winter tire Red
Close the factory

*/
