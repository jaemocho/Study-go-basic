package main

import "fmt"

type User struct {
	Name string
	Age  int
}

//function 내의 지역변수는 stack 에 쌓인다
// escape analyzing (탈출 분석) : 탈출하는 변수를 분석해서 heap으로 옮긴다
func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
} // u는 사라진다

func main() {

	// *User       Escape
	userPointer := NewUser("AAA", 23)

	fmt.Println(userPointer)
}
