package v1

import (
	"fmt"
	"strings"
)

// @see https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes

// Language is a model natural language (ISO 639-1)
const (
	LanguageInvalid Language = iota //
	LanguageAa
	LanguageAb
	LanguageAe
	LanguageAf
	LanguageAk
	LanguageAm
	LanguageAn
	LanguageAr
	LanguageAs
	LanguageAv
	LanguageAy
	LanguageAz
	LanguageBa
	LanguageBe
	LanguageBg
	LanguageBh
	LanguageBi
	LanguageBm
	LanguageBn
	LanguageBo
	LanguageBr
	LanguageBs
	LanguageCa
	LanguageCe
	LanguageCh
	LanguageCo
	LanguageCr
	LanguageCs
	LanguageCu
	LanguageCv
	LanguageCy
	LanguageDa
	LanguageDe
	LanguageDv
	LanguageDz
	LanguageEe
	LanguageEl
	LanguageEn
	LanguageEo
	LanguageEs
	LanguageEt
	LanguageEu
	LanguageFa
	LanguageFf
	LanguageFi
	LanguageFj
	LanguageFo
	LanguageFr
	LanguageFy
	LanguageGa
	LanguageGd
	LanguageGl
	LanguageGn
	LanguageGu
	LanguageGv
	LanguageHa
	LanguageHe
	LanguageHi
	LanguageHo
	LanguageHr
	LanguageHt
	LanguageHu
	LanguageHy
	LanguageHz
	LanguageIa
	LanguageId
	LanguageIe
	LanguageIg
	LanguageIi
	LanguageIk
	LanguageIo
	LanguageIs
	LanguageIt
	LanguageIu
	LanguageJa
	LanguageJv
	LanguageKa
	LanguageKg
	LanguageKi
	LanguageKj
	LanguageKk
	LanguageKl
	LanguageKm
	LanguageKn
	LanguageKo
	LanguageKr
	LanguageKs
	LanguageKu
	LanguageKv
	LanguageKw
	LanguageKy
	LanguageLa
	LanguageLb
	LanguageLg
	LanguageLi
	LanguageLn
	LanguageLo
	LanguageLt
	LanguageLu
	LanguageLv
	LanguageMg
	LanguageMh
	LanguageMi
	LanguageMk
	LanguageMl
	LanguageMn
	LanguageMr
	LanguageMs
	LanguageMt
	LanguageMy
	LanguageNa
	LanguageNb
	LanguageNd
	LanguageNe
	LanguageNg
	LanguageNl
	LanguageNn
	LanguageNo
	LanguageNr
	LanguageNv
	LanguageNy
	LanguageOc
	LanguageOj
	LanguageOm
	LanguageOr
	LanguageOs
	LanguagePa
	LanguagePi
	LanguagePl
	LanguagePs
	LanguagePt
	LanguageQu
	LanguageRm
	LanguageRn
	LanguageRo
	LanguageRu
	LanguageRw
	LanguageSa
	LanguageSc
	LanguageSd
	LanguageSe
	LanguageSg
	LanguageSi
	LanguageSk
	LanguageSl
	LanguageSm
	LanguageSn
	LanguageSo
	LanguageSq
	LanguageSr
	LanguageSs
	LanguageSt
	LanguageSu
	LanguageSv
	LanguageSw
	LanguageTa
	LanguageTe
	LanguageTg
	LanguageTh
	LanguageTi
	LanguageTk
	LanguageTl
	LanguageTn
	LanguageTo
	LanguageTr
	LanguageTs
	LanguageTt
	LanguageTw
	LanguageTy
	LanguageUg
	LanguageUk
	LanguageUr
	LanguageUz
	LanguageVe
	LanguageVi
	LanguageVo
	LanguageWa
	LanguageWo
	LanguageXh
	LanguageYi
	LanguageYo
	LanguageZa
	LanguageZh
	LanguageZu
)

var (
	// LanguageCodes is a map of Language to code.
	LanguageCodes = map[Language]string{
		LanguageInvalid: "invalid",
		LanguageAa:      "aa",
		LanguageAb:      "ab",
		LanguageAe:      "ae",
		LanguageAf:      "af",
		LanguageAk:      "ak",
		LanguageAm:      "am",
		LanguageAn:      "an",
		LanguageAr:      "ar",
		LanguageAs:      "as",
		LanguageAv:      "av",
		LanguageAy:      "ay",
		LanguageAz:      "az",
		LanguageBa:      "ba",
		LanguageBe:      "be",
		LanguageBg:      "bg",
		LanguageBh:      "bh",
		LanguageBi:      "bi",
		LanguageBm:      "bm",
		LanguageBn:      "bn",
		LanguageBo:      "bo",
		LanguageBr:      "br",
		LanguageBs:      "bs",
		LanguageCa:      "ca",
		LanguageCe:      "ce",
		LanguageCh:      "ch",
		LanguageCo:      "co",
		LanguageCr:      "cr",
		LanguageCs:      "cs",
		LanguageCu:      "cu",
		LanguageCv:      "cv",
		LanguageCy:      "cy",
		LanguageDa:      "da",
		LanguageDe:      "de",
		LanguageDv:      "dv",
		LanguageDz:      "dz",
		LanguageEe:      "ee",
		LanguageEl:      "el",
		LanguageEn:      "en",
		LanguageEo:      "eo",
		LanguageEs:      "es",
		LanguageEt:      "et",
		LanguageEu:      "eu",
		LanguageFa:      "fa",
		LanguageFf:      "ff",
		LanguageFi:      "fi",
		LanguageFj:      "fj",
		LanguageFo:      "fo",
		LanguageFr:      "fr",
		LanguageFy:      "fy",
		LanguageGa:      "ga",
		LanguageGd:      "gd",
		LanguageGl:      "gl",
		LanguageGn:      "gn",
		LanguageGu:      "gu",
		LanguageGv:      "gv",
		LanguageHa:      "ha",
		LanguageHe:      "he",
		LanguageHi:      "hi",
		LanguageHo:      "ho",
		LanguageHr:      "hr",
		LanguageHt:      "ht",
		LanguageHu:      "hu",
		LanguageHy:      "hy",
		LanguageHz:      "hz",
		LanguageIa:      "ia",
		LanguageId:      "id",
		LanguageIe:      "ie",
		LanguageIg:      "ig",
		LanguageIi:      "ii",
		LanguageIk:      "ik",
		LanguageIo:      "io",
		LanguageIs:      "is",
		LanguageIt:      "it",
		LanguageIu:      "iu",
		LanguageJa:      "ja",
		LanguageJv:      "jv",
		LanguageKa:      "ka",
		LanguageKg:      "kg",
		LanguageKi:      "ki",
		LanguageKj:      "kj",
		LanguageKk:      "kk",
		LanguageKl:      "kl",
		LanguageKm:      "km",
		LanguageKn:      "kn",
		LanguageKo:      "ko",
		LanguageKr:      "kr",
		LanguageKs:      "ks",
		LanguageKu:      "ku",
		LanguageKv:      "kv",
		LanguageKw:      "kw",
		LanguageKy:      "ky",
		LanguageLa:      "la",
		LanguageLb:      "lb",
		LanguageLg:      "lg",
		LanguageLi:      "li",
		LanguageLn:      "ln",
		LanguageLo:      "lo",
		LanguageLt:      "lt",
		LanguageLu:      "lu",
		LanguageLv:      "lv",
		LanguageMg:      "mg",
		LanguageMh:      "mh",
		LanguageMi:      "mi",
		LanguageMk:      "mk",
		LanguageMl:      "ml",
		LanguageMn:      "mn",
		LanguageMr:      "mr",
		LanguageMs:      "ms",
		LanguageMt:      "mt",
		LanguageMy:      "my",
		LanguageNa:      "na",
		LanguageNb:      "nb",
		LanguageNd:      "nd",
		LanguageNe:      "ne",
		LanguageNg:      "ng",
		LanguageNl:      "nl",
		LanguageNn:      "nn",
		LanguageNo:      "no",
		LanguageNr:      "nr",
		LanguageNv:      "nv",
		LanguageNy:      "ny",
		LanguageOc:      "oc",
		LanguageOj:      "oj",
		LanguageOm:      "om",
		LanguageOr:      "or",
		LanguageOs:      "os",
		LanguagePa:      "pa",
		LanguagePi:      "pi",
		LanguagePl:      "pl",
		LanguagePs:      "ps",
		LanguagePt:      "pt",
		LanguageQu:      "qu",
		LanguageRm:      "rm",
		LanguageRn:      "rn",
		LanguageRo:      "ro",
		LanguageRu:      "ru",
		LanguageRw:      "rw",
		LanguageSa:      "sa",
		LanguageSc:      "sc",
		LanguageSd:      "sd",
		LanguageSe:      "se",
		LanguageSg:      "sg",
		LanguageSi:      "si",
		LanguageSk:      "sk",
		LanguageSl:      "sl",
		LanguageSm:      "sm",
		LanguageSn:      "sn",
		LanguageSo:      "so",
		LanguageSq:      "sq",
		LanguageSr:      "sr",
		LanguageSs:      "ss",
		LanguageSt:      "st",
		LanguageSu:      "su",
		LanguageSv:      "sv",
		LanguageSw:      "sw",
		LanguageTa:      "ta",
		LanguageTe:      "te",
		LanguageTg:      "tg",
		LanguageTh:      "th",
		LanguageTi:      "ti",
		LanguageTk:      "tk",
		LanguageTl:      "tl",
		LanguageTn:      "tn",
		LanguageTo:      "to",
		LanguageTr:      "tr",
		LanguageTs:      "ts",
		LanguageTt:      "tt",
		LanguageTw:      "tw",
		LanguageTy:      "ty",
		LanguageUg:      "ug",
		LanguageUk:      "uk",
		LanguageUr:      "ur",
		LanguageUz:      "uz",
		LanguageVe:      "ve",
		LanguageVi:      "vi",
		LanguageVo:      "vo",
		LanguageWa:      "wa",
		LanguageWo:      "wo",
		LanguageXh:      "xh",
		LanguageYi:      "yi",
		LanguageYo:      "yo",
		LanguageZa:      "za",
		LanguageZh:      "zh",
		LanguageZu:      "zu",
	}

	// LanguageNames is a map of Language to string.
	LanguageNames = map[Language]string{
		LanguageInvalid: "Invalid",
		LanguageAa:      "Afar",
		LanguageAb:      "Abkhazian",
		LanguageAe:      "Avestan",
		LanguageAf:      "Afrikaans",
		LanguageAk:      "Akan",
		LanguageAm:      "Amharic",
		LanguageAn:      "Aragonese",
		LanguageAr:      "Arabic",
		LanguageAs:      "Assamese",
		LanguageAv:      "Avaric",
		LanguageAy:      "Aymara",
		LanguageAz:      "Azerbaijani",
		LanguageBa:      "Bashkir",
		LanguageBe:      "Belarusian",
		LanguageBg:      "Bulgarian",
		LanguageBh:      "Bihari",
		LanguageBi:      "Bislama",
		LanguageBm:      "Bambara",
		LanguageBn:      "Bengali",
		LanguageBo:      "Tibetan",
		LanguageBr:      "Breton",
		LanguageBs:      "Bosnian",
		LanguageCa:      "Catalan",
		LanguageCe:      "Chechen",
		LanguageCh:      "Chamorro",
		LanguageCo:      "Corsican",
		LanguageCr:      "Cree",
		LanguageCs:      "Czech",
		LanguageCu:      "Church Slavic",
		LanguageCv:      "Chuvash",
		LanguageCy:      "Welsh",
		LanguageDa:      "Danish",
		LanguageDe:      "German",
		LanguageDv:      "Divehi",
		LanguageDz:      "Dzongkha",
		LanguageEe:      "Ewe",
		LanguageEl:      "Greek",
		LanguageEn:      "English",
		LanguageEo:      "Esperanto",
		LanguageEs:      "Spanish",
		LanguageEt:      "Estonian",
		LanguageEu:      "Basque",
		LanguageFa:      "Persian",
		LanguageFf:      "Fulah",
		LanguageFi:      "Finnish",
		LanguageFj:      "Fijian",
		LanguageFo:      "Faroese",
		LanguageFr:      "French",
		LanguageFy:      "Western Frisian",
		LanguageGa:      "Irish",
		LanguageGd:      "Scottish Gaelic",
		LanguageGl:      "Galician",
		LanguageGn:      "Guarani",
		LanguageGu:      "Gujarati",
		LanguageGv:      "Manx",
		LanguageHa:      "Hausa",
		LanguageHe:      "Hebrew",
		LanguageHi:      "Hindi",
		LanguageHo:      "Hiri Motu",
		LanguageHr:      "Croatian",
		LanguageHt:      "Haitian",
		LanguageHu:      "Hungarian",
		LanguageHy:      "Armenian",
		LanguageHz:      "Herero",
		LanguageIa:      "Interlingua",
		LanguageId:      "Indonesian",
		LanguageIe:      "Interlingue",
		LanguageIg:      "Igbo",
		LanguageIi:      "Sichuan Yi",
		LanguageIk:      "Inupiaq",
		LanguageIo:      "Ido",
		LanguageIs:      "Icelandic",
		LanguageIt:      "Italian",
		LanguageIu:      "Inuktitut",
		LanguageJa:      "Japanese",
		LanguageJv:      "Javanese",
		LanguageKa:      "Georgian",
		LanguageKg:      "Kongo",
		LanguageKi:      "Kikuyu",
		LanguageKj:      "Kuanyama",
		LanguageKk:      "Kazakh",
		LanguageKl:      "Kalaallisut",
		LanguageKm:      "Central Khmer",
		LanguageKn:      "Kannada",
		LanguageKo:      "Korean",
		LanguageKr:      "Kanuri",
		LanguageKs:      "Kashmiri",
		LanguageKu:      "Kurdish",
		LanguageKv:      "Komi",
		LanguageKw:      "Cornish",
		LanguageKy:      "Kirghiz",
		LanguageLa:      "Latin",
		LanguageLb:      "Luxembourgish",
		LanguageLg:      "Ganda",
		LanguageLi:      "Limburgan",
		LanguageLn:      "Lingala",
		LanguageLo:      "Lao",
		LanguageLt:      "Lithuanian",
		LanguageLu:      "Luba-Katanga",
		LanguageLv:      "Latvian",
		LanguageMg:      "Malagasy",
		LanguageMh:      "Marshallese",
		LanguageMi:      "Maori",
		LanguageMk:      "Macedonian",
		LanguageMl:      "Malayalam",
		LanguageMn:      "Mongolian",
		LanguageMr:      "Marathi",
		LanguageMs:      "Malay",
		LanguageMt:      "Maltese",
		LanguageMy:      "Burmese",
		LanguageNa:      "Nauru",
		LanguageNb:      "Norwegian Bokmal",
		LanguageNd:      "North Ndebele",
		LanguageNe:      "Nepali",
		LanguageNg:      "Ndonga",
		LanguageNl:      "Dutch",
		LanguageNn:      "Norwegian Nynorsk",
		LanguageNo:      "Norwegian",
		LanguageNr:      "South Ndebele",
		LanguageNv:      "Navajo",
		LanguageNy:      "Chichewa",
		LanguageOc:      "Occitan",
		LanguageOj:      "Ojibwa",
		LanguageOm:      "Oromo",
		LanguageOr:      "Oriya",
		LanguageOs:      "Ossetian",
		LanguagePa:      "Panjabi",
		LanguagePi:      "Pali",
		LanguagePl:      "Polish",
		LanguagePs:      "Pushto",
		LanguagePt:      "Portuguese",
		LanguageQu:      "Quechua",
		LanguageRm:      "Romansh",
		LanguageRn:      "Rundi",
		LanguageRo:      "Romanian",
		LanguageRu:      "Russian",
		LanguageRw:      "Kinyarwanda",
		LanguageSa:      "Sanskrit",
		LanguageSc:      "Sardinian",
		LanguageSd:      "Sindhi",
		LanguageSe:      "Northern Sami",
		LanguageSg:      "Sango",
		LanguageSi:      "Sinhala",
		LanguageSk:      "Slovak",
		LanguageSl:      "Slovenian",
		LanguageSm:      "Samoan",
		LanguageSn:      "Shona",
		LanguageSo:      "Somali",
		LanguageSq:      "Albanian",
		LanguageSr:      "Serbian",
		LanguageSs:      "Swati",
		LanguageSt:      "Southern Sotho",
		LanguageSu:      "Sundanese",
		LanguageSv:      "Swedish",
		LanguageSw:      "Swahili",
		LanguageTa:      "Tamil",
		LanguageTe:      "Telugu",
		LanguageTg:      "Tajik",
		LanguageTh:      "Thai",
		LanguageTi:      "Tigrinya",
		LanguageTk:      "Turkmen",
		LanguageTl:      "Tagalog",
		LanguageTn:      "Tswana",
		LanguageTo:      "Tonga",
		LanguageTr:      "Turkish",
		LanguageTs:      "Tsonga",
		LanguageTt:      "Tatar",
		LanguageTw:      "Twi",
		LanguageTy:      "Tahitian",
		LanguageUg:      "Uighur",
		LanguageUk:      "Ukrainian",
		LanguageUr:      "Urdu",
		LanguageUz:      "Uzbek",
		LanguageVe:      "Venda",
		LanguageVi:      "Vietnamese",
		LanguageVo:      "Volapuk",
		LanguageWa:      "Walloon",
		LanguageWo:      "Wolof",
		LanguageXh:      "Xhosa",
		LanguageYi:      "Yiddish",
		LanguageYo:      "Yoruba",
		LanguageZa:      "Zhuang",
		LanguageZh:      "Chinese",
		LanguageZu:      "Zulu",
	}
)

// String outputs the Language as a string.
func (l Language) String() string {
	return LanguageNames[l]
}

// Code outputs the Language code.
func (l Language) Code() string {
	return LanguageCodes[l]
}

// MarshalJSON outputs the Language as a json.
func (l Language) MarshalJSON() ([]byte, error) {
	if !l.Valid() {
		return nil, fmt.Errorf("invalid language")
	}

	return []byte(`"` + l.String() + `"`), nil
}

// UnmarshalJSON parses the Language from json.
func (l *Language) UnmarshalJSON(data []byte) error {
	language, err := ParseLanguage(string(data))
	if err != nil {
		return err
	}

	*l = language

	return nil
}

// Valid returns true if the Language is valid.
func (l Language) Valid() bool {
	return l != LanguageInvalid
}

// ParseLanguage parses the Language string.
func ParseLanguage(value string) (Language, error) {
	v := strings.ToLower(strings.Trim(value, `"`))

	for k, n := range LanguageNames {
		if v == strings.ToLower(n) {
			return k, nil
		}
	}

	return LanguageInvalid, fmt.Errorf("invalid language: %s", value)
}
