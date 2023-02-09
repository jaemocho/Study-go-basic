package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}

	// function 종료 전에 반드시 실행
	defer fmt.Println("반드시 호출 됩니다 ")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다")
	fmt.Println("파일에 hello world 를 씁니다 ")
	fmt.Fprintf(f, "Hello World")

}

/*

defer는 stack에 들어가서 출력순서가 거꾸로

PS C:\Users\조재모\goprojects\example\ex21.2> .\ex21.2.exe
파일에 hello world 를 씁니다
파일을 닫았습니다
반드시 호출 됩니다

*/
