// Tests for the gates package
package gates

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

// Test for the And gate
func TestAnd(t *testing.T) {
		cases :=  []struct {
	a, b, out Bit
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},

	}

	for _, c := range cases {
		got := And(c.a, c.b)
		if got != c.out {
			t.Errorf("And(%v, %v) = %v", c.a, c.b, c.out)
		}
	}

}




// Test for the Or gate
func TestOr(t *testing.T) {
		cases :=  []struct {
	a, b, out Bit
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 1},

	}

	for _, c := range cases {
		got := Or(c.a, c.b)
		if got != c.out {
			t.Errorf("Or(%v, %v) = %v. Expected %v", c.a, c.b, got, c.out)
		}
	}

}



// Test for the Xor gate
func TestXor(t *testing.T) {
		cases :=  []struct {
	a, b, out Bit
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},

	}

	for _, c := range cases {
		got := Xor(c.a, c.b)
		if got != c.out {
			t.Errorf("Or(%v, %v) = %v. Expected %v", c.a, c.b, got, c.out)
		}
	}

}


// Test Mux
func TestMux(t *testing.T) {
		cases :=  []struct {
			a, b, sel, out Bit
	}{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 1},
		{1, 1, 0, 1},
		{0, 0, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 0},
		{1, 1, 1, 1},

	}

	for _, c := range cases {
		got := Mux(c.a, c.b, c.sel)
		if got != c.out {
			t.Errorf("Mux(%v, %v, %v) = %v. Expected %v", c.a, c.b, c.sel, got, c.out)
		}
	}

}




// Test DMux
func TestDMux(t *testing.T) {
		cases :=  []struct {
			in, sel, a, b Bit
	}{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{1, 0, 1, 0},
		{1, 1, 0, 1},
		//{0, 0, 0, 0},
		//{0, 1, 0, 0},
		//{1, 0, 1, 0},
		//{1, 1, 0, 1},

	}

	for _, c := range cases {
		a, b := DMux(c.in, c.sel)
		if a != c.a && b != c.b {
			t.Errorf("DMux(%v, %v) = %v %v. Expected %v %v", c.in, c.sel, a,
		b, c.a , c.b)
		}
	}

}

// Test for Not16
func TestNot16(t *testing.T) {
	cases := struct {
		in, out []Bit
	}{
		[]Bit{1,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		[]Bit{0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
	}

	got := Not16(cases.in)
	for i, e := range got {
		if got[i] != cases.out[i] {
			t.Errorf("Error computing bit number %v: Not(%v) = %v, expected %v", i,
				e, e, cases.out[i] )
		}
	}
}


// Test the power function
func TestPow(t *testing.T) {
	cases := []struct{
		b, e,  out int
	}{
		{2, 0, 1},
		{3, 0, 1},
		{2, 3, 8},
		{3, 3, 27},
	}

	for _, c := range cases {
		got := pow(c.b, c.e)
		if got != c.out {
			t.Errorf("pow(%v, %v) = %v. Expected %v", c.b, c.e, got, c.out)
		}
	}

}


func TestBinToDec(t *testing.T) {
	cases := []struct{
		n []Bit
		out int
	}{
		{[]Bit{1, 0, 0, 1, 1}, 19},
		{[]Bit{0, 0, 0, 0, 0}, 0},
	}

	for _, c := range cases {
		got := BinToDec(c.n)
		if got != c.out {
			t.Errorf("BinToDec(%v) = %v. Expected %v\n", c.n, got, c.out)
		}
	}
}
