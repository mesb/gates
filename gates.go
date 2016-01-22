// This package implements basic logic gates.
// It uses Nand as the most basic gate and build all others from it

package gates


import (
	_"math"
	_"fmt"
)

// Bit
type Bit uint8

// Unary function type
type UnaryFunc func(Bit) Bit

// Binary function type
type BinaryFunc func(Bit, Bit) Bit

// n-bit function type
type MultiBitBinFunc func([]Bit, []Bit) []Bit

// n-bit unary function type
type MultiBitUnaryFunc func([]Bit) []Bit

// Construct an n-Bit Binary function
func NewMultiBitBinFunc(nbits int, funcName BinaryFunc) MultiBitBinFunc {

	out := make([]Bit, nbits)
	f := func(a, b []Bit) []Bit {
		for i:= 0; i < nbits-1; i++ {
			out[i] = funcName(a[i], b[i])
		}

		return out
	}

	return f
}


// Constructor for n-bit Unary function
func NewMultiBitUnaryFunc(nbits int, funcName UnaryFunc) MultiBitUnaryFunc {
	out := make([]Bit, nbits)
	f := func(a []Bit) []Bit {
		for i := 0; i < nbits-1; i++ {
			out[i] = funcName(a[i])
		}

		return out

	}

	return f
}

// Nand gate. This is the building block for all other gates
// if a=b=1 then out=0 else out=1
func Nand(a, b Bit) Bit {
	if a == b && a == 1 {
		return 0
	}

	return 1
}

// Not gate
// If in=0 then out=1 else out=0
func Not(in Bit) Bit {
	return Nand(in, in)
}

// And receives 2 input: a, b and returns out
// if a=b=1 then out=1 else out=0
func And(a, b Bit) Bit {
	in := Nand(a, b)

	return Not(in)
}


// Or
// input: a, b
// output: out
// if a=b=0 then out=0 else out=1
func Or(a, b Bit) Bit{
	a = Not(a)
	b = Not(b)
	return Nand(a, b)
}


// Xor
// input: a, b
//output: out
// if a!=b then out=1 else out= 0

func Xor(a, b Bit) Bit{
	i := Nand(a, Not(b))
	j := Nand(Not(a), b)
	return Nand(i, j)
}


// Mux: Multiplexor
// inputs a, b, sel
//outputs: out
// if sel=0 then out=a else out=b
func Mux(a, b, sel Bit) Bit{
	a = And(a, Not(sel))
	b = And(b, sel)
	return Or(a, b)
}


// DMux: Demultuplexor
// inputs in, sel
// outputs: a, b
// Functions: If sel=0 then (a=in, b=0) else (a=0, b=in)
func DMux(in, sel Bit) (a, b Bit) {
	a = And(in, Not(sel))
	b = And(in, sel)

	return a, b
}


// Not16
// Inputs: in[16] // 1 16-bit pin
// out[16]
// for i=0..15 out[i] = Not(in[i])
var Not16 = NewMultiBitUnaryFunc(16, Not)


var And16 = NewMultiBitBinFunc(16, And)


func pow(a, b int) int{
	if b == 0 {
		return 1
	}

	return a * pow(a, b-1)
}


// Convert a binary number to a decimal
func BinToDec(b []Bit) int{
	sum := 0
	n := len(b)

	for i := 0; i < n; i++  {
		p := pow(2, n-1-i)

		sum += (p * int(b[i]))
	}

	return sum
}
