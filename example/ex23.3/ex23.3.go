package main

import (
	"fmt"
)

type PasswordError struct {
	Len        int
	RequireLen int
}

/*
	type error interface {
	    Error() string
	}

error 의 interface가 위와 같아서 PasswordError struct 에 Error 구현
*/
func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다"
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		// 에러 객체를 만들어서 반환
		return PasswordError{len(password), 8}
	}
	return nil
}

func main() {
	err := RegisterAccount("MyId", "MyPw")
	if err != nil {
		// type 변환,  error type 인지 확인,     ok가 조건문 앞에는 초기문
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입이 됐습니다. ")
	}
}
