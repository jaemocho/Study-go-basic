package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	// context 를 사용해 값을 지정할 수 있다
	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {

	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("square:%d", n*n)
	}
	wg.Done()
}

/*
PS C:\Users\조재모\goprojects\example\ex25.9> .\ex25.9.exe
square:81
*/
