// 어떤 package에 속하는지 지정
// main 은 특별한 의미, main은 program 시작점을 포함하는 package, 1개만
package main

// fmt package 를 가져오겠다
import "fmt"

// function
// main 은 예약어, program 시작점
func main() {
	fmt.Println("Hello World")
}

// 모듈 생성 (go.mod)
// go mod init goprojects/hello
// 실행 파일 생성 (hello.exe)
// go build

// 실행파일 -> memory로 load -> cpu가 한줄 씩 읽어서 실행 (튜랭, 폰노이만 구조)
