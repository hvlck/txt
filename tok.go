// Tokenization utilities, primarily for use as a tokenizer in a full-text search engine.
package txt

import (
	"strings"
	"unicode"
)

// Normalizes a list of tokens.
type Tokenizer func(tokens []string) []string
}

// Splits a string into individual tokens.
type Splitter func(text string) []string

// Produces a list of normalized text tokens. If no options are provided, the DefaultTokenizer is used.
func Tokenize(text string, splitter Splitter, options ...Tokenizer) []string {
	if len(text) == 0 {
		return make([]string, 0)
	}

	if len(options) == 0 {
		options = append(options, DefaultTokenizer...)
	}

	text = Normalize(text)
	tokens := splitter(text)

	for _, v := range options {
		tokens = v(tokens)
	}

	return tokens
}
