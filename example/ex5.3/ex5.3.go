package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("Exception : ", err)
	} else {
		fmt.Println(n, a, b)
	}

}
