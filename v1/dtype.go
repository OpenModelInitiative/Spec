package v1

import (
	"bytes"
	"strings"
)

const (
	DTypeInvalid    DType = iota //
	DTypeBool                    // boolean
	DTypeInt2                    // A 2-bit signed integer can represent values from -2 to 1.
	DTypeInt3                    // A 3-bit signed integer can represent values from -4 to 3.
	DTypeInt4                    // A 4-bit signed integer can represent values from -8 to 7.
	DTypeInt5                    // A 5-bit signed integer can represent values from -16 to 15.
	DTypeInt6                    // A 6-bit signed integer can represent values from -32 to 31.
	DTypeInt7                    // A 7-bit signed integer can represent values from -64 to 63.
	DTypeInt8                    // A 8-bit signed integer can represent values from -128 to 127.
	DTypeInt16                   // A 16-bit signed integer can represent values from -32768 to 32767.
	DTypeInt32                   // A 32-bit signed integer can represent values from -2147483648 to 2147483647.
	DTypeInt64                   // A 64-bit signed integer can represent values from -9223372036854775808 to 9223372036854775807.
	DTypeUint2                   // A 2-bit unsigned integer can represent values from 0 to 3.
	DTypeUint3                   // A 3-bit unsigned integer can represent values from 0 to 7.
	DTypeUint4                   // A 4-bit unsigned integer can represent values from 0 to 15.
	DTypeUint5                   // A 5-bit unsigned integer can represent values from 0 to 31.
	DTypeUint6                   // A 6-bit unsigned integer can represent values from 0 to 63.
	DTypeUint7                   // A 7-bit unsigned integer can represent values from 0 to 127.
	DTypeUint8                   // A 8-bit unsigned integer can represent values from 0 to 255.
	DTypeUint16                  // A 16-bit unsigned integer can represent values from 0 to 65535.
	DTypeUint32                  // A 32-bit unsigned integer can represent values from 0 to 4294967295.
	DTypeUint64                  // A 64-bit unsigned integer can represent values from 0 to 18446744073709551615.
	DTypeFloat16                 // (half-precision) A 16-bit floating point number.
	DTypeFloat32                 // (single-precision) A 32-bit floating point number.
	DTypeFloat64                 // (double-precision) A 64-bit floating point number.
	DTypeBFloat16                // (Brain Float) A 16-bit floating point number.
	DTypeComplex64               // A 64-bit complex
	DTypeComplex128              // A 128-bit complex
)

var (
	// DTypeNames is a map of DType names.
	DTypeNames = map[DType]string{
		DTypeBool:       "bool",
		DTypeInt2:       "int2",
		DTypeInt3:       "int3",
		DTypeInt4:       "int4",
		DTypeInt5:       "int5",
		DTypeInt6:       "int6",
		DTypeInt7:       "int7",
		DTypeInt8:       "int8",
		DTypeInt16:      "int16",
		DTypeInt32:      "int32",
		DTypeInt64:      "int64",
		DTypeUint2:      "uint2",
		DTypeUint3:      "uint3",
		DTypeUint4:      "uint4",
		DTypeUint5:      "uint5",
		DTypeUint6:      "uint6",
		DTypeUint7:      "uint7",
		DTypeUint8:      "uint8",
		DTypeUint16:     "uint16",
		DTypeUint32:     "uint32",
		DTypeUint64:     "uint64",
		DTypeFloat16:    "float16",
		DTypeFloat32:    "float32",
		DTypeFloat64:    "float64",
		DTypeBFloat16:   "bfloat16",
		DTypeComplex64:  "complex64",
		DTypeComplex128: "complex128",
	}

	// DTypeDescriptions is a map of DType descriptions.
	DTypeDescriptions = map[DType]string{
		DTypeBool:       "boolean",
		DTypeInt2:       "A 2-bit signed integer can represent values from -2 to 1.",
		DTypeInt3:       "A 3-bit signed integer can represent values from -4 to 3.",
		DTypeInt4:       "A 4-bit signed integer can represent values from -8 to 7.",
		DTypeInt5:       "A 5-bit signed integer can represent values from -16 to 15.",
		DTypeInt6:       "A 6-bit signed integer can represent values from -32 to 31.",
		DTypeInt7:       "A 7-bit signed integer can represent values from -64 to 63.",
		DTypeInt8:       "A 8-bit signed integer can represent values from -128 to 127.",
		DTypeInt16:      "A 16-bit signed integer can represent values from -32768 to 32767.",
		DTypeInt32:      "A 32-bit signed integer can represent values from -2147483648 to 2147483647.",
		DTypeInt64:      "A 64-bit signed integer can represent values from -9223372036854775808 to 9223372036854775807.",
		DTypeUint2:      "A 2-bit unsigned integer can represent values from 0 to 3.",
		DTypeUint3:      "A 3-bit unsigned integer can represent values from 0 to 7.",
		DTypeUint4:      "A 4-bit unsigned integer can represent values from 0 to 15.",
		DTypeUint5:      "A 5-bit unsigned integer can represent values from 0 to 31.",
		DTypeUint6:      "A 6-bit unsigned integer can represent values from 0 to 63.",
		DTypeUint7:      "A 7-bit unsigned integer can represent values from 0 to 127.",
		DTypeUint8:      "A 8-bit unsigned integer can represent values from 0 to 255.",
		DTypeUint16:     "A 16-bit unsigned integer can represent values from 0 to 65535.",
		DTypeUint32:     "A 32-bit unsigned integer can represent values from 0 to 4294967295.",
		DTypeUint64:     "A 64-bit unsigned integer can represent values from 0 to 18446744073709551615.",
		DTypeFloat16:    "(half-precision) A 16-bit floating point number.",
		DTypeFloat32:    "(single-precision) A 32-bit floating point number.",
		DTypeFloat64:    "(double-precision) A 64-bit floating point number.",
		DTypeBFloat16:   "(Brain Float) A 16-bit floating point number.",
		DTypeComplex64:  "A 64-bit complex",
		DTypeComplex128: "A 128-bit complex",
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

// MarshalJSON outputs the DType as a json.
func (dt DType) MarshalJSON() ([]byte, error) {
	if !dt.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + dt.String() + `"`), nil
}

// UnmarshalJSON parses the DType from json.
func (dt *DType) UnmarshalJSON(data []byte) error {
	str := string(bytes.Trim(data, `"`))
	if dType := ParseDType(str); dType.Validate() {
		*dt = dType
	}

	return nil
}

// Validate returns true if the DType is valid.
func (dt DType) Validate() bool {
	return dt != DTypeInvalid
}

// ParseDType parses the DType string.
func ParseDType(value string) DType {
	value = strings.ToLower(value)
	for k, v := range DTypeNames {
		if v == value {
			return k
		}
	}

	return DTypeInvalid
}
