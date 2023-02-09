package main

import "fmt"

type account struct {
	balance int
}

// function
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// receiver가 있어서 method
func (a *account) withdrawFunc(amount int) {
	a.balance -= amount
}

type myInt int

func (m myInt) Add(a int) myInt {
	rst := int(m) + a // 형변환
	return myInt(rst) // 형변환 리턴
}

func addFunc(m myInt, a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func main() {
	a := &account{100} // *account

	withdrawFunc(a, 30)

	a.withdrawFunc(30)

	fmt.Printf("%d\n", a.balance)

	var b myInt = 3

	b = b.Add(10)

	b = addFunc(b, 10)

	fmt.Printf("%d\n", b)

}

/*
PS C:\Users\조재모\goprojects\example\ex19.1> .\ex19.1.exe
40
*/
