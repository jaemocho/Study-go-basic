package main

import (
	"fmt"
	"os"
)

type Writer func(string)

// 의존성 주입, 어떤 function이 와도 hello world를 쓰겠다
func writeHello(writer Writer) {
	writer("Hello world")
}

func main() {
	f, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
