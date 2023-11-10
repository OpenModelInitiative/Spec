package v1

import (
	"fmt"
	"strings"
)

// Task is a model task.
const (
	TaskInvalid Task = iota //
	TaskTextClassification
	TaskTextZeroShotClassification
	TaskTextTokenClassification
	TaskTextTranslation
	TaskTextSummarization
	TaskTextSimilarity
	TaskTextFillMask
	TaskTextToText
	TaskTextToEmbedding
	TaskTextToImage
	TaskTextToAudio
	TaskTextToSpeech
	TaskTextToVideo
	TaskImageClassification
	TaskImageZeroShotClassification
	TaskImageSimilarity
	TaskImageQuestionAnswering
	TaskImageToImage
	TaskImageToEmbedding
	TaskImageToText
	TaskAudioClassification
	TaskAudioZeroShotClassification
	TaskAudioSimilarity
	TaskAudioQuestionAnswering
	TaskAudioToAudio
	TaskAudioToEmbedding
	TaskAudioToText
	TaskVideoClassification
	TaskVideoZeroShotClassification
	TaskVideoSimilarity
	TaskVideoQuestionAnswering
	TaskVideoToVideo
	TaskVideoToEmbedding
	TaskVideoToText
)

var (
	// TaskNames is a map of Task names.
	TaskNames = map[Task]string{
		TaskInvalid:                     "invalid",
		TaskTextClassification:          "text-classification",
		TaskTextZeroShotClassification:  "text-zero-shot-classification",
		TaskTextTokenClassification:     "text-token-classification",
		TaskTextTranslation:             "text-translation",
		TaskTextSummarization:           "text-summarization",
		TaskTextSimilarity:              "text-similarity",
		TaskTextFillMask:                "text-fill-mask",
		TaskTextToText:                  "text-to-text",
		TaskTextToEmbedding:             "text-to-embedding",
		TaskTextToImage:                 "text-to-image",
		TaskTextToAudio:                 "text-to-audio",
		TaskTextToSpeech:                "text-to-speech",
		TaskTextToVideo:                 "text-to-video",
		TaskImageClassification:         "image-classification",
		TaskImageZeroShotClassification: "image-zero-shot-classification",
		TaskImageSimilarity:             "image-similarity",
		TaskImageQuestionAnswering:      "image-question-answering",
		TaskImageToImage:                "image-to-image",
		TaskImageToEmbedding:            "image-to-embedding",
		TaskImageToText:                 "image-to-text",
		TaskAudioClassification:         "audio-classification",
		TaskAudioZeroShotClassification: "audio-zero-shot-classification",
		TaskAudioSimilarity:             "audio-similarity",
		TaskAudioQuestionAnswering:      "audio-question-answering",
		TaskAudioToAudio:                "audio-to-audio",
		TaskAudioToEmbedding:            "audio-to-embedding",
		TaskAudioToText:                 "audio-to-text",
		TaskVideoClassification:         "video-classification",
		TaskVideoZeroShotClassification: "video-zero-shot-classification",
		TaskVideoSimilarity:             "video-similarity",
		TaskVideoQuestionAnswering:      "video-question-answering",
		TaskVideoToVideo:                "video-to-video",
		TaskVideoToEmbedding:            "video-to-embedding",
		TaskVideoToText:                 "video-to-text",
	}

	// TaskDescriptions is a map of Task descriptions.
	TaskDescriptions = map[Task]string{
		TaskInvalid:                     "Invalid task.",
		TaskTextClassification:          "Text classification is the task of assigning a set of predefined categories to open-ended text.",
		TaskTextZeroShotClassification:  "Zero-shot classification is the task of predicting a predefined category for a piece of text without any training examples.",
		TaskTextTokenClassification:     "Token classification is the task of assigning a label to each token in a sequence.",
		TaskTextTranslation:             "Text translation is the task of translating a piece of text from one language to another.",
		TaskTextSummarization:           "Text summarization is the task of summarizing a piece of text into a shorter text.",
		TaskTextSimilarity:              "Text similarity is the task of predicting a similarity score between two pieces of text.",
		TaskTextFillMask:                "Text fill mask is the task of predicting a masked word in a piece of text.",
		TaskTextToText:                  "Text to text is the task of generating a piece of text from another piece of text.",
		TaskTextToEmbedding:             "Text to embedding is the task of generating an embedding vector from a piece of text.",
		TaskTextToImage:                 "Text to image is the task of generating an image from a piece of text.",
		TaskTextToAudio:                 "Text to audio is the task of generating an audio from a piece of text.",
		TaskTextToSpeech:                "Text to speech is the task of generating a speech from a piece of text.",
		TaskTextToVideo:                 "Text to video is the task of generating a video from a piece of text.",
		TaskImageClassification:         "Image classification is the task of assigning a set of predefined categories to an image.",
		TaskImageZeroShotClassification: "Zero-shot classification is the task of predicting a predefined category for an image without any training examples.",
		TaskImageSimilarity:             "Image similarity is the task of predicting a similarity score between two images.",
		TaskImageQuestionAnswering:      "Image question answering is the task of answering a question about an image.",
		TaskImageToImage:                "Image to image is the task of generating an image from another image.",
		TaskImageToEmbedding:            "Image to embedding is the task of generating an embedding vector from an image.",
		TaskImageToText:                 "Image to text is the task of generating a piece of text from an image.",
		TaskAudioClassification:         "Audio classification is the task of assigning a set of predefined categories to an audio.",
		TaskAudioZeroShotClassification: "Zero-shot classification is the task of predicting a predefined category for an audio without any training examples.",
		TaskAudioSimilarity:             "Audio similarity is the task of predicting a similarity score between two audios.",
		TaskAudioQuestionAnswering:      "Audio question answering is the task of answering a question about an audio.",
		TaskAudioToAudio:                "Audio to audio is the task of generating an audio from another audio.",
		TaskAudioToEmbedding:            "Audio to embedding is the task of generating an embedding vector from an audio.",
		TaskAudioToText:                 "Audio to text is the task of generating a piece of text from an audio.",
		TaskVideoClassification:         "Video classification is the task of assigning a set of predefined categories to a video.",
		TaskVideoZeroShotClassification: "Zero-shot classification is the task of predicting a predefined category for a video without any training examples.",
		TaskVideoSimilarity:             "Video similarity is the task of predicting a similarity score between two videos.",
		TaskVideoQuestionAnswering:      "Video question answering is the task of answering a question about a video.",
		TaskVideoToVideo:                "Video to video is the task of generating a video from another video.",
		TaskVideoToEmbedding:            "Video to embedding is the task of generating an embedding vector from a video.",
		TaskVideoToText:                 "Video to text is the task of generating a piece of text from a video.",
	}
)

// String outputs the Task as a string.
func (t Task) String() string {
	return TaskNames[t]
}

// Description outputs the Task description.
func (t Task) Description() string {
	return TaskDescriptions[t]
}

// MarshalJSON outputs the Task as a json.
func (t Task) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return nil, fmt.Errorf("invalid task")
	}

	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON parses the Task from json.
func (t *Task) UnmarshalJSON(data []byte) error {
	task, err := ParseTask(string(data))
	if err != nil {
		return err
	}

	*t = task

	return nil
}

// Valid returns true if the Task is valid.
func (t Task) Valid() bool {
	return t != TaskInvalid
}

// ParseTask parses the Task string.
func ParseTask(value string) (Task, error) {
	v := strings.ToLower(strings.Trim(value, `"`))

	for k, n := range TaskNames {
		if v == n {
			return k, nil
		}
	}

	return TaskInvalid, fmt.Errorf("invalid task: %s", value)
}
