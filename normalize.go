// Text normalization.
package txt

import (
	"strings"
)

// Normalizes a given string.
type Normalizer func(text string) string

var (
	// Converts all characters to lowercase.
	NormalizerToLower Normalizer = func(text string) string { return strings.ToLower(text) }

	// Converts excess whitespace into a single space.
	NormalizeSingleSpace Normalizer = func(text string) string {
		return replaceSpaces(text)
	}

	// Removes most punctuation. Periods and slashes are not removed.
	NormalizePunctuation Normalizer = func(text string) string {
		return removeChars(text, ",", ":", ";", "!", "?", "\"", "'", "(", ")", "&")
	}

	// Removes all punctuation, including periods and slashes.
	NormalizeAllPunctuation Normalizer = func(text string) string {
		return removeChars(NormalizePunctuation(text), ".", "/", "\\")
	}

	NormalizeSpecial Normalizer = func(text string) string {
		return removeChars(text, "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "[", "]", "{", "}", ";", ":", "'", "\"", "<", ">", "?", "/", "\\", "~", "`")
	}

	// Replaces instances of hyphens with spaces.
	NormalizeHyphens Normalizer = func(text string) string {
		return strings.ReplaceAll(text, "-", " ")
	}
)

// Default normalizer used if no normalizers are provided.
var DefaultNormalizer = []Normalizer{
	NormalizerToLower,
	NormalizePunctuation,
	NormalizeSingleSpace,
	NormalizeHyphens,
	NormalizeSpecial,
}

// Normalizes a given string using the provided options. Options are executed in the order that they're provided in.
// If no options are provided, the default normalizer is used.
func Normalize(text string, options ...Normalizer) string {
	if len(options) == 0 {
		options = append(options, DefaultNormalizer...)
	}

	for _, v := range options {
		text = v(text)
	}

	return text
}
