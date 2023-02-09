package main

import "fmt"

func main() {
	var x int8 = 16
	var y int8 = -128
	var z int8 = -1
	var w uint8 = 128

	//범위 넘어가지 않게, 부호에 따라 채워지는 숫자가 다르다 양수면 0 음수면 1
	fmt.Printf("x:%08b x>>2:%08b x>>2: %d\n", x, x>>2, x>>2)
	fmt.Printf("y:%08b y>>2:%08b y>>2: %d\n", y, y>>2, y>>2)
	fmt.Printf("z:%08b z>>2:%08b z>>2: %d\n", z, z>>2, z>>2)
	fmt.Printf("w:%08b w>>2:%08b w>>2: %d\n", w, w>>2, w>>2)
}
