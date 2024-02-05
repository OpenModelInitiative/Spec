package v1

import (
	"net/url"
	"time"

	"github.com/opencontainers/go-digest"
)

type (
	Named struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	// ObjectKind represents the object version and kind.
	ObjectKind struct {
		Kind    Kind   `json:"kind"`
		Version string `json:"version"`
		ID      string `json:"id"`
	}

	// ObjectMeta represents the object metadata.
	ObjectMeta struct {
		Named   `json:",inline"`
		Labels  map[string]string `json:"labels"`
		Author  string            `json:"author"`
		Url     *url.URL          `json:"url"`
		Licence string            `json:"licence"`
	}

	// ObjectStatus represents the object status.
	ObjectStatus struct {
		Status map[string]any `json:"status"`
	}

	// DatasetAssembler represents a model dataset.
	DatasetAssembler struct {
		ObjectKind `json:",inline"`
		Spec       DatasetAssemblerSpec `json:"spec"`
		Status     ObjectStatus         `json:",inline,omitempty"`
	}

	// DatasetAssemblerSpec represents a model dataset specification.
	DatasetAssemblerSpec struct {
		ObjectMeta `json:",inline"`
		Tasks      []Task     `json:"tasks"`
		Languages  []Language `json:"languages"`
		Layers     []Layer    `json:"layers"`
		Entrypoint Entrypoint `json:"entrypoint"`
	}

	// TrainerAssembler represents a model trainer.
	TrainerAssembler struct {
		ObjectKind `json:",inline"`
		Spec       TrainerAssemblerSpec `json:"spec"`
		Status     ObjectStatus         `json:",inline,omitempty"`
	}

	// TrainerAssemblerSpec represents a model trainer specification.
	TrainerAssemblerSpec struct {
		ObjectMeta   `json:",inline"`
		Architecture Architecture `json:"architecture"`
		Tasks        []Task       `json:"tasks"`
		Languages    []Language   `json:"languages"`
		Layers       []Layer      `json:"layers"`
		Entrypoint   Entrypoint   `json:"entrypoint"`
	}

	// ModelAssembler represents a model assembler.
	ModelAssembler struct {
		ObjectKind `json:",inline"`
		Spec       ModelAssemblerSpec `json:"spec"`
		Status     ObjectStatus       `json:",inline,omitempty"`
	}

	// ModelAssemblerSpec represents a model specification.
	ModelAssemblerSpec struct {
		ObjectMeta   `json:",inline"`
		Architecture Architecture `json:"architecture"`
		Tasks        []Task       `json:"tasks"`
		Languages    []Language   `json:"languages"`
		Layers       []Layer      `json:"layers"`
		Entrypoint   Entrypoint   `json:"entrypoint"`
	}

	// InferenceAssembler represents a model.
	InferenceAssembler struct {
		ObjectKind `json:",inline"`
		Spec       InferenceAssemblerSpec `json:"spec"`
		Status     ObjectStatus           `json:",inline,omitempty"`
	}

	// InferenceAssemblerSpec represents a model specification.
	InferenceAssemblerSpec struct {
		ObjectMeta   `json:",inline"`
		Architecture Architecture `json:"architecture"`
		Layers       []Layer      `json:"layers"`
		Entrypoint   Entrypoint   `json:"entrypoint"`
	}

	// Dataset represents a model dataset.
	Dataset struct {
		ObjectKind `json:",inline"`
		Spec       DatasetSpec  `json:"spec"`
		Status     ObjectStatus `json:",inline,omitempty"`
	}

	// DatasetSpec represents a model dataset specification.
	DatasetSpec struct {
		ObjectMeta `json:",inline"`
		// History describes the history of each assembler step.
		History   []History  `json:"history"`
		Tasks     []Task     `json:"tasks"`
		Languages []Language `json:"languages"`
		Layers    []Layer    `json:"layers"`
	}

	// Trainer represents a model trainer.
	Trainer struct {
		ObjectKind `json:",inline"`
		Spec       TrainerSpec  `json:"spec"`
		Status     ObjectStatus `json:",inline,omitempty"`
	}

	// TrainerSpec represents a model trainer specification.
	TrainerSpec struct {
		ObjectMeta `json:",inline"`
		// History describes the history of each assembler step.
		History      []History    `json:"history"`
		Architecture Architecture `json:"architecture"`
		Tasks        []Task       `json:"tasks"`
		Languages    []Language   `json:"languages"`
		Layers       []Layer      `json:"layers"`
		Entrypoint   Entrypoint   `json:"entrypoint"`
	}

	// Model represents a model.
	Model struct {
		ObjectKind `json:",inline"`
		Spec       ModelSpec    `json:"spec"`
		Status     ObjectStatus `json:",inline,omitempty"`
	}

	// ModelSpec represents a model specification.
	ModelSpec struct {
		ObjectMeta `json:",inline"`
		// History describes the history of each assembler step.
		History      []History               `json:"history"`
		Architecture Architecture            `json:"architecture"`
		Tasks        []Task                  `json:"tasks"`
		Languages    []Language              `json:"languages"`
		Variants     map[string]ModelVariant `json:"variants"`
	}

	// ModelVariant represents a model variant.
	ModelVariant struct {
		DataKind      DataKind `json:"data_kind"`
		ParamsSize    uint64   `json:"params_size"`
		VocabSize     uint32   `json:"vocab_size"`
		ContextSize   uint64   `json:"context_size"`
		EmbeddingSize uint64   `json:"embedding_size"`
		MinVRAMSize   uint64   `json:"min_vram_size"`
		Finetuned     bool     `json:"finetuned"`
		Quantized     bool     `json:"quantized"`
		Layers        []Layer  `json:"layers"`
	}

	// Inference represents a model inference.
	Inference struct {
		ObjectKind `json:",inline"`
		Spec       InferenceSpec `json:"spec"`
		Status     ObjectStatus  `json:",inline,omitempty"`
	}

	// InferenceSpec represents a model inference specification.
	InferenceSpec struct {
		ObjectMeta   `json:",inline"`
		Architecture Architecture       `json:"architecture"`
		Variants     []InferenceVariant `json:"variants"`
	}

	// InferenceVariant represents a model inference variant.
	InferenceVariant struct {
		// History describes the history of each assembler step.
		History    []History  `json:"history"`
		Platform   Platform   `json:"platform"`
		Layers     []Layer    `json:"layers"`
		Entrypoint Entrypoint `json:"entrypoint"`
	}

	// Architecture represents a model architecture.
	Architecture struct {
		Kind    ArchitectureKind `json:"kind"`
		Version string           `json:"version"`
	}

	// Layer is a layer of an object.
	Layer struct {
		Named     `json:",inline"`
		MediaType string        `json:"media_type"`
		Size      uint64        `json:"size"`
		Digest    digest.Digest `json:"digest"`
	}

	// History describes the history of assembler steps.
	History struct {
		Created   *time.Time `json:"created,omitempty"`
		CreatedBy string     `json:"created_by,omitempty"`
		Comment   string     `json:"comment,omitempty"`
	}

	// Entrypoint defines the execution parameters which should be used when running a command.
	Entrypoint struct {
		Env  []string `json:"Env,omitempty"`
		Cmd  string   `json:"Cmd"`
		Args []string `json:"Args,omitempty"`
	}

	// Platform represents a inference process platform.
	Platform struct {
		Architecture string   `json:"architecture"`
		OS           string   `json:"os"`
		OSVersion    string   `json:"os.version,omitempty"`
		OSFeatures   []string `json:"os.features,omitempty"`
		Variant      string   `json:"variant,omitempty"`
	}

	// Tensor is a tensor.
	Tensor struct {
		device IDevice
		kind   DataKind
		shape  []uint64
		data   []byte
	}

	// TensorNamed is a tensor with a name.
	TensorNamed struct {
		name        string
		description string
		Tensor
	}

	// Device is a device.
	Device struct {
		kind        DeviceKind
		index       int
		name        string
		description string
		vendor      string
		features    []string
		memory      Memory
	}

	// Memory is a memory struct.
	Memory struct {
		Total     uint64
		Available uint64
	}

	// IDevice is a device interface.
	IDevice interface {
		Kind() Kind
		Index() int
		Name() string
		Description() string
		Vendor() string
		Features() []string
		Memory() Memory
		String() string
	}

	// ArchitectureKind is a model architecture (e.g. transformer, lstm, etc.)
	ArchitectureKind uint16

	// DataKind is a data kind of tensor (e.g. F32, I16, etc.)
	DataKind uint8

	DeviceKind uint8

	// Kind is a object kind (e.g. model, inference, etc.)
	Kind uint8

	// Language is a model natural language (ISO 639-1)
	Language uint8

	// Task is a model task (e.g. classification, translation, etc.)
	Task uint8
)
