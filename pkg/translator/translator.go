package translator

type (
	Language   string
	Translator interface {
		Translate(key string, lang Language) string
	}
)
