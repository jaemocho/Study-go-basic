package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   //4
	Score float64 //12
}

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user))
	/*
		PS C:\Users\조재모\goprojects\example\ex13.5>
		PS C:\Users\조재모\goprojects\example\ex13.5> .\ex13.5.exe
		16
	*/
}

// age 4 byte score 12byte 인데 결과는 12가아닌 16byte => memory padding
// 정렬 및 데이터 가져올 때 편의성을 위해 padding 이 생성
