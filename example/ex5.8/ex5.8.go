package main

import (
	"bufio"
	"fmt"
	"os"
)

// error 발생 시 입력 버퍼 비우기
func main() {

	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)

		// 개행 문자가 나올 때까지 읽어와라 << 버퍼를 비우는 목적
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	//위에서 n err 선언했으니 사용만
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)

		// 개행 문자가 나올 때까지 읽어와라 << 버퍼를 비우는 목적
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

}
