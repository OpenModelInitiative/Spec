package v1

import (
	"fmt"
	"strings"
)

// DeviceKind is the device kind for memory allocation.
const (
	DeviceKindInvalid DeviceKind = iota //
	DeviceKindBPU
	DeviceKindCPU
	DeviceKindDPU
	DeviceKindFPU
	DeviceKindFPGA
	DeviceKindGPU
	DeviceKindIPU
	DeviceKindNPU
	DeviceKindTPU
	DeviceKindVPU
)

var (
	// DeviceKindNames is a map of DeviceKind names.
	DeviceKindNames = map[DeviceKind]string{
		DeviceKindInvalid: "Invalid",
		DeviceKindBPU:     "BPU",
		DeviceKindCPU:     "CPU",
		DeviceKindDPU:     "DPU",
		DeviceKindFPU:     "FPU",
		DeviceKindFPGA:    "FPGA",
		DeviceKindGPU:     "GPU",
		DeviceKindIPU:     "IPU",
		DeviceKindNPU:     "NPU",
		DeviceKindTPU:     "TPU",
		DeviceKindVPU:     "VPU",
	}

	// DeviceKindDescriptions is a map of DeviceKind names.
	DeviceKindDescriptions = map[DeviceKind]string{
		DeviceKindBPU:  "Brain Processing Unit",
		DeviceKindCPU:  "Central Processing Unit",
		DeviceKindDPU:  "Data Processing Unit",
		DeviceKindFPU:  "Floating Processing Unit,",
		DeviceKindFPGA: "Field-Programmable Gate Array",
		DeviceKindGPU:  "Graphics Processing Unit",
		DeviceKindIPU:  "Intelligent Processing Unit",
		DeviceKindNPU:  "Neural Processing Unit",
		DeviceKindTPU:  "Tensor Processing Unit",
		DeviceKindVPU:  "Vector Processing Unit",
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
		return nil, fmt.Errorf("invalid device kind")
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
