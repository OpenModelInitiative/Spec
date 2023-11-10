package v1

import (
	"fmt"
	"strings"
)

// Architecture is a model kind
const (
	ArchitectureInvalid Architecture = iota
	ArchitectureAlbert
	ArchitectureAlign
	ArchitectureAltclip
	ArchitectureAudioSpectrogramTransformer
	ArchitectureAutoformer
	ArchitectureBark
	ArchitectureBart
	ArchitectureBeit
	ArchitectureBert
	ArchitectureBertGeneration
	ArchitectureBigBird
	ArchitectureBigbirdPegasus
	ArchitectureBioGpt
	ArchitectureBit
	ArchitectureBlenderbot
	ArchitectureBlenderbotSmall
	ArchitectureBlip
	ArchitectureBlip2
	ArchitectureBloom
	ArchitectureBridgeTower
	ArchitectureBros
	ArchitectureCamembert
	ArchitectureCanine
	ArchitectureChineseClip
	ArchitectureClap
	ArchitectureClip
	ArchitectureClipseg
	ArchitectureCodegen
	ArchitectureConditionalDetr
	ArchitectureConvbert
	ArchitectureConvnext
	ArchitectureConvnextv2
	ArchitectureCpmant
	ArchitectureCtrl
	ArchitectureCvt
	ArchitectureData2Vec
	ArchitectureData2VecText
	ArchitectureData2VecVision
	ArchitectureDeberta
	ArchitectureDebertaV2
	ArchitectureDecisionTransformer
	ArchitectureDeformableDetr
	ArchitectureDeit
	ArchitectureDeta
	ArchitectureDetr
	ArchitectureDinat
	ArchitectureDinov2
	ArchitectureDistilbert
	ArchitectureDpr
	ArchitectureDpt
	ArchitectureEfficientformer
	ArchitectureEfficientnet
	ArchitectureElectra
	ArchitectureEncodec
	ArchitectureEncoderDecoder
	ArchitectureErnie
	ArchitectureErnieM
	ArchitectureEsm
	ArchitectureFalcon
	ArchitectureFlaubert
	ArchitectureFlava
	ArchitectureFnet
	ArchitectureFocalnet
	ArchitectureFsmt
	ArchitectureFunnel
	ArchitectureGit
	ArchitectureGlpn
	ArchitectureGpt2
	ArchitectureGptBigcode
	ArchitectureGptNeo
	ArchitectureGptNeox
	ArchitectureGptNeoxJapanese
	ArchitectureGptj
	ArchitectureGptsanJapanese
	ArchitectureGraphormer
	ArchitectureGroupvit
	ArchitectureHubert
	ArchitectureIbert
	ArchitectureIdefics
	ArchitectureImagegpt
	ArchitectureInformer
	ArchitectureInstructblip
	ArchitectureJukebox
	ArchitectureLayoutlm
	ArchitectureLayoutlmv2
	ArchitectureLayoutlmv3
	ArchitectureLed
	ArchitectureLevit
	ArchitectureLilt
	ArchitectureLlama
	ArchitectureLongformer
	ArchitectureLongt5
	ArchitectureLuke
	ArchitectureLxmert
	ArchitectureM2m100
	ArchitectureMarian
	ArchitectureMarkuplm
	ArchitectureMask2former
	ArchitectureMaskformer
	ArchitectureMbart
	ArchitectureMega
	ArchitectureMegatronBert
	ArchitectureMgpStr
	ArchitectureMistral
	ArchitectureMobilebert
	ArchitectureMobilenetV1
	ArchitectureMobilenetV2
	ArchitectureMobilevit
	ArchitectureMobilevitv2
	ArchitectureMpnet
	ArchitectureMpt
	ArchitectureMra
	ArchitectureMt5
	ArchitectureMusicgen
	ArchitectureMvp
	ArchitectureNat
	ArchitectureNezha
	ArchitectureNllbMoe
	ArchitectureNystromformer
	ArchitectureOneformer
	ArchitectureOpenai
	ArchitectureOpt
	ArchitectureOwlv2
	ArchitectureOwlvit
	ArchitecturePegasus
	ArchitecturePegasusX
	ArchitecturePerceiver
	ArchitecturePersimmon
	ArchitecturePix2struct
	ArchitecturePlbart
	ArchitecturePoolformer
	ArchitecturePop2piano
	ArchitectureProphetnet
	ArchitecturePvt
	ArchitectureRealm
	ArchitectureReformer
	ArchitectureRegnet
	ArchitectureRembert
	ArchitectureResnet
	ArchitectureRoberta
	ArchitectureRocBert
	ArchitectureRoformer
	ArchitectureRwkv
	ArchitectureSam
	ArchitectureSegformer
	ArchitectureSew
	ArchitectureSewD
	ArchitectureSpeechEncoderDecoder
	ArchitectureSpeechToText
	ArchitectureSpeechToText2
	ArchitectureSpeecht5
	ArchitectureSplinter
	ArchitectureSqueezebert
	ArchitectureSwiftformer
	ArchitectureSwin
	ArchitectureSwin2sr
	ArchitectureSwinv2
	ArchitectureSwitchTransformers
	ArchitectureT5
	ArchitectureTableTransformer
	ArchitectureTapas
	ArchitectureTimeSeriesTransformer
	ArchitectureTimesformer
	ArchitectureTransfoXl
	ArchitectureTrocr
	ArchitectureTvlt
	ArchitectureUmt5
	ArchitectureUnispeech
	ArchitectureUnispeechSat
	ArchitectureUpernet
	ArchitectureVideomae
	ArchitectureVilt
	ArchitectureVisionEncoderDecoder
	ArchitectureVisionTextDualEncoder
	ArchitectureVisualBert
	ArchitectureVit
	ArchitectureVitHybrid
	ArchitectureVitMae
	ArchitectureVitMsn
	ArchitectureVitdet
	ArchitectureVitmatte
	ArchitectureVits
	ArchitectureVivit
	ArchitectureWav2vec2
	ArchitectureWav2vec2Conformer
	ArchitectureWavlm
	ArchitectureWhisper
	ArchitectureXClip
	ArchitectureXglm
	ArchitectureXlm
	ArchitectureXlmProphetnet
	ArchitectureXlmRoberta
	ArchitectureXlmRobertaXl
	ArchitectureXlnet
	ArchitectureXmod
	ArchitectureYolos
	ArchitectureYoso
)

var (
	// ArchitectureNames is a map of Architecture to string.
	ArchitectureNames = map[Architecture]string{
		ArchitectureInvalid:                     "Invalid",
		ArchitectureAlbert:                      "Albert",
		ArchitectureAlign:                       "Align",
		ArchitectureAltclip:                     "AltCLIP",
		ArchitectureAudioSpectrogramTransformer: "AST",
		ArchitectureAutoformer:                  "Autoformer",
		ArchitectureBark:                        "Bark",
		ArchitectureBart:                        "Bart",
		ArchitectureBeit:                        "Beit",
		ArchitectureBert:                        "Bert",
		ArchitectureBertGeneration:              "BertGeneration",
		ArchitectureBigBird:                     "BigBird",
		ArchitectureBigbirdPegasus:              "BigBirdPegasus",
		ArchitectureBioGpt:                      "BioGpt",
		ArchitectureBit:                         "Bit",
		ArchitectureBlenderbot:                  "Blenderbot",
		ArchitectureBlenderbotSmall:             "BlenderbotSmall",
		ArchitectureBlip:                        "Blip",
		ArchitectureBlip2:                       "Blip2",
		ArchitectureBloom:                       "Bloom",
		ArchitectureBridgeTower:                 "BridgeTower",
		ArchitectureBros:                        "Bros",
		ArchitectureCamembert:                   "Camembert",
		ArchitectureCanine:                      "Canine",
		ArchitectureChineseClip:                 "ChineseClip",
		ArchitectureClap:                        "Clap",
		ArchitectureClip:                        "CLIP",
		ArchitectureClipseg:                     "CLIPSeg",
		ArchitectureCodegen:                     "CodeGen",
		ArchitectureConditionalDetr:             "ConditionalDETR",
		ArchitectureConvbert:                    "ConvBert",
		ArchitectureConvnext:                    "ConvNext",
		ArchitectureConvnextv2:                  "ConvNextV2",
		ArchitectureCpmant:                      "CPMAnt",
		ArchitectureCtrl:                        "CTRL",
		ArchitectureCvt:                         "Cvt",
		ArchitectureData2Vec:                    "Data2Vec",
		ArchitectureData2VecText:                "Data2VecText",
		ArchitectureData2VecVision:              "Data2VecVision",
		ArchitectureDeberta:                     "Deberta",
		ArchitectureDebertaV2:                   "DebertaV2",
		ArchitectureDecisionTransformer:         "DecisionTransformer",
		ArchitectureDeformableDetr:              "DeformableDetr",
		ArchitectureDeit:                        "DeiT",
		ArchitectureDeta:                        "Deta",
		ArchitectureDetr:                        "Detr",
		ArchitectureDinat:                       "Dinat",
		ArchitectureDinov2:                      "Dinov2",
		ArchitectureDistilbert:                  "DistilBert",
		ArchitectureDpr:                         "DPR",
		ArchitectureDpt:                         "DPT",
		ArchitectureEfficientformer:             "EfficientFormer",
		ArchitectureEfficientnet:                "EfficientNet",
		ArchitectureElectra:                     "Electra",
		ArchitectureEncodec:                     "Encodec",
		ArchitectureEncoderDecoder:              "EncoderDecoder",
		ArchitectureErnie:                       "Ernie",
		ArchitectureErnieM:                      "ErnieM",
		ArchitectureEsm:                         "ESM",
		ArchitectureFalcon:                      "Falcon",
		ArchitectureFlaubert:                    "Flaubert",
		ArchitectureFlava:                       "Flava",
		ArchitectureFnet:                        "FNet",
		ArchitectureFocalnet:                    "FocalNet",
		ArchitectureFsmt:                        "FSMT",
		ArchitectureFunnel:                      "Funnel",
		ArchitectureGit:                         "Git",
		ArchitectureGlpn:                        "GLPN",
		ArchitectureGpt2:                        "GPT2",
		ArchitectureGptBigcode:                  "GPTBigCode",
		ArchitectureGptNeo:                      "GPTNeo",
		ArchitectureGptNeox:                     "GPTNeox",
		ArchitectureGptNeoxJapanese:             "GPTNeoxJapanese",
		ArchitectureGptj:                        "GPTJ",
		ArchitectureGptsanJapanese:              "GPTSanJapanese",
		ArchitectureGraphormer:                  "Graphormer",
		ArchitectureGroupvit:                    "GroupViT",
		ArchitectureHubert:                      "Hubert",
		ArchitectureIbert:                       "IBert",
		ArchitectureIdefics:                     "Idefics",
		ArchitectureImagegpt:                    "ImageGPT",
		ArchitectureInformer:                    "Informer",
		ArchitectureInstructblip:                "InstructBlip",
		ArchitectureJukebox:                     "Jukebox",
		ArchitectureLayoutlm:                    "LayoutLM",
		ArchitectureLayoutlmv2:                  "LayoutLMv2",
		ArchitectureLayoutlmv3:                  "LayoutLMv3",
		ArchitectureLed:                         "LED",
		ArchitectureLevit:                       "Levit",
		ArchitectureLilt:                        "Lilt",
		ArchitectureLlama:                       "Llama",
		ArchitectureLongformer:                  "Longformer",
		ArchitectureLongt5:                      "LongT5",
		ArchitectureLuke:                        "Luke",
		ArchitectureLxmert:                      "Lxmert",
		ArchitectureM2m100:                      "M2M100",
		ArchitectureMarian:                      "MarianMT",
		ArchitectureMarkuplm:                    "MarkupLM",
		ArchitectureMask2former:                 "Mask2Former",
		ArchitectureMaskformer:                  "MaskFormer",
		ArchitectureMbart:                       "MBart",
		ArchitectureMega:                        "Mega",
		ArchitectureMegatronBert:                "MegatronBert",
		ArchitectureMgpStr:                      "MGPSTR",
		ArchitectureMistral:                     "Mistral",
		ArchitectureMobilebert:                  "MobileBert",
		ArchitectureMobilenetV1:                 "MobileNetV1",
		ArchitectureMobilenetV2:                 "MobileNetV2",
		ArchitectureMobilevit:                   "MobileViT",
		ArchitectureMobilevitv2:                 "MobileViTv2",
		ArchitectureMpnet:                       "MPNet",
		ArchitectureMpt:                         "MPT",
		ArchitectureMra:                         "Mra",
		ArchitectureMt5:                         "MT5",
		ArchitectureMusicgen:                    "Musicgen",
		ArchitectureMvp:                         "Mvp",
		ArchitectureNat:                         "Nat",
		ArchitectureNezha:                       "NeZha",
		ArchitectureNllbMoe:                     "NllbMoe",
		ArchitectureNystromformer:               "Nystromformer",
		ArchitectureOneformer:                   "OneFormer",
		ArchitectureOpenai:                      "OpenAIGPTLMHead",
		ArchitectureOpt:                         "OPT",
		ArchitectureOwlv2:                       "Owlv2",
		ArchitectureOwlvit:                      "OwlViT",
		ArchitecturePegasus:                     "Pegasus",
		ArchitecturePegasusX:                    "PegasusX",
		ArchitecturePerceiver:                   "Perceiver",
		ArchitecturePersimmon:                   "Persimmon",
		ArchitecturePix2struct:                  "Pix2Struct",
		ArchitecturePlbart:                      "PLBart",
		ArchitecturePoolformer:                  "PoolFormer",
		ArchitecturePop2piano:                   "Pop2Piano",
		ArchitectureProphetnet:                  "ProphetNet",
		ArchitecturePvt:                         "Pvt",
		ArchitectureRealm:                       "RealM",
		ArchitectureReformer:                    "Reformer",
		ArchitectureRegnet:                      "RegNet",
		ArchitectureRembert:                     "RemBert",
		ArchitectureResnet:                      "ResNet",
		ArchitectureRoberta:                     "Roberta",
		ArchitectureRocBert:                     "RoCBert",
		ArchitectureRoformer:                    "RoFormer",
		ArchitectureRwkv:                        "Rwkv",
		ArchitectureSam:                         "Sam",
		ArchitectureSegformer:                   "Segformer",
		ArchitectureSew:                         "SEW",
		ArchitectureSewD:                        "SEWD",
		ArchitectureSpeechEncoderDecoder:        "SpeechEncoderDecoder",
		ArchitectureSpeechToText:                "SpeechToText",
		ArchitectureSpeechToText2:               "SpeechToText2",
		ArchitectureSpeecht5:                    "SpeechT5",
		ArchitectureSplinter:                    "Splinter",
		ArchitectureSqueezebert:                 "SqueezeBert",
		ArchitectureSwiftformer:                 "SwiftFormer",
		ArchitectureSwin:                        "Swin",
		ArchitectureSwin2sr:                     "Swin2SR",
		ArchitectureSwinv2:                      "Swinv2",
		ArchitectureSwitchTransformers:          "SwitchTransformers",
		ArchitectureT5:                          "T5",
		ArchitectureTableTransformer:            "TableTransformer",
		ArchitectureTapas:                       "Tapas",
		ArchitectureTimeSeriesTransformer:       "TimeSeriesTransformer",
		ArchitectureTimesformer:                 "Timesformer",
		ArchitectureTransfoXl:                   "TransfoXL",
		ArchitectureTrocr:                       "TrOCR",
		ArchitectureTvlt:                        "Tvlt",
		ArchitectureUmt5:                        "UMT5",
		ArchitectureUnispeech:                   "UniSpeech",
		ArchitectureUnispeechSat:                "UniSpeechSat",
		ArchitectureUpernet:                     "UperNet",
		ArchitectureVideomae:                    "VideoMAE",
		ArchitectureVilt:                        "Vilt",
		ArchitectureVisionEncoderDecoder:        "VisionEncoderDecoder",
		ArchitectureVisionTextDualEncoder:       "VisionTextDualEncoder",
		ArchitectureVisualBert:                  "VisualBert",
		ArchitectureVit:                         "ViT",
		ArchitectureVitHybrid:                   "ViTHybrid",
		ArchitectureVitMae:                      "ViTMAE",
		ArchitectureVitMsn:                      "ViTMSN",
		ArchitectureVitdet:                      "ViTDet",
		ArchitectureVitmatte:                    "VitMatte",
		ArchitectureVits:                        "Vits",
		ArchitectureVivit:                       "ViViT",
		ArchitectureWav2vec2:                    "Wav2Vec2",
		ArchitectureWav2vec2Conformer:           "Wav2Vec2Conformer",
		ArchitectureWavlm:                       "WavLM",
		ArchitectureWhisper:                     "Whisper",
		ArchitectureXClip:                       "XClip",
		ArchitectureXglm:                        "XGLM",
		ArchitectureXlm:                         "XLM",
		ArchitectureXlmProphetnet:               "XLMProphetNet",
		ArchitectureXlmRoberta:                  "XLMRoberta",
		ArchitectureXlmRobertaXl:                "XLMRobertaXL",
		ArchitectureXlnet:                       "XLNet",
		ArchitectureXmod:                        "Xmod",
		ArchitectureYolos:                       "Yolos",
		ArchitectureYoso:                        "Yoso",
	}
)

// String outputs the Architecture as a string.
func (a Architecture) String() string {
	return ArchitectureNames[a]
}

// MarshalJSON outputs the Architecture as a json.
func (a Architecture) MarshalJSON() ([]byte, error) {
	if !a.Valid() {
		return nil, fmt.Errorf("invalid architecture")
	}

	return []byte(`"` + a.String() + `"`), nil
}

// UnmarshalJSON parses the Architecture from json.
func (a *Architecture) UnmarshalJSON(data []byte) error {
	architecture, err := ParseArchitecture(string(data))
	if err != nil {
		return err
	}

	*a = architecture

	return nil
}

// Valid returns true if the Architecture is valid.
func (a Architecture) Valid() bool {
	return a != ArchitectureInvalid
}

// ParseArchitecture parses the Architecture string.
func ParseArchitecture(value string) (Architecture, error) {
	v := strings.ToLower(strings.Trim(value, `"`))

	for k, n := range ArchitectureNames {
		if v == strings.ToLower(n) {
			return k, nil
		}
	}

	return ArchitectureInvalid, fmt.Errorf("invalid architecture: %s", value)
}
