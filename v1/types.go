package v1

type (
	Model struct {
		Name          string       `json:"name"`
		Description   string       `json:"description"`
		Author        string       `json:"author"`
		Url           string       `json:"url"`
		Licence       string       `json:"licence"`
		Architecture  Architecture `json:"architecture"`
		DType         DType        `json:"d_type"`
		Languages     []Language   `json:"languages"`
		Tasks         []Task       `json:"tasks"`
		ParamsSize    uint64       `json:"params_size"`
		VocabSize     uint32       `json:"vocab_size"`
		ContextSize   uint16       `json:"context_size"`
		EmbeddingSize uint16       `json:"embedding_size"`
		Quantized     bool         `json:"quantized"`
		FineTuned     bool         `json:"fine_tuned"`
	}

	// Architecture is a model kind
	Architecture uint16

	// DType is a data type of tensor
	DType uint8

	// Language is a model natural language (ISO 639-1)
	Language uint8

	// Task is a model task
	Task uint8
)
