package main

import "fmt"

// 모든 type 이 가능, 불특정 type을 인자로 받고 싶은 경우 사용
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})
}

/*
PS C:\Users\조재모\goprojects\example\ex20.3> .\ex20.3.exe
v is int 10
v is float64 3.140000
v is string Hello
Not supported type main.Student:{15}

*/
