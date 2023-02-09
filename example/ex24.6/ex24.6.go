package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과 %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}

/*

S C:\Users\조재모\goprojects\example\ex24.6> .\ex24.6.exe
9 작업시작
4 작업시작
6 작업시작
1 작업시작
5 작업시작
2 작업시작
3 작업시작
7 작업시작
8 작업시작
0 작업시작
0 작업 완료 - 결과 0
5 작업 완료 - 결과 25
9 작업 완료 - 결과 81
4 작업 완료 - 결과 16
6 작업 완료 - 결과 36
1 작업 완료 - 결과 1
3 작업 완료 - 결과 9
2 작업 완료 - 결과 4
7 작업 완료 - 결과 49
8 작업 완료 - 결과 64

*/
