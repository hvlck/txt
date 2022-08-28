// Tokenization utilities, primarily for use as a tokenizer in a full-text search engine.
package txt

import (
	"bytes"
	"os"
	"strings"
	"unicode"
)

// Normalizes a list of tokens.
type Tokenizer func(tokens []string) []string

var (
	// Removes stopwords from a token list using the default stopword list.
	TokenizerStopwords Tokenizer = func(tokens []string) []string { return FilterStopwords(tokens) }
	// Stems tokens using a Porter stemmer.
	TokenizerStemmer Tokenizer = func(tokens []string) []string { return StemTokens(tokens) }


var (
	// Splits a string at non-alphanumeric characters (whitespace, punctuation, etc).
	SplitNonAlphanumeric Splitter = func(text string) []string {
		return strings.FieldsFunc(text, func(r rune) bool {
			return !unicode.IsNumber(r) && !unicode.IsLetter(r)
		})
	}

	DefaultSplitter Splitter = SplitNonAlphanumeric
)

// DefaultTokenizer removes stopwords and stems tokens.
var DefaultTokenizer = []Tokenizer{
	TokenizerStopwords,
	TokenizerStemmer,
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
