package goav

import "testing"

func TestAbc(t *testing.T) {
	k := Minus(2, 1)
	if k != 1 {
		t.Error() // to indicate test failed
	}
}
