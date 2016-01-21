package gates
// Testing the gates package

import (
	"testing"
)



// Test the Nand gate
func TestNand(t *testing.T) {

	cases :=  []struct {
	a, b, out Bit
	}{
		{0, 0, 1},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},

	}

	for _, c := range cases {
		got := Nand(c.a, c.b)
		if got != c.out {
			t.Errorf("Nand(%v, %v) == %v, but out is %v, which is wrong", c.a, c.b, got, c.out)
		}
	}

}


// Test the Not gate
func TestNot(t *testing.T) {
	var cases =  []struct {
	in, out Bit
	}{
		{0, 1},
		{1, 0},
	}

	for _, c := range cases {
		got := Not(c.in)
		if got != c.out {
			t.Errorf("Not(%v) == %v, but out is %v, which is wrong", c.in, got, c.out)
		}
	}

}
