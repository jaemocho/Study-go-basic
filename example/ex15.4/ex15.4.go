package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello 월드"
	str2 := str1

	//str1 의 type을 raw pointer  type으로 변환 StringHeader pointer type으로 변환
	stringheader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringheader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringheader1)
	fmt.Println(stringheader2)
}
