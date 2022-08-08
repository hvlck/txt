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

// Returns the individual words in a string.
// Also removes punctuation (".", ",", ":", ";", "!", "?"). Parantheses are kept, as well as brackets and quotation
// marks.
func Words(s string) []string {
	s = replaceSpaces(s)
	s = strings.ReplaceAll(s, "-", " ")
	s = removeChars(s, ".", ",", ":", ";", "!", "?")
	return strings.Split(s, " ")
}

func CharCount(s string) int {
	return len(s)
}
