package mathhelper

// Untyped constants containing min/max values for (u)int.
// Also define number of bit in int as untyped constant.
// Useful if code can be compiled on architectures with different int size.
// This method should works only for 8, 16, 32, 64 and 128 bits int.

const (
	_max = ^uint(0)                                         // uint max value (constant typed)
	_log = _max>>8&1 + _max>>16&1 + _max>>32&1 + _max>>64&1 // temp

	// UintBytes contains size of uint in bytes as untyped constant
	UintBytes = 1 << _log
	// UintBits contains size of uint in bits as untyped constant
	UintBits = IntBytes << 3

	// MaxUint is uint maximum value as untyped constant
	MaxUint = (1 << IntBits) - 1
	// MinUint is uint minimum value as untyped constant (always 0)
	MinUint = 0

	// IntBytes contains size of int in bytes as untyped constant
	IntBytes = UintBytes
	// IntBits contains size of int in bits as untyped constant
	IntBits = UintBits

	// MaxInt is int maximum value as untyped constant
	MaxInt = MaxUint >> 1
	// MinInt is int minimum value as untyped constant
	MinInt = -MaxInt - 1
)
