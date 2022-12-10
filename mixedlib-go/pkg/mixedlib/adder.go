package mixedlib

// Adder is a binary operation that sums.
type Adder struct {
}

var _ BinaryOperation = (*Adder)(nil)

// NewAdder creates a new adder.
func NewAdder() *Adder {
	return &Adder{}
}

// Int adds two ints.
func (Adder) Int(x int, y int) int {
	return x + y
}

// Float32 adds two float32s.
func (Adder) Float32(x float32, y float32) float32 {
	return x + y
}
