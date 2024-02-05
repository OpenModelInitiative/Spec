package v1

import (
	"fmt"
	"strings"
)

// @see https://huggingface.co/docs/safetensors/index
// @see also: https://github.com/ggerganov/ggml/blob/master/src/ggml-quants.h

// DataKind is the data kind.
const (
	DataKindInvalid DataKind = iota //
	DataKindBool
	DataKindQ2
	DataKindQ3
	DataKindQ4
	DataKindU8
	DataKindI8
	DataKindI16
	DataKindU16
	DataKindBF16
	DataKindF16
	DataKindI32
	DataKindU32
	DataKindF32
	DataKindI64
	DataKindU64
	DataKindF64
)

var (
	// DataKindNames is a map of DataKind names.
	DataKindNames = map[DataKind]string{
		DataKindInvalid: "INVALID",
		DataKindBool:    "BOOL",
		DataKindQ2:      "Q2",
		DataKindQ3:      "Q3",
		DataKindQ4:      "Q4",
		DataKindU8:      "U8",
		DataKindI8:      "I8",
		DataKindI16:     "I16",
		DataKindU16:     "U16",
		DataKindBF16:    "BF16",
		DataKindF16:     "F16",
		DataKindI32:     "I32",
		DataKindU32:     "U32",
		DataKindF32:     "F32",
		DataKindI64:     "I64",
		DataKindU64:     "U64",
		DataKindF64:     "F64",
	}

	// DataKindDescriptions is a map of DataKind descriptions.
	DataKindDescriptions = map[DataKind]string{
		DataKindInvalid: "Invalid data kind.",
		DataKindBool:    "Represents a boolean kind.",
		DataKindQ2:      "Represents a 2-bit quantization.",
		DataKindQ3:      "Represents a 3-bit quantization.",
		DataKindQ4:      "Represents a 4-bit quantization.",
		DataKindU8:      "Represents a 8-bit unsigned integer kind.",
		DataKindI8:      "Represents a 8-bit signed integer kind.",
		DataKindI16:     "Represents a 16-bit signed integer kind.",
		DataKindU16:     "Represents a 16-bit unsigned integer kind.",
		DataKindBF16:    "Represents a 16-bit brain floating point kind.",
		DataKindF16:     "Represents a 16-bit half-precision floating point kind.",
		DataKindI32:     "Represents a 32-bit signed integer kind.",
		DataKindU32:     "Represents a 32-bit unsigned integer kind.",
		DataKindF32:     "Represents a 32-bit floating point kind.",
		DataKindI64:     "Represents a 64-bit signed integer kind.",
		DataKindU64:     "Represents a 64-bit unsigned integer kind.",
		DataKindF64:     "Represents a 64-bit floating point kind.",
	}

	// DataKindInBits is a map of DataKind bits.
	DataKindInBits = map[DataKind]uint8{
		DataKindInvalid: 0,
		DataKindBool:    1,
		DataKindQ2:      2,
		DataKindQ3:      3,
		DataKindQ4:      4,
		DataKindU8:      8,
		DataKindI8:      8,
		DataKindI16:     16,
		DataKindU16:     16,
		DataKindBF16:    16,
		DataKindF16:     16,
		DataKindI32:     32,
		DataKindU32:     32,
		DataKindF32:     32,
		DataKindI64:     64,
		DataKindU64:     64,
		DataKindF64:     64,
	}
)

// String outputs the DataKind as a string.
func (dk DataKind) String() string {
	return DataKindNames[dk]
}

// Description outputs the DataKind description.
func (dk DataKind) Description() string {
	return DataKindDescriptions[dk]
}

// ByteSize returns the DataKind as a byte.
func (dk DataKind) ByteSize() uint8 {
	return dk.BitsSize() / 8
}

// BitsSize returns the DataKind as a bit.
func (dk DataKind) BitsSize() uint8 {
	return DataKindInBits[dk]
}

// MarshalJSON outputs the DataKind as a json.
func (dk DataKind) MarshalJSON() ([]byte, error) {
	if !dk.Valid() {
		return nil, fmt.Errorf("invalid data kind")
	}

	return []byte(`"` + dk.String() + `"`), nil
}

// UnmarshalJSON parses the DataKind from json.
func (dk *DataKind) UnmarshalJSON(data []byte) error {
	dataKind, err := ParseDataKind(string(data))
	if err != nil {
		return err
	}

	*dk = dataKind

	return nil
}

// Valid returns true if the DataKind is valid.
func (dk DataKind) Valid() bool {
	return dk != DataKindInvalid
}

// ParseDataKind parses the DataKind string.
func ParseDataKind(value string) (DataKind, error) {
	v := strings.ToLower(strings.Trim(value, `"`))

	for k, n := range DataKindNames {
		if v == n {
			return k, nil
		}
	}

	return DataKindInvalid, fmt.Errorf("invalid data kind: %s", value)
}
