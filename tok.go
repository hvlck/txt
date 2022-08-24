// Tokenization utilities, primarily for use as a tokenizer in a full-text search engine.
package txt

import (
	"strings"
	"unicode"
)

func Normalize(text string) []string {
	text = strings.ToLower(text)
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	})
}

func Tokenize(text string) []string {
	if len(text) == 0 {
		return make([]string, 0)
	}

	tokens := Normalize(text)
	tokens = FilterStopwords(tokens)
	tokens = StemTokens(tokens)

	return tokens
}
