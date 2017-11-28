package amount_test

import (
	"testing"

	"gitlab.com/tokend/go/amount"
)

var Tests = []struct {
	S string
	I int64
}{
	{"100.0000", 1000000},
	{"100.0001", 1000001},
	{"123.0001", 1230001},
}

func TestParse(t *testing.T) {
	for _, v := range Tests {
		o, err := amount.Parse(v.S)
		if err != nil {
			t.Errorf("Couldn't parse %s: %v+", v.S, err)
			continue
		}

		if o != v.I {
			t.Errorf("%s parsed to %d, not %d", v.S, o, v.I)
		}
	}
}

func TestString(t *testing.T) {
	for _, v := range Tests {
		o := amount.String(v.I)

		if o != v.S {
			t.Errorf("%d stringified to %s, not %s", v.I, o, v.S)
		}
	}
}
