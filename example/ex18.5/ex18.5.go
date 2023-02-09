package main

import "fmt"

func addNum(slice []int) {
	slice = append(slice, 4)
}

func main() {
	slice := []int{1, 2, 3}

	addNum(slice)
	fmt.Println(slice)
}

// 흔히 하는 실수
/*

PS C:\Users\조재모\goprojects\example\ex18.5> .\ex18.5.exe
[1 2 3]

같은 주소를 가르키고 있긴 하지만
sliceHeader struct가 복사된 것이라
addNum에 넣었던 인자와 addNum 안의 slice는 다르다
변화 없음

아래와 같이 넘기면 변경 가능

func addNum(slice *[]int) {
	*slice = append(*slice, 4)
}


func main() {
	slice := []int{1, 2, 3}

	addNum(&slice)
	fmt.Println(slice)
}

혹은 return 값으로( slice는 아래 방법을 추천 )

func addNum(slice []int) []int {
	slice = append(slice, 4)
	return slice
}

func main() {
	slice := []int{1, 2, 3}
	slice = addNum(slice)
	fmt.Println(slice)
}


*/
