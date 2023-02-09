package main

import (
	"fmt"
	"unsafe"
)

type Stringer interface {
	String() string
}

type Student struct {
	Name string
}

type User struct {
	Name string
	Age  int
}

// pointer type
func (s *Student) String() string {
	return s.Name
}

// value type
func (u User) String() string {
	return u.Name
}

func main() {

	var stringer1 Stringer
	fmt.Printf("type:%T size:%d\n", stringer1, unsafe.Sizeof(stringer1))

	student := &Student{"AAA"}
	stringer1 = student
	fmt.Printf("type:%T size:%d\n", stringer1, unsafe.Sizeof(stringer1))

	var stringer2 Stringer
	fmt.Printf("type:%T size:%d\n", stringer2, unsafe.Sizeof(stringer2))

	user := User{"BBB", 20}
	stringer2 = user
	fmt.Printf("type:%T size:%d\n", stringer2, unsafe.Sizeof(stringer2))

}

/*
PS C:\Users\조재모\goprojects\example\ex20.2> .\ex20.2.exe
type:<nil> size:16
type:*main.Student size:16
type:<nil> size:16
type:main.User size:16

*/
