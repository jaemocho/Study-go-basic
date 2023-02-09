package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	// rst := square(9)
	// if rst != 81 {
	// 	t.Errorf("square(9) should be 81 but returns %d", rst)
	// }

	assert.Equal(t, 82, square(9), "square(9) should be 81")
}

/*
	PS C:\Users\조재모\goprojects\example\ex28.1> go test
	PASS
	ok      goprojects/ex28.1       0.279s


	PS C:\Users\조재모\goprojects\example\ex28.1> go test
	--- FAIL: TestSquare1 (0.00s)
		ex28.1_test.go:14:
					Error Trace:    C:\Users\조재모\goprojects\example\ex28.1\ex28.1_test.go:14
					Error:          Not equal:
									expected: 82
									actual  : 81
					Test:           TestSquare1
					Messages:       square(9) should be 81
	FAIL
	exit status 1
	FAIL    goprojects/ex28.1       0.315s
*/
