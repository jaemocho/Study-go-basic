package main

import (
	"fmt"
	"time"
)

// 무한 반복문 예시
func main() {
	i := 1
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}
