package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕 ! 나는 %d 살 %s라고 해", s.Age, s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer

	// assign type이 다르다 (interface 이기 때문에 가능)
	stringer = student
	fmt.Printf("%s\n", stringer.String())

	// GetAge 는 호출할 수가 없다 interface 에 속하지 않았기 때문에

}
