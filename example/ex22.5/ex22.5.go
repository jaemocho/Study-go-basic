package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}

/*

PS C:\Users\조재모\goprojects\example\ex22.5> .\ex22.5.exe
A B C D E
A E D C B

*/
