package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

// arg 값이 복사되므로 function 에서 변경하여도 값은 그대로
func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func ChangeDataP(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	// arg 값이 복사되므로 function 에서 변경하여도 값은 그대로
	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	// 값 변경
	ChangeDataP(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}

/*
PS C:\Users\조재모\goprojects\example\ex14.2> .\ex14.2.exe
value = 0
data[100] = 0
*/
