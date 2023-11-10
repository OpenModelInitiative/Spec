package v1

import (
	"net/url"
	"time"

	"github.com/opencontainers/go-digest"
)

type (
	// ObjectKind represents the object version and kind.
	ObjectKind struct {
		Version string `json:"version"`
		Kind    Kind   `json:"kind"`
	}

	// ObjectMeta represents the object metadata.
	ObjectMeta struct {
		Name        string            `json:"name"`
		Description string            `json:"description"`
		Labels      map[string]string `json:"labels"`
		Author      string            `json:"author"`
		Url         *url.URL          `json:"url"`
		Licence     string            `json:"licence"`
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
		Run        Run        `json:"run"`
	}

	// TrainerAssembler represents a model trainer.
	TrainerAssembler struct {
		ObjectKind `json:",inline"`
		Spec       TrainerAssemblerSpec `json:"spec"`
		Status     ObjectStatus         `json:",inline,omitempty"`
	}

	// TrainerAssemblerSpec represents a model trainer specification.
	TrainerAssemblerSpec struct {
		ObjectMeta `json:",inline"`
		Layers     []Layer `json:"layers"`
		Run        Run     `json:"run"`
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
		Datasets     []string     `json:"datasets"`
		Trainer      string       `json:"trainer"`
		Layers       []Layer      `json:"layers"`
		Run          Run          `json:"run"`
	}

	// InferenceAssembler represents a model.
	InferenceAssembler struct {
		ObjectKind `json:",inline"`
		Spec       InferenceAssemblerSpec `json:"spec"`
		Status     ObjectStatus           `json:",inline,omitempty"`
	}

	// InferenceAssemblerSpec represents a model specification.
	InferenceAssemblerSpec struct {
		ObjectMeta `json:",inline"`
		Model      string  `json:"model"`
		Layers     []Layer `json:"layers"`
		Run        Run     `json:"run"`
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
		History []History `json:"history"`

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
		History []History `json:"history"`

		Datasets []string `json:"datasets"`
		Layers   []Layer  `json:"layers"`
		Run      Run      `json:"run"`
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
		History []History `json:"history"`

		Architecture Architecture            `json:"architecture"`
		Tasks        []Task                  `json:"tasks"`
		Languages    []Language              `json:"languages"`
		Inference    string                  `json:"inference"`
		Variants     map[string]ModelVariant `json:"variants"`
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
		ObjectKind `json:",inline"`
		Spec       InferenceSpec `json:"spec"`
		Status     ObjectStatus  `json:",inline,omitempty"`
	}

	// InferenceSpec represents a model inference specification.
	InferenceSpec struct {
		ObjectMeta `json:",inline"`
		// History describes the history of each assembler step.
		History []History `json:"history"`

		Run      Run                `json:"run"`
		Variants []InferenceVariant `json:"variants"`
	}

	// InferenceVariant represents a model inference variant.
	InferenceVariant struct {
		Platform Platform `json:"platform"`
		Layers   []Layer  `json:"layers"`
	}

	// Layer is
	Layer struct {
		Name        string        `json:"name"`
		Description string        `json:"description"`
		MediaType   string        `json:"media_type"`
		Artifact    *url.URL      `json:"artifact"`
		Size        uint64        `json:"size"`
		Digest      digest.Digest `json:"digest"`
	}

	// History describes the history of assembler steps.
	History struct {
		// Created is the combined date and time at which was created, formatted as defined by RFC 3339, section 5.6.
		Created *time.Time `json:"created,omitempty"`

		// CreatedBy is the command which created.
		CreatedBy string `json:"created_by,omitempty"`

		// Comment is a custom message.
		Comment string `json:"comment,omitempty"`
	}

	// Run defines the execution parameters which should be used when running a command.
	Run struct {
		// Env is a list of environment variables to be used in a container.
		Env []string `json:"Env,omitempty"`

		// Cmd defines the default arguments to the entrypoint of the container.
		Cmd []string `json:"Cmd,omitempty"`
	}

	// Platform represents a inference process platform.
	Platform struct {
		// Architecture field specifies the CPU architecture, for example
		// `amd64` or `ppc64le`.
		Architecture string `json:"architecture"`

		// OS specifies the operating system, for example `linux` or `windows`.
		OS string `json:"os"`

		// OSVersion is an optional field specifying the operating system
		// version, for example on Windows `10.0.14393.1066`.
		OSVersion string `json:"os.version,omitempty"`

		// OSFeatures is an optional field specifying an array of strings,
		// each listing a required OS feature (for example on Windows `win32k`).
		OSFeatures []string `json:"os.features,omitempty"`

		// Variant is an optional field specifying a variant of the CPU, for
		// example `v7` to specify ARMv7 when architecture is `arm`.
		Variant string `json:"variant,omitempty"`
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
