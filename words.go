package txt

import "strings"

// Average words read per minute
var READING_WPM = 200

// Duration, in seconds, to read the provided string `s`.
func ReadTime(s string) int {
	return len(Words(s)) / READING_WPM
}

func WordCount(s string, spaces bool) int {
	return len(Words(s))
}

func replaceSpaces(s string) string {
	s = strings.ReplaceAll(s, "  ", " ")
	if strings.Contains(s, "  ") {
		return replaceSpaces(s)
	}

	return s
}

func removeChars(s string, characters ...string) string {
	for _, v := range characters {
		s = strings.ReplaceAll(s, v, "")
	}
	return s
}

// Returns the individual words in a string, in lowercase.
// Also removes punctuation (".", ",", ":", ";", "!", "?"). Parantheses are kept, as well as brackets and quotation
// marks.
func Words(s string) []string {
	s = normalize(s)
	return strings.Split(s, " ")
}

type WordOffset struct {
	word   string
	offset int
}

func normalize(s string) string {
	s = strings.ToLower(s)
	s = replaceSpaces(s)
	s = strings.ReplaceAll(s, "-", " ")
	s = removeChars(s, ".", ",", ":", ";", "!", "?")

	return s
}

// Returns the offsets and individual words in a string.
func WordOffsets(s string) []WordOffset {
	s = normalize(s)
	words := strings.Split(s, " ")
	offsets := make([]WordOffset, len(words))

	for idx, v := range words {
		offsets[idx] = WordOffset{
			word:   v,
			offset: strings.Index(s, v),
		}
	}

	return offsets
}

func WordFrequency(s, word string) uint {
	words := Words(s)

	var i uint
	for _, v := range words {
		if v == word {
			i += 1
		}
	}

	return i
}

func CharCount(s string) int {
	return len(s)
}
