// Tokenization utilities, primarily for use as a tokenizer in a full-text search engine.
package txt

import (
	_ "embed"
	"strings"
	"unicode"
)

//go:embed dicts/colors.txt
var COLORS_DICT_EMBED string

// Loads a replacer file from a given path.
// Files must be in the format:
// original,replaced value
// with each tuple on a new line.
// The values are loaded into memory; the replaced value is assigned the `Data` field on the final
// node of a trie branch, which can be accessed using Trie.NodeAt(x).
func LoadDictionary(data string, caseSensitive bool) *Node {
	t := NewTrie()

	for _, v := range strings.Split(data, "\n") {
		tuple := strings.Split(v, ",")
		if len(tuple) < 2 {
			break
		}

		original := tuple[0]
		if !caseSensitive {
			original = strings.ToLower(original)
		}
		replaced := tuple[1]

		t.Insert(string(original), []byte(replaced))
	}

	return t
}

var (
	// English colors and their hex equivalents.
	Colors *Node = LoadDictionary(COLORS_DICT_EMBED, false)
)

// Normalizes a list of tokens.
type Tokenizer func(tokens []string) []string

var (
	// Removes stopwords from a token list using the default stopword list.
	TokenizerStopwords Tokenizer = func(tokens []string) []string { return FilterStopwords(tokens) }
	// Stems tokens using a Porter stemmer.
	TokenizerStemmer Tokenizer = func(tokens []string) []string { return StemTokens(tokens) }

	// Normalizes color names (e.g. `red`, `navy`) into their hex values.
	TokenizerColors Tokenizer = func(tokens []string) []string {
		for i, v := range tokens {
			if n, exists := Colors.Find(v); exists {
				tokens[i] = string(n.Data.([]byte))
			}
		}

		return tokens
	}
)

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
func Tokenize(text string, splitter Splitter, normalizer []Normalizer, options ...Tokenizer) []string {
	if len(text) == 0 {
		return make([]string, 0)
	}

	if len(options) == 0 {
		options = append(options, DefaultTokenizer...)
	}

	text = Normalize(text, normalizer...)
	tokens := splitter(text)

	for _, v := range options {
		tokens = v(tokens)
	}

	return tokens
}
