package main

import "fmt"

func main() {
	//var slice []int // error

	slice := []int{1, 2, 3}

	// 5번째 2 10번째 3 넣어라
	var slice2 = []int{1, 5: 2, 10: 3}

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10
	fmt.Println(slice, slice2)
}
