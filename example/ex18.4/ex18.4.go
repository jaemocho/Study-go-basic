package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)

}

/*

PS C:\Users\조재모\goprojects\example\ex18.4> .\ex18.4.exe
[1 2 3 4 5]
[1 2 200 4 5]

array 는 복사가 되어서 변화가 없고
slice는 sliceHeader struct가 복사된다, 주소값을 가지고 있어서 (ptr type) 변경이 된다
*/
