package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed((time.Now().UnixNano()))

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

/*

PS C:\Users\조재모\goprojects\example\ex24.5> .\ex24.5.exe
B 밥을 먹으려 합니다.
B 수저 획득
B 포크 획득
B 밥을 먹습니다
A 밥을 먹으려 합니다.
A 포크 획득
A 수저 획득
A 밥을 먹습니다
B 밥을 먹으려 합니다.
A 밥을 먹으려 합니다.
A 포크 획득
B 수저 획득
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc000008090?)
        C:/Program Files/Go/src/runtime/sema.go:62 +0x25
sync.(*WaitGroup).Wait(0x1356a9?)
        C:/Program Files/Go/src/sync/waitgroup.go:139 +0x52
main.main()
        C:/Users/조재모/goprojects/example/ex24.5/ex24.5.go:21 +0x171

goroutine 6 [semacquire]:
sync.runtime_SemacquireMutex(0x10?, 0x60?, 0x2?)
        C:/Program Files/Go/src/runtime/sema.go:77 +0x25
sync.(*Mutex).lockSlow(0xc000016090)
        C:/Program Files/Go/src/sync/mutex.go:171 +0x165
sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/sync/mutex.go:90
main.diningProblem({0x1db866, 0x1}, 0xc000016078, 0xc000016090, {0x1dbe5d, 0x6}, {0x1dbe57, 0x6})
        C:/Users/조재모/goprojects/example/ex24.5/ex24.5.go:29 +0x1f1
created by main.main
        C:/Users/조재모/goprojects/example/ex24.5/ex24.5.go:19 +0x10f

goroutine 7 [semacquire]:
sync.runtime_SemacquireMutex(0x10?, 0x60?, 0x2?)
        C:/Program Files/Go/src/runtime/sema.go:77 +0x25
sync.(*Mutex).lockSlow(0xc000016078)
        C:/Program Files/Go/src/sync/mutex.go:171 +0x165
sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/sync/mutex.go:90
main.diningProblem({0x1db867, 0x1}, 0xc000016090, 0xc000016078, {0x1dbe57, 0x6}, {0x1dbe5d, 0x6})
        C:/Users/조재모/goprojects/example/ex24.5/ex24.5.go:29 +0x1f1

*/
