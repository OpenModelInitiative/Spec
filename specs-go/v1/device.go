package v1

import (
	"fmt"
	"strings"
)

// Kind returns the kind of the device.
func (d *Device) Kind() DeviceKind {
	return d.kind
}

// Index returns the index of the device.
func (d *Device) Index() int {
	return d.index
}

// Name returns the name of the device.
func (d *Device) Name() string {
	return d.name
}

func (d *Device) Description() string {
	return d.description
}

// Vendor returns the vendor of the device.
func (d *Device) Vendor() string {
	return d.vendor
}

// Features returns the features of the device.
func (d *Device) Features() []string {
	return d.features
}

func (d *Device) Memory() Memory {
	return d.memory
}

func (d *Device) String() string {
	return fmt.Sprintf("device(%s:%d name=%s description=%s vendor=%s features=%s memory.total=%d memory.available=%d)", d.kind, d.index, d.name, d.description, d.vendor, strings.Join(d.features, " "), d.memory.Total, d.memory.Available)
}
