package mixedlib

// BinaryOperation is an operation that is defined over the numeric types.
type BinaryOperation interface {
	Int(x int, y int) int
	Float32(x float32, y float32) float32
}
