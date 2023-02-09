package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)

	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	// 같은 key 로 지정하면 데이터 변경
	m["최번개"] = "청주시 덕진구"

	fmt.Printf("송하나의 주소는 %s 입니다.\n", m["송하나"])
	fmt.Printf("백두산의 주소는 %s 입니다.\n", m["백두산"])

	for k, v := range m {
		fmt.Println(k, v)
	}

}
