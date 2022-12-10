// The main purpose of this file is to define the public API for mixedlib and to
// make client code imports more ergonomic.

package mixedlib

import "github.com/dtanabe/mixedlib/mixedlib-go/pkg/mixedlib"

type BinaryOperation = mixedlib.BinaryOperation
type Adder = mixedlib.Adder
type Multiplier = mixedlib.Multiplier

var NewAdder = mixedlib.NewAdder
var NewMultiplier = mixedlib.NewMultiplier
