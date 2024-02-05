package v1

import (
	"fmt"
)

// Device is the device for memory allocation and operations.
func (t *Tensor) Device() IDevice {
	return t.device
}

// Kind returns the data kind of the tensor.
func (t *Tensor) Kind() DataKind {
	return t.kind
}

// Shape returns the shape of the tensor.
func (t *Tensor) Shape() []uint64 {
	return t.shape
}

// Size returns the size of the tensor.
func (t *Tensor) Size() uint64 {
	var size uint64 = 0
	for _, d := range t.shape {
		size += d
	}

	return size
}

// String outputs the Tensor as a string.
func (t *Tensor) String() string {
	return fmt.Sprintf("tensor(%s device=%s:%d shape=%v size=%d)", t.kind, t.device.Kind(), t.device.Index(), t.shape, t.Size())
}
