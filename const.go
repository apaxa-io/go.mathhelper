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

const (
	Uint8Bytes = 1
	Uint8Bits  = Uint8Bytes * 8
	Int8Bytes  = Uint8Bytes
	Int8Bits   = Uint8Bits
)

const (
	Uint16Bytes = 2
	Uint16Bits  = Uint16Bytes * 8
	Int16Bytes  = Uint16Bytes
	Int16Bits   = Uint16Bits
)

const (
	Uint32Bytes = 4
	Uint32Bits  = Uint32Bytes * 8
	Int32Bytes  = Uint32Bytes
	Int32Bits   = Uint32Bits
)

const (
	Uint64Bytes = 8
	Uint64Bits  = Uint64Bytes * 8
	Int64Bytes  = Uint64Bytes
	Int64Bits   = Uint64Bits
)
