package txt

import "strings"

// Average words read per minute
var READING_WPM = 200

// Duration, in seconds, to read the provided string `s`.
func ReadTime(s string) int {
	return len(Words(s)) / READING_WPM
}

// Counts the number of words within a string.
func WordCount(s string, spaces bool) int {
	return len(Words(s))
}

// Replaces instances of multiple spaces with a single space.
func replaceSpaces(s string) string {
	s = strings.ReplaceAll(s, "  ", " ")
	if strings.Contains(s, "  ") {
		return replaceSpaces(s)
	}

	return s
}

// Removes the provided characters from a string.
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
	s = Normalize(s, NormalizeSingleSpace, NormalizeHyphens, NormalizeAllPunctuation)
	return strings.Split(s, " ")
}

// Occurence of a word in a string, as well as its offset within the string.
// todo: documentation on exact offset meaning
type WordOffset struct {
	Word   string
	Offset int
}

// Returns the offsets and individual words in a string.
func WordOffsets(s string) []WordOffset {
	s = Normalize(s)
	words := strings.Split(s, " ")
	offsets := make([]WordOffset, len(words))

	for idx, v := range words {
		offsets[idx] = WordOffset{
			Word:   v,
			Offset: strings.Index(s, v),
		}
	}

	return offsets
}

func WordFrequency(s, word string) uint {
	words := Tokenize(s, DefaultSplitter)

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
