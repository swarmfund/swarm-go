package amount

import (
	"math"
	"testing"
)

var Tests = []struct {
	A        int64
	B        int64
	C        int64
	Rounding Rounding
	Result   int64
	Overflow bool
}{
	{3, 200, 6, ROUND_UP, 100, false},
	{3, 200, 6, ROUND_DOWN, 100, false},
	{1, 1, 2, ROUND_UP, 1, false},
	{1, 1, 2, ROUND_DOWN, 0, false},
	{1, 1, 2 * One, ROUND_UP, 1, false},
	{1, 1, 2 * One, ROUND_DOWN, 0, false},

	// A*B overflows
	{math.MaxInt64, 2, math.MaxInt64, ROUND_UP, 2, false},
	{math.MaxInt64, 2, math.MaxInt64, ROUND_DOWN, 2, false},
	{math.MaxInt64, math.MaxInt64, math.MaxInt64, ROUND_DOWN, math.MaxInt64, false},
	{math.MaxInt64, math.MaxInt64, math.MaxInt64, ROUND_UP, math.MaxInt64, false},

	// overflow
	{math.MaxInt64, 4, 3, ROUND_UP, 0, true},
	{math.MaxInt64, 4, 3, ROUND_DOWN, 0, true},
	{math.MaxInt64, 3, 2, ROUND_UP, 0, true},
	{math.MaxInt64, 3, 2, ROUND_DOWN, 0, true},
}

func TestParse(t *testing.T) {
	for _, v := range Tests {
		result, overflow := BigDivide(v.A, v.B, v.C, v.Rounding)
		if overflow != v.Overflow {
			t.Fatalf("Unexpected overflow result. Result: %d, overflow: %t. Data: %+v.", result, overflow, v)
		}

		if overflow {
			continue
		}

		if result != v.Result {
			t.Fatalf("Unexpected result. Result: %d, overflow: %t. Data: %+v.", result, overflow, v)
		}
	}
}
