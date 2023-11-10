package v1

import (
	"fmt"
	"strings"
)

// @see https://huggingface.co/docs/safetensors/index
// @see also: https://github.com/ggerganov/ggml/blob/master/src/ggml-quants.h

// DType represents a data type of tensor.
const (
	DTypeInvalid DType = iota //
	DTypeBool
	DTypeU8
	DTypeI8
	DTypeI16
	DTypeU16
	DTypeF16
	DTypeBF16
	DTypeI32
	DTypeU32
	DTypeF32
	DTypeI64
	DTypeU64
	DTypeF64
	DTypeQ2K
	DTypeQ3KS
	DTypeQ3KM
	DTypeQ3KL
	DTypeQ40
	DTypeQ41
	DTypeQ4KS
	DTypeQ4KM
	DTypeQ50
	DTypeQ51
	DTypeQ5KS
	DTypeQ5KM
	DTypeQ6KS
	DTypeQ80
)

var (
	// DTypeNames is a map of DType names.
	DTypeNames = map[DType]string{
		DTypeInvalid: "INVALID",
		DTypeBool:    "BOOL",
		DTypeU8:      "U8",
		DTypeI8:      "I8",
		DTypeI16:     "I16",
		DTypeU16:     "U16",
		DTypeF16:     "F16",
		DTypeBF16:    "BF16",
		DTypeI32:     "I32",
		DTypeU32:     "U32",
		DTypeF32:     "F32",
		DTypeI64:     "I64",
		DTypeU64:     "U64",
		DTypeF64:     "F64",
		DTypeQ2K:     "Q2_K",
		DTypeQ3KS:    "Q3_K_S",
		DTypeQ3KM:    "Q3_K_M",
		DTypeQ3KL:    "Q3_K_L",
		DTypeQ40:     "Q4_0",
		DTypeQ41:     "Q4_1",
		DTypeQ4KS:    "Q4_K_S",
		DTypeQ4KM:    "Q4_K_M",
		DTypeQ50:     "Q5_0",
		DTypeQ51:     "Q5_1",
		DTypeQ5KS:    "Q5_K_S",
		DTypeQ5KM:    "Q5_K_M",
		DTypeQ6KS:    "Q6_K_S",
		DTypeQ80:     "Q8_0",
	}

	// DTypeDescriptions is a map of DType descriptions.
	DTypeDescriptions = map[DType]string{
		DTypeInvalid: "Invalid data type.",
		DTypeBool:    "Represents a boolean type.",
		DTypeU8:      "Represents an unsigned byte type.",
		DTypeI8:      "Represents a signed byte type.",
		DTypeI16:     "Represents a 16-bit signed integer type.",
		DTypeU16:     "Represents a 16-bit unsigned integer type.",
		DTypeF16:     "Represents a half-precision (16-bit) floating point type.",
		DTypeBF16:    "Represents a brain (16-bit) floating point type.",
		DTypeI32:     "Represents a 32-bit signed integer type.",
		DTypeU32:     "Represents a 32-bit unsigned integer type.",
		DTypeF32:     "Represents a 32-bit floating point type.",
		DTypeI64:     "Represents a 64-bit signed integer type.",
		DTypeU64:     "Represents a 64-bit unsigned integer type.",
		DTypeF64:     "Represents a 64-bit floating point type.",
		DTypeQ2K:     "Represents a 2-bit quantization.",
		DTypeQ3KS:    "Represents a 3-bit quantization.",
		DTypeQ3KM:    "Represents a 3-bit quantization.",
		DTypeQ3KL:    "Represents a 3-bit quantization.",
		DTypeQ40:     "Represents a 4-bit quantization.",
		DTypeQ41:     "Represents a 4-bit quantization.",
		DTypeQ4KS:    "Represents a 4-bit quantization.",
		DTypeQ4KM:    "Represents a 4-bit quantization.",
		DTypeQ50:     "Represents a 5-bit quantization.",
		DTypeQ51:     "Represents a 5-bit quantization.",
		DTypeQ5KS:    "Represents a 5-bit quantization.",
		DTypeQ5KM:    "Represents a 5-bit quantization.",
		DTypeQ6KS:    "Represents a 6-bit quantization.",
		DTypeQ80:     "Represents a 8-bit quantization.",
	}

	// DTypeBytes is a map of DType bytes.
	DTypeBytes = map[DType]uint8{
		DTypeInvalid: 0,
		DTypeBool:    1,
		DTypeU8:      1,
		DTypeI8:      1,
		DTypeI16:     2,
		DTypeU16:     2,
		DTypeF16:     2,
		DTypeBF16:    2,
		DTypeI32:     4,
		DTypeU32:     4,
		DTypeF32:     4,
		DTypeI64:     8,
		DTypeU64:     8,
		DTypeF64:     8,
		DTypeQ2K:     1,
		DTypeQ3KS:    1,
		DTypeQ3KM:    1,
		DTypeQ3KL:    1,
		DTypeQ40:     1,
		DTypeQ41:     1,
		DTypeQ4KS:    1,
		DTypeQ4KM:    1,
		DTypeQ50:     1,
		DTypeQ51:     1,
		DTypeQ5KS:    1,
		DTypeQ5KM:    1,
		DTypeQ6KS:    1,
		DTypeQ80:     1,
	}

	// DTypeBits is a map of DType bits.
	DTypeBits = map[DType]uint8{
		DTypeInvalid: 0,
		DTypeBool:    1,
		DTypeU8:      8,
		DTypeI8:      8,
		DTypeI16:     16,
		DTypeU16:     16,
		DTypeF16:     16,
		DTypeBF16:    16,
		DTypeI32:     32,
		DTypeU32:     32,
		DTypeF32:     32,
		DTypeI64:     64,
		DTypeU64:     64,
		DTypeF64:     64,
		DTypeQ2K:     2,
		DTypeQ3KS:    3,
		DTypeQ3KM:    3,
		DTypeQ3KL:    3,
		DTypeQ40:     4,
		DTypeQ41:     4,
		DTypeQ4KS:    4,
		DTypeQ4KM:    4,
		DTypeQ50:     5,
		DTypeQ51:     5,
		DTypeQ5KS:    5,
		DTypeQ5KM:    5,
		DTypeQ6KS:    6,
		DTypeQ80:     8,
	}
)

// String outputs the DType as a string.
func (dt DType) String() string {
	return DTypeNames[dt]
}

// Description outputs the DType description.
func (dt DType) Description() string {
	return DTypeDescriptions[dt]
}

// ByteSize returns the DType as a byte.
func (dt DType) ByteSize() uint8 {
	return DTypeBytes[dt]
}

// BitSize returns the DType as a bit.
func (dt DType) BitSize() uint8 {
	return DTypeBits[dt]
}

// MarshalJSON outputs the DType as a json.
func (dt DType) MarshalJSON() ([]byte, error) {
	if !dt.Valid() {
		return nil, fmt.Errorf("invalid dtype")
	}

	return []byte(`"` + dt.String() + `"`), nil
}

// UnmarshalJSON parses the DType from json.
func (dt *DType) UnmarshalJSON(data []byte) error {
	dType, err := ParseDType(string(data))
	if err != nil {
		return err
	}

	*dt = dType

	return nil
}

// Valid returns true if the DType is valid.
func (dt DType) Valid() bool {
	return dt != DTypeInvalid
}

// ParseDType parses the DType string.
func ParseDType(value string) (DType, error) {
	v := strings.ToLower(strings.Trim(value, `"`))

	for k, n := range DTypeNames {
		if v == n {
			return k, nil
		}
	}

	return DTypeInvalid, fmt.Errorf("invalid dtype: %s", value)
}
