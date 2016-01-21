package gates

// This package implements basic logic gates

// Bit
type Bit uint8


// Nand gate. This is the building block for all other gates
func Nand(a, b Bit) Bit {
	if a == b && a == 1 {
		return 0
	}

	return 1
}

// Not gate
func Not(a Bit) Bit {
	return Nand(a, a)
}
