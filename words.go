package txt

import "strings"

// Average words read per minute
var READING_WPM = 200

// Duration, in seconds, to read the provided string `s`.
func ReadTime(s string) int {
	return len(strings.Split(s, " ")) / READING_WPM
}

func WordCount(s string, spaces bool) int {
	return len(strings.Split(s, " "))
}

func CharCount(s string) int {
	return len(s)
}
