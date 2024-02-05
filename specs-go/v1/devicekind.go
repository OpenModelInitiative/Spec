package v1

import (
	"fmt"
	"strings"
)

// DeviceKind is the device kind for memory allocation.
const (
	DeviceKindInvalid DeviceKind = iota //
	DeviceKindCPU
	DeviceKindGPU
	DeviceKindNPU
	DeviceKindTPU
	DeviceKindFPGA
)

var (
	// DeviceKindNames is a map of DeviceKind names.
	DeviceKindNames = map[DeviceKind]string{
		DeviceKindInvalid: "Invalid",
		DeviceKindCPU:     "CPU",
		DeviceKindGPU:     "GPU",
		DeviceKindNPU:     "NPU",
		DeviceKindTPU:     "TPU",
		DeviceKindFPGA:    "FPGA",
	}

	// DeviceKindDescriptions is a map of DeviceKind names.
	DeviceKindDescriptions = map[DeviceKind]string{
		DeviceKindCPU:  "Central Processing Unit",
		DeviceKindGPU:  "Graphics Processing Unit",
		DeviceKindNPU:  "Neural Processing Unit",
		DeviceKindTPU:  "Tensor Processing Unit",
		DeviceKindFPGA: "Field-Programmable Gate Array",
	}
)

// String outputs the DeviceKind as a string.
func (dk DeviceKind) String() string {
	return DeviceKindNames[dk]
}

// Description outputs the DeviceKind description.
func (dk DeviceKind) Description() string {
	return DeviceKindDescriptions[dk]
}

// MarshalJSON outputs the DeviceKind as a json.
func (dk DeviceKind) MarshalJSON() ([]byte, error) {
	if !dk.Valid() {
		return nil, nil
	}

	return []byte(`"` + dk.String() + `"`), nil
}

// UnmarshalJSON parses the DeviceKind from json.
func (dk *DeviceKind) UnmarshalJSON(data []byte) error {
	deviceKind, err := ParseDeviceKind(string(data))
	if err != nil {
		return err
	}

	*dk = deviceKind

	return nil
}

// Valid returns true if the DeviceKind is valid.
func (dk DeviceKind) Valid() bool {
	return dk != DeviceKindInvalid
}

// ParseDeviceKind parses the DeviceKind string.
func ParseDeviceKind(value string) (DeviceKind, error) {
	v := strings.ToUpper(strings.Trim(value, `"`))

	for k, n := range DeviceKindNames {
		if v == n {
			return k, nil
		}
	}

	return DeviceKindInvalid, fmt.Errorf("invalid device kind: %s", value)
}
