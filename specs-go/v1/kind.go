package v1

import (
	"fmt"
	"strings"
)

const (
	KindInvalid Kind = iota //
	KindDatasetAssembler
	KindTrainerAssembler
	KindModelAssembler
	KindInferenceAssembler
	KindDataset
	KindTrainer
	KindModel
	KindInference
)

var (
	KindNames = map[Kind]string{
		KindInvalid:            "Invalid",
		KindDatasetAssembler:   "DatasetAssembler",
		KindTrainerAssembler:   "TrainerAssembler",
		KindModelAssembler:     "ModelAssembler",
		KindInferenceAssembler: "InferenceAssembler",
		KindDataset:            "Dataset",
		KindTrainer:            "Trainer",
		KindModel:              "Model",
		KindInference:          "Inference",
	}
)

// String outputs the Kind as a string.
func (k Kind) String() string {
	return KindNames[k]
}

// MarshalJSON outputs the Kind as a json.
func (k Kind) MarshalJSON() ([]byte, error) {
	if !k.Valid() {
		return nil, fmt.Errorf("invalid kind")
	}

	return []byte(`"` + k.String() + `"`), nil
}

// UnmarshalJSON parses the Kind from json.
func (k *Kind) UnmarshalJSON(data []byte) error {
	kind, err := ParseKind(string(data))
	if err != nil {
		return err
	}

	*k = kind

	return nil
}

// Valid returns true if the Kind is valid.
func (k Kind) Valid() bool {
	return k != KindInvalid
}

// ParseKind parses the Kind string.
func ParseKind(value string) (Kind, error) {
	v := strings.ToLower(strings.Trim(value, `"`))

	for k, n := range KindNames {
		if v == strings.ToLower(n) {
			return k, nil
		}
	}

	return KindInvalid, fmt.Errorf("invalid kind: %s", value)
}
