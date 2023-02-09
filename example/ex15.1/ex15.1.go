package main

import "fmt"

func main() {
	//str[0] = 'H'이지만  str[7] ='드' 가 아니다
	str := "Hello 월드"

	//len() 은 byte 길이를 반환
	for i := 0; i < len(str); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c \n", str[i], str[i], str[i])
	}

	// rune -> int32 별칭 타입
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c \n", arr[i], arr[i], arr[i])
	}

	for _, v := range str {
		fmt.Printf("타입:%T 값:%d 문자값:%c \n", v, v, v)
	}

}

/*
PS C:\Users\조재모\goprojects\example\ex15.1> .\ex15.1.exe
타입:uint8 값:72 문자값:H
타입:uint8 값:101 문자값:e
타입:uint8 값:108 문자값:l
타입:uint8 값:108 문자값:l
타입:uint8 값:111 문자값:o
타입:uint8 값:32 문자값:
타입:uint8 값:236 문자값:ì
타입:uint8 값:155 문자값:타
입:uint8 값:148 문자값:
타입:uint8 값:235 문자값:ë   << 문자가 깨지는 이유는 1byte 씩 읽어와서 ('월' 과 '드' 는 3byte )
타입:uint8 값:147 문자값:
타입:uint8 값:156 문자값:
*/
