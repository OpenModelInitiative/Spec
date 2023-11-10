package v1

// NewObjectKind creates a new object kind.
func NewObjectKind(kind Kind) ObjectKind {
	return ObjectKind{Version: Version, Kind: kind}
}

// NewObjectDatasetAssemblerKind creates a new object kind for a dataset assembler.
func NewObjectDatasetAssemblerKind() ObjectKind {
	return NewObjectKind(KindDatasetAssembler)
}

// NewObjectTrainerAssemblerKind creates a new object kind for a trainer assembler.
func NewObjectTrainerAssemblerKind() ObjectKind {
	return NewObjectKind(KindTrainerAssembler)
}

// NewObjectModelAssemblerKind creates a new object kind for a model assembler.
func NewObjectModelAssemblerKind() ObjectKind {
	return NewObjectKind(KindModelAssembler)
}

// NewObjectInferenceAssemblerKind creates a new object kind for a inference assembler.
func NewObjectInferenceAssemblerKind() ObjectKind {
	return NewObjectKind(KindInferenceAssembler)
}

// NewObjectDatasetKind creates a new object kind for a dataset.
func NewObjectDatasetKind() ObjectKind {
	return NewObjectKind(KindDataset)
}

// NewObjectTrainerKind creates a new object kind for a trainer.
func NewObjectTrainerKind() ObjectKind {
	return NewObjectKind(KindTrainer)
}

// NewObjectModelKind creates a new object kind for a model.
func NewObjectModelKind() ObjectKind {
	return NewObjectKind(KindModel)
}

// NewObjectInferenceKind creates a new object kind for a inference.
func NewObjectInferenceKind() ObjectKind {
	return NewObjectKind(KindInference)
}
