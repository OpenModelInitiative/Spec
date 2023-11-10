package v1

import (
	"github.com/leliuga/cdk/http"
	"github.com/leliuga/cdk/types"
	"github.com/opencontainers/go-digest"
)

type (
	// TypeMeta represents the object type metadata.
	TypeMeta struct {
		ResourceVersion string `json:"resourceVersion"`
		Kind            string `json:"kind"`
	}

	// ObjectMeta represents the object metadata.
	ObjectMeta struct {
		Name        string            `json:"name"`
		Description string            `json:"description"`
		Labels      types.Map[string] `json:"labels"`
		Author      string            `json:"author"`
		Url         types.URI         `json:"url"`
		Licence     string            `json:"licence"`
	}

	// ObjectStatus represents the object status.
	ObjectStatus struct {
		Status types.Map[any] `json:"status"`
	}

	DatasetAssembler struct {
		TypeMeta   `json:",inline"`
		ObjectMeta `json:",inline"`
		Spec       DatasetAssemblerSpec `json:"spec"`
		Status     ObjectStatus         `json:",inline,omitempty"`
	}

	DatasetAssemblerSpec struct {
		Init   Process `json:"init"`
		Layers []Layer `json:"layers"`
	}

	// TrainerAssembler represents a model trainer.
	TrainerAssembler struct {
		TypeMeta   `json:",inline"`
		ObjectMeta `json:",inline"`
		Spec       TrainerAssemblerSpec `json:"spec"`
		Status     ObjectStatus         `json:",inline,omitempty"`
	}

	// TrainerAssemblerSpec represents a model trainer specification.
	TrainerAssemblerSpec struct {
		Init   Process `json:"init"`
		Layers []Layer `json:"layers"`
	}

	// ModelAssembler represents a model assembler.
	ModelAssembler struct {
		TypeMeta   `json:",inline"`
		ObjectMeta `json:",inline"`
		Spec       ModelAssemblerSpec `json:"spec"`
		Status     ObjectStatus       `json:",inline,omitempty"`
	}

	// ModelAssemblerSpec represents a model specification.
	ModelAssemblerSpec struct {
		Architecture Architecture `json:"architecture"`
		Tasks        []Task       `json:"tasks"`
		Languages    []Language   `json:"languages"`
		Datasets     []string     `json:"datasets"`
		Trainer      []string     `json:"trainer"`
		Init         Process      `json:"init"`
		Layers       []Layer      `json:"layers"`
	}

	// InferenceAssembler represents a model.
	InferenceAssembler struct {
		TypeMeta   `json:",inline"`
		ObjectMeta `json:",inline"`
		Spec       InferenceAssemblerSpec `json:"spec"`
		Status     ObjectStatus           `json:",inline,omitempty"`
	}

	// InferenceAssemblerSpec represents a model specification.
	InferenceAssemblerSpec struct {
		Model     string  `json:"model"`
		Inference string  `json:"inference"`
		Init      Process `json:"init"`
		Layers    []Layer `json:"layers"`
	}

	// Dataset represents a model dataset.
	Dataset struct {
		TypeMeta   `json:",inline"`
		ObjectMeta `json:",inline"`
		Spec       DatasetSpec  `json:"spec"`
		Status     ObjectStatus `json:",inline,omitempty"`
	}

	// DatasetSpec represents a model dataset specification.
	DatasetSpec struct {
		Layers []*Layer `json:"layers"`
	}

	// Trainer represents a model trainer.
	Trainer struct {
		TypeMeta   `json:",inline"`
		ObjectMeta `json:",inline"`
		Spec       TrainerSpec  `json:"spec"`
		Status     ObjectStatus `json:",inline,omitempty"`
	}

	// TrainerSpec represents a model trainer specification.
	TrainerSpec struct {
		Init   Process `json:"init"`
		Layers []Layer `json:"layers"`
	}

	// Model represents a model.
	Model struct {
		TypeMeta   `json:",inline"`
		ObjectMeta `json:",inline"`
		Spec       ModelSpec    `json:"spec"`
		Status     ObjectStatus `json:",inline,omitempty"`
	}

	// ModelSpec represents a model specification.
	ModelSpec struct {
		Architecture Architecture            `json:"architecture"`
		Tasks        []Task                  `json:"tasks"`
		Languages    []Language              `json:"languages"`
		Variants     types.Map[ModelVariant] `json:"variants"`
	}

	// ModelVariant represents a model variant.
	ModelVariant struct {
		DType         DType   `json:"dtype"`
		ParamsSize    uint64  `json:"params_size"`
		VocabSize     uint32  `json:"vocab_size"`
		ContextSize   uint64  `json:"context_size"`
		EmbeddingSize uint64  `json:"embedding_size"`
		MinVRAMSize   uint64  `json:"min_vram_size"`
		Finetuned     bool    `json:"finetuned"`
		Quantized     bool    `json:"quantized"`
		Layers        []Layer `json:"layers"`
	}

	// Inference represents a model inference.
	Inference struct {
		TypeMeta   `json:",inline"`
		ObjectMeta `json:",inline"`
		Spec       InferenceSpec `json:"spec"`
		Status     ObjectStatus  `json:",inline,omitempty"`
	}

	// InferenceSpec represents a model inference specification.
	InferenceSpec struct {
		Init      Process    `json:"init"`
		Layers    []Layer    `json:"layers"`
		Endpoints []Endpoint `json:"endpoints"`
	}

	// Layer is a Model ModelVariant layer (e.g. model, backend, etc.)
	Layer struct {
		Name        string        `json:"name"`
		Description string        `json:"description"`
		MediaType   string        `json:"media_type"`
		Artifact    types.URI     `json:"artifact"`
		Size        uint64        `json:"size"`
		Digest      digest.Digest `json:"digest"`
	}

	// Process represents a model backend process.
	Process struct {
		Command string                   `json:"command"`
		Args    types.Map[types.Options] `json:"args"`
	}

	// Endpoint represents a model backend endpoint.
	Endpoint struct {
		Name        string                   `json:"name"`
		Description string                   `json:"description"`
		Method      http.Method              `json:"method"`
		Path        string                   `json:"path"`
		Fields      types.Map[types.Options] `json:"fields"`
	}

	// Architecture is a model architecture (e.g. transformer, lstm, etc.)
	Architecture uint16

	// DType is a data type of tensor (e.g. float32, int32, etc.)
	DType uint8

	// Kind is a object kind (e.g. model, inference, etc.)
	Kind uint8

	// Language is a model natural language (ISO 639-1)
	Language uint8

	// Task is a model task (e.g. classification, translation, etc.)
	Task uint8
)
