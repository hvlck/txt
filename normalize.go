// Text normalization.
package txt

import (
	"strings"
)

// Normalizes a given string.
type Normalizer func(text string) string


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
