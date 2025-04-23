package mutils

import "fmt"

func FormatLocalizedField(filed string, language *string, extra ...string) string {
	if language == nil {
		language = StringPointer("en")
	}
	formattedFiled := fmt.Sprintf("$%s.%s", filed, *language)
	for _, v := range extra {
		formattedFiled = fmt.Sprintf("%s.%s", formattedFiled, v)
	}
	return formattedFiled
}
