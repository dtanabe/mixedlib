package mixedlib

// Multiplier is a binary operation that multiplies.
type Multiplier struct {
}

// NewMultiplier creates a new multiplier.
func NewMultiplier() *Multiplier {
	return &Multiplier{}
}

// Int multiplies two ints.
func (Multiplier) Int(x int, y int) int {
	return x * y
}

// Float32 multiplies two float32s.
func (Multiplier) Float32(x float32, y float32) float32 {
	return x * y
}
