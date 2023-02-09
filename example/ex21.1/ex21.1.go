package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}

/*
PS C:\Users\조재모\goprojects\example\ex21.1> .\ex21.1.exe
nums 타입: []int
15
nums 타입: []int
30
nums 타입: []int
0
PS
*/
