package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2]
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice[0] = 100 // array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500)
	fmt.Println("After append slice")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}

/*
PS C:\Users\조재모\goprojects\example\ex18.6> .\ex18.6.exe
array: [1 2 3 4 5]
slice: [2] 1 4
After change second element
array: [1 100 3 4 5]
slice: [100] 1 4
After append slice
array: [1 100 500 4 5]
slice: [100 500] 2 4

*/
