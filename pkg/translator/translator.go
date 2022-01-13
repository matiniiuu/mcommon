package translator

import (
	"strings"
)

type (
	Translator interface {
		Translate(key string, lang ...Language) string
	}

	Language string
)

const (
	AR Language = "ar"
	EN Language = "en"
	FA Language = "fa"
	TR Language = "tr"
)

func GetLanguage(lang string) Language {
	switch strings.ToLower(lang) {
	case "ar", "ar-ar":
		return AR
	case "en", "en-us":
		return EN
	case "fa", "fa-ir":
		return FA
	case "tr", "tr-tr":
		return TR
	default:
		return EN
	}
}
