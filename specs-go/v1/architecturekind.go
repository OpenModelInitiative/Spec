package v1

import (
	"fmt"
	"strings"
)

// Architecture is the ArchitectureKind kind.
const (
	ArchitectureKindInvalid ArchitectureKind = iota
	ArchitectureKindAlbert
	ArchitectureKindAlign
	ArchitectureKindAltclip
	ArchitectureKindAudioSpectrogramTransformer
	ArchitectureKindAutoformer
	ArchitectureKindBark
	ArchitectureKindBart
	ArchitectureKindBeit
	ArchitectureKindBert
	ArchitectureKindBertGeneration
	ArchitectureKindBigBird
	ArchitectureKindBigbirdPegasus
	ArchitectureKindBioGpt
	ArchitectureKindBit
	ArchitectureKindBlenderbot
	ArchitectureKindBlenderbotSmall
	ArchitectureKindBlip
	ArchitectureKindBlip2
	ArchitectureKindBloom
	ArchitectureKindBridgeTower
	ArchitectureKindBros
	ArchitectureKindCamembert
	ArchitectureKindCanine
	ArchitectureKindChineseClip
	ArchitectureKindClap
	ArchitectureKindClip
	ArchitectureKindClipseg
	ArchitectureKindCodegen
	ArchitectureKindConditionalDetr
	ArchitectureKindConvbert
	ArchitectureKindConvnext
	ArchitectureKindConvnextv2
	ArchitectureKindCpmant
	ArchitectureKindCtrl
	ArchitectureKindCvt
	ArchitectureKindData2Vec
	ArchitectureKindData2VecText
	ArchitectureKindData2VecVision
	ArchitectureKindDeberta
	ArchitectureKindDebertaV2
	ArchitectureKindDecisionTransformer
	ArchitectureKindDeformableDetr
	ArchitectureKindDeit
	ArchitectureKindDeta
	ArchitectureKindDetr
	ArchitectureKindDinat
	ArchitectureKindDinov2
	ArchitectureKindDistilbert
	ArchitectureKindDpr
	ArchitectureKindDpt
	ArchitectureKindEfficientformer
	ArchitectureKindEfficientnet
	ArchitectureKindElectra
	ArchitectureKindEncodec
	ArchitectureKindEncoderDecoder
	ArchitectureKindErnie
	ArchitectureKindErnieM
	ArchitectureKindEsm
	ArchitectureKindFalcon
	ArchitectureKindFlaubert
	ArchitectureKindFlava
	ArchitectureKindFnet
	ArchitectureKindFocalnet
	ArchitectureKindFsmt
	ArchitectureKindFunnel
	ArchitectureKindGit
	ArchitectureKindGlpn
	ArchitectureKindGpt2
	ArchitectureKindGptBigcode
	ArchitectureKindGptNeo
	ArchitectureKindGptNeox
	ArchitectureKindGptNeoxJapanese
	ArchitectureKindGptj
	ArchitectureKindGptsanJapanese
	ArchitectureKindGraphormer
	ArchitectureKindGroupvit
	ArchitectureKindHubert
	ArchitectureKindIbert
	ArchitectureKindIdefics
	ArchitectureKindImagegpt
	ArchitectureKindInformer
	ArchitectureKindInstructblip
	ArchitectureKindJukebox
	ArchitectureKindLayoutlm
	ArchitectureKindLayoutlmv2
	ArchitectureKindLayoutlmv3
	ArchitectureKindLed
	ArchitectureKindLevit
	ArchitectureKindLilt
	ArchitectureKindLlama
	ArchitectureKindLongformer
	ArchitectureKindLongt5
	ArchitectureKindLuke
	ArchitectureKindLxmert
	ArchitectureKindM2m100
	ArchitectureKindMarian
	ArchitectureKindMarkuplm
	ArchitectureKindMask2former
	ArchitectureKindMaskformer
	ArchitectureKindMbart
	ArchitectureKindMega
	ArchitectureKindMegatronBert
	ArchitectureKindMgpStr
	ArchitectureKindMistral
	ArchitectureKindMobilebert
	ArchitectureKindMobilenetV1
	ArchitectureKindMobilenetV2
	ArchitectureKindMobilevit
	ArchitectureKindMobilevitv2
	ArchitectureKindMpnet
	ArchitectureKindMpt
	ArchitectureKindMra
	ArchitectureKindMt5
	ArchitectureKindMusicgen
	ArchitectureKindMvp
	ArchitectureKindNat
	ArchitectureKindNezha
	ArchitectureKindNllbMoe
	ArchitectureKindNystromformer
	ArchitectureKindOneformer
	ArchitectureKindOpenai
	ArchitectureKindOpt
	ArchitectureKindOwlv2
	ArchitectureKindOwlvit
	ArchitectureKindPegasus
	ArchitectureKindPegasusX
	ArchitectureKindPerceiver
	ArchitectureKindPersimmon
	ArchitectureKindPix2struct
	ArchitectureKindPlbart
	ArchitectureKindPoolformer
	ArchitectureKindPop2piano
	ArchitectureKindProphetnet
	ArchitectureKindPvt
	ArchitectureKindRealm
	ArchitectureKindReformer
	ArchitectureKindRegnet
	ArchitectureKindRembert
	ArchitectureKindResnet
	ArchitectureKindRoberta
	ArchitectureKindRocBert
	ArchitectureKindRoformer
	ArchitectureKindRwkv
	ArchitectureKindSam
	ArchitectureKindSegformer
	ArchitectureKindSew
	ArchitectureKindSewD
	ArchitectureKindSpeechEncoderDecoder
	ArchitectureKindSpeechToText
	ArchitectureKindSpeechToText2
	ArchitectureKindSpeecht5
	ArchitectureKindSplinter
	ArchitectureKindSqueezebert
	ArchitectureKindSwiftformer
	ArchitectureKindSwin
	ArchitectureKindSwin2sr
	ArchitectureKindSwinv2
	ArchitectureKindSwitchTransformers
	ArchitectureKindT5
	ArchitectureKindTableTransformer
	ArchitectureKindTapas
	ArchitectureKindTimeSeriesTransformer
	ArchitectureKindTimesformer
	ArchitectureKindTransfoXl
	ArchitectureKindTrocr
	ArchitectureKindTvlt
	ArchitectureKindUmt5
	ArchitectureKindUnispeech
	ArchitectureKindUnispeechSat
	ArchitectureKindUpernet
	ArchitectureKindVideomae
	ArchitectureKindVilt
	ArchitectureKindVisionEncoderDecoder
	ArchitectureKindVisionTextDualEncoder
	ArchitectureKindVisualBert
	ArchitectureKindVit
	ArchitectureKindVitHybrid
	ArchitectureKindVitMae
	ArchitectureKindVitMsn
	ArchitectureKindVitdet
	ArchitectureKindVitmatte
	ArchitectureKindVits
	ArchitectureKindVivit
	ArchitectureKindWav2vec2
	ArchitectureKindWav2vec2Conformer
	ArchitectureKindWavlm
	ArchitectureKindWhisper
	ArchitectureKindXClip
	ArchitectureKindXglm
	ArchitectureKindXlm
	ArchitectureKindXlmProphetnet
	ArchitectureKindXlmRoberta
	ArchitectureKindXlmRobertaXl
	ArchitectureKindXlnet
	ArchitectureKindXmod
	ArchitectureKindYolos
	ArchitectureKindYoso
)

var (
	// ArchitectureKindNames is a map of ArchitectureKind to string.
	ArchitectureKindNames = map[ArchitectureKind]string{
		ArchitectureKindInvalid:                     "Invalid",
		ArchitectureKindAlbert:                      "Albert",
		ArchitectureKindAlign:                       "Align",
		ArchitectureKindAltclip:                     "AltCLIP",
		ArchitectureKindAudioSpectrogramTransformer: "AST",
		ArchitectureKindAutoformer:                  "Autoformer",
		ArchitectureKindBark:                        "Bark",
		ArchitectureKindBart:                        "Bart",
		ArchitectureKindBeit:                        "Beit",
		ArchitectureKindBert:                        "Bert",
		ArchitectureKindBertGeneration:              "BertGeneration",
		ArchitectureKindBigBird:                     "BigBird",
		ArchitectureKindBigbirdPegasus:              "BigBirdPegasus",
		ArchitectureKindBioGpt:                      "BioGpt",
		ArchitectureKindBit:                         "Bit",
		ArchitectureKindBlenderbot:                  "Blenderbot",
		ArchitectureKindBlenderbotSmall:             "BlenderbotSmall",
		ArchitectureKindBlip:                        "Blip",
		ArchitectureKindBlip2:                       "Blip2",
		ArchitectureKindBloom:                       "Bloom",
		ArchitectureKindBridgeTower:                 "BridgeTower",
		ArchitectureKindBros:                        "Bros",
		ArchitectureKindCamembert:                   "Camembert",
		ArchitectureKindCanine:                      "Canine",
		ArchitectureKindChineseClip:                 "ChineseClip",
		ArchitectureKindClap:                        "Clap",
		ArchitectureKindClip:                        "CLIP",
		ArchitectureKindClipseg:                     "CLIPSeg",
		ArchitectureKindCodegen:                     "CodeGen",
		ArchitectureKindConditionalDetr:             "ConditionalDETR",
		ArchitectureKindConvbert:                    "ConvBert",
		ArchitectureKindConvnext:                    "ConvNext",
		ArchitectureKindConvnextv2:                  "ConvNextV2",
		ArchitectureKindCpmant:                      "CPMAnt",
		ArchitectureKindCtrl:                        "CTRL",
		ArchitectureKindCvt:                         "Cvt",
		ArchitectureKindData2Vec:                    "Data2Vec",
		ArchitectureKindData2VecText:                "Data2VecText",
		ArchitectureKindData2VecVision:              "Data2VecVision",
		ArchitectureKindDeberta:                     "Deberta",
		ArchitectureKindDebertaV2:                   "DebertaV2",
		ArchitectureKindDecisionTransformer:         "DecisionTransformer",
		ArchitectureKindDeformableDetr:              "DeformableDetr",
		ArchitectureKindDeit:                        "DeiT",
		ArchitectureKindDeta:                        "Deta",
		ArchitectureKindDetr:                        "Detr",
		ArchitectureKindDinat:                       "Dinat",
		ArchitectureKindDinov2:                      "Dinov2",
		ArchitectureKindDistilbert:                  "DistilBert",
		ArchitectureKindDpr:                         "DPR",
		ArchitectureKindDpt:                         "DPT",
		ArchitectureKindEfficientformer:             "EfficientFormer",
		ArchitectureKindEfficientnet:                "EfficientNet",
		ArchitectureKindElectra:                     "Electra",
		ArchitectureKindEncodec:                     "Encodec",
		ArchitectureKindEncoderDecoder:              "EncoderDecoder",
		ArchitectureKindErnie:                       "Ernie",
		ArchitectureKindErnieM:                      "ErnieM",
		ArchitectureKindEsm:                         "ESM",
		ArchitectureKindFalcon:                      "Falcon",
		ArchitectureKindFlaubert:                    "Flaubert",
		ArchitectureKindFlava:                       "Flava",
		ArchitectureKindFnet:                        "FNet",
		ArchitectureKindFocalnet:                    "FocalNet",
		ArchitectureKindFsmt:                        "FSMT",
		ArchitectureKindFunnel:                      "Funnel",
		ArchitectureKindGit:                         "Git",
		ArchitectureKindGlpn:                        "GLPN",
		ArchitectureKindGpt2:                        "GPT2",
		ArchitectureKindGptBigcode:                  "GPTBigCode",
		ArchitectureKindGptNeo:                      "GPTNeo",
		ArchitectureKindGptNeox:                     "GPTNeox",
		ArchitectureKindGptNeoxJapanese:             "GPTNeoxJapanese",
		ArchitectureKindGptj:                        "GPTJ",
		ArchitectureKindGptsanJapanese:              "GPTSanJapanese",
		ArchitectureKindGraphormer:                  "Graphormer",
		ArchitectureKindGroupvit:                    "GroupViT",
		ArchitectureKindHubert:                      "Hubert",
		ArchitectureKindIbert:                       "IBert",
		ArchitectureKindIdefics:                     "Idefics",
		ArchitectureKindImagegpt:                    "ImageGPT",
		ArchitectureKindInformer:                    "Informer",
		ArchitectureKindInstructblip:                "InstructBlip",
		ArchitectureKindJukebox:                     "Jukebox",
		ArchitectureKindLayoutlm:                    "LayoutLM",
		ArchitectureKindLayoutlmv2:                  "LayoutLMv2",
		ArchitectureKindLayoutlmv3:                  "LayoutLMv3",
		ArchitectureKindLed:                         "LED",
		ArchitectureKindLevit:                       "Levit",
		ArchitectureKindLilt:                        "Lilt",
		ArchitectureKindLlama:                       "Llama",
		ArchitectureKindLongformer:                  "Longformer",
		ArchitectureKindLongt5:                      "LongT5",
		ArchitectureKindLuke:                        "Luke",
		ArchitectureKindLxmert:                      "Lxmert",
		ArchitectureKindM2m100:                      "M2M100",
		ArchitectureKindMarian:                      "MarianMT",
		ArchitectureKindMarkuplm:                    "MarkupLM",
		ArchitectureKindMask2former:                 "Mask2Former",
		ArchitectureKindMaskformer:                  "MaskFormer",
		ArchitectureKindMbart:                       "MBart",
		ArchitectureKindMega:                        "Mega",
		ArchitectureKindMegatronBert:                "MegatronBert",
		ArchitectureKindMgpStr:                      "MGPSTR",
		ArchitectureKindMistral:                     "Mistral",
		ArchitectureKindMobilebert:                  "MobileBert",
		ArchitectureKindMobilenetV1:                 "MobileNetV1",
		ArchitectureKindMobilenetV2:                 "MobileNetV2",
		ArchitectureKindMobilevit:                   "MobileViT",
		ArchitectureKindMobilevitv2:                 "MobileViTv2",
		ArchitectureKindMpnet:                       "MPNet",
		ArchitectureKindMpt:                         "MPT",
		ArchitectureKindMra:                         "Mra",
		ArchitectureKindMt5:                         "MT5",
		ArchitectureKindMusicgen:                    "Musicgen",
		ArchitectureKindMvp:                         "Mvp",
		ArchitectureKindNat:                         "Nat",
		ArchitectureKindNezha:                       "NeZha",
		ArchitectureKindNllbMoe:                     "NllbMoe",
		ArchitectureKindNystromformer:               "Nystromformer",
		ArchitectureKindOneformer:                   "OneFormer",
		ArchitectureKindOpenai:                      "OpenAIGPTLMHead",
		ArchitectureKindOpt:                         "OPT",
		ArchitectureKindOwlv2:                       "Owlv2",
		ArchitectureKindOwlvit:                      "OwlViT",
		ArchitectureKindPegasus:                     "Pegasus",
		ArchitectureKindPegasusX:                    "PegasusX",
		ArchitectureKindPerceiver:                   "Perceiver",
		ArchitectureKindPersimmon:                   "Persimmon",
		ArchitectureKindPix2struct:                  "Pix2Struct",
		ArchitectureKindPlbart:                      "PLBart",
		ArchitectureKindPoolformer:                  "PoolFormer",
		ArchitectureKindPop2piano:                   "Pop2Piano",
		ArchitectureKindProphetnet:                  "ProphetNet",
		ArchitectureKindPvt:                         "Pvt",
		ArchitectureKindRealm:                       "RealM",
		ArchitectureKindReformer:                    "Reformer",
		ArchitectureKindRegnet:                      "RegNet",
		ArchitectureKindRembert:                     "RemBert",
		ArchitectureKindResnet:                      "ResNet",
		ArchitectureKindRoberta:                     "Roberta",
		ArchitectureKindRocBert:                     "RoCBert",
		ArchitectureKindRoformer:                    "RoFormer",
		ArchitectureKindRwkv:                        "Rwkv",
		ArchitectureKindSam:                         "Sam",
		ArchitectureKindSegformer:                   "Segformer",
		ArchitectureKindSew:                         "SEW",
		ArchitectureKindSewD:                        "SEWD",
		ArchitectureKindSpeechEncoderDecoder:        "SpeechEncoderDecoder",
		ArchitectureKindSpeechToText:                "SpeechToText",
		ArchitectureKindSpeechToText2:               "SpeechToText2",
		ArchitectureKindSpeecht5:                    "SpeechT5",
		ArchitectureKindSplinter:                    "Splinter",
		ArchitectureKindSqueezebert:                 "SqueezeBert",
		ArchitectureKindSwiftformer:                 "SwiftFormer",
		ArchitectureKindSwin:                        "Swin",
		ArchitectureKindSwin2sr:                     "Swin2SR",
		ArchitectureKindSwinv2:                      "Swinv2",
		ArchitectureKindSwitchTransformers:          "SwitchTransformers",
		ArchitectureKindT5:                          "T5",
		ArchitectureKindTableTransformer:            "TableTransformer",
		ArchitectureKindTapas:                       "Tapas",
		ArchitectureKindTimeSeriesTransformer:       "TimeSeriesTransformer",
		ArchitectureKindTimesformer:                 "Timesformer",
		ArchitectureKindTransfoXl:                   "TransfoXL",
		ArchitectureKindTrocr:                       "TrOCR",
		ArchitectureKindTvlt:                        "Tvlt",
		ArchitectureKindUmt5:                        "UMT5",
		ArchitectureKindUnispeech:                   "UniSpeech",
		ArchitectureKindUnispeechSat:                "UniSpeechSat",
		ArchitectureKindUpernet:                     "UperNet",
		ArchitectureKindVideomae:                    "VideoMAE",
		ArchitectureKindVilt:                        "Vilt",
		ArchitectureKindVisionEncoderDecoder:        "VisionEncoderDecoder",
		ArchitectureKindVisionTextDualEncoder:       "VisionTextDualEncoder",
		ArchitectureKindVisualBert:                  "VisualBert",
		ArchitectureKindVit:                         "ViT",
		ArchitectureKindVitHybrid:                   "ViTHybrid",
		ArchitectureKindVitMae:                      "ViTMAE",
		ArchitectureKindVitMsn:                      "ViTMSN",
		ArchitectureKindVitdet:                      "ViTDet",
		ArchitectureKindVitmatte:                    "VitMatte",
		ArchitectureKindVits:                        "Vits",
		ArchitectureKindVivit:                       "ViViT",
		ArchitectureKindWav2vec2:                    "Wav2Vec2",
		ArchitectureKindWav2vec2Conformer:           "Wav2Vec2Conformer",
		ArchitectureKindWavlm:                       "WavLM",
		ArchitectureKindWhisper:                     "Whisper",
		ArchitectureKindXClip:                       "XClip",
		ArchitectureKindXglm:                        "XGLM",
		ArchitectureKindXlm:                         "XLM",
		ArchitectureKindXlmProphetnet:               "XLMProphetNet",
		ArchitectureKindXlmRoberta:                  "XLMRoberta",
		ArchitectureKindXlmRobertaXl:                "XLMRobertaXL",
		ArchitectureKindXlnet:                       "XLNet",
		ArchitectureKindXmod:                        "Xmod",
		ArchitectureKindYolos:                       "Yolos",
		ArchitectureKindYoso:                        "Yoso",
	}
)

// String outputs the ArchitectureKind as a string.
func (ak ArchitectureKind) String() string {
	return ArchitectureKindNames[ak]
}

// MarshalJSON outputs the ArchitectureKind as a json.
func (ak ArchitectureKind) MarshalJSON() ([]byte, error) {
	if !ak.Valid() {
		return nil, fmt.Errorf("invalid architecture kind")
	}

	return []byte(`"` + ak.String() + `"`), nil
}

// UnmarshalJSON parses the ArchitectureKind from json.
func (ak *ArchitectureKind) UnmarshalJSON(data []byte) error {
	architectureKind, err := ParseArchitectureKind(string(data))
	if err != nil {
		return err
	}

	*ak = architectureKind

	return nil
}

// Valid returns true if the ArchitectureKind is valid.
func (ak ArchitectureKind) Valid() bool {
	return ak != ArchitectureKindInvalid
}

// ParseArchitectureKind parses the ArchitectureKind string.
func ParseArchitectureKind(value string) (ArchitectureKind, error) {
	v := strings.ToLower(strings.Trim(value, `"`))

	for k, n := range ArchitectureKindNames {
		if v == strings.ToLower(n) {
			return k, nil
		}
	}

	return ArchitectureKindInvalid, fmt.Errorf("invalid architecture kind: %s", value)
}
