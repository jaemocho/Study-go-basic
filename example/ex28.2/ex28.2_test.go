package main

import "testing"

func BenchmarkFibonacci1(b *testing.B) {
	// b.N 은 function 의 반복횟수 같은 function을 N번 반복
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	// b.N 은 function 의 반복횟수 같은 function을 N번 반복
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}

/*
PS C:\Users\조재모\goprojects\example\ex28.2> go test -bench .
goos: windows
goarch: amd64
pkg: goprojects/ex28.2
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkFibonacci1-4            4575316               266.5 ns/op
PASS
ok      goprojects/ex28.2       1.988s
PS C:\Users\조재모\goprojects\example\ex28.2> go test
testing: warning: no tests to run
PASS
ok      goprojects/ex28.2       0.282s

*/
