package amount_test

import (
	"gitlab.com/swarmfund/go/amount"
	"math"
	"testing"
)

var testsDataU = []struct {
	S string
	I uint64
}{
	{"100.0000", 1000000},
	{"100.0001", 1000001},
	{"123.0001", 1230001},
	{"1844674407370955.1615", math.MaxUint64},
}

func TestUint64Parse(t *testing.T) {
	for _, v := range testsDataU {
		o, err := amount.ParseU(v.S)
		if err != nil {
			t.Errorf("Couldn't parse %s: %v+", v.S, err)
			continue
		}

		if o != v.I {
			t.Errorf("%s parsed to %d, not %d", v.S, o, v.I)
		}
	}
}

func TestUString(t *testing.T) {
	for _, v := range testsDataU {
		o := amount.StringU(v.I)

		if o != v.S {
			t.Errorf("%d stringified to %s, not %s", v.I, o, v.S)
		}
	}
}
