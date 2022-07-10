// levenshtein distance
package txt

import (
	"bytes"
	"io/ioutil"
)

// loads the dictionary words list
func loadDict() ([][]byte, error) {
	b, err := ioutil.ReadFile("./words.txt")
	if err != nil {
		return nil, nil
	}

	return bytes.Split(b, []byte("\n")), nil
}

var dict, dictErr = loadDict()

// generates a list of spelling corrections for the provided `word`
// `lim` is the maximum levenshtein distance away for a correction to be returned (inclusive)
// e.g. a correction with a LD of 3 would be returned with a limit of `3`, but a word with a LD of 4 would not
// in the return values, the `uint8` in the map corresponds to levenshtein distance of the corrected word
// todo: better, more focused results; current implementation returns many options, especially for smaller words
// could possibly be fixed by weighting spelling errors that are closer on keyboards
// e.g. with the input `vad`
// `tad` and `bad` are both options, but the "b" in `bad` is closer physically on the keyboard than the "t" in
// `tab`, and so would be the better choice
func Correct(word string, lim uint8) (map[string]uint8, error) {
	if dictErr != nil {
		return nil, dictErr
	}

	// all found matches
	matches := map[string]uint8{}

	for i := 0; i < len(dict); i++ {
		// levenshtein distance of correction
		l := Ld(word, string(dict[i]))
		if l <= lim {
			matches[string(dict[i])] = l
			lim = l
		}
	}

	return matches, nil
}

// returns the minimum of a function
func min(v ...uint8) uint8 {
	m := v[0]

	for _, k := range v {
		if k < m {
			m = k
		}
	}

	return m
}

// levenshtein distance
// based in part on https://rosettacode.org/wiki/Levenshtein_distance#Go, some modifications made to use one-dimensional array
// this version usually takes about half the time as the second version, and usually less than half the time of the first version on RosettaCode
func Ld(a, b string) uint8 {
	if a == "" {
		return uint8(len(b))
	}
	if b == "" {
		return uint8(len(a))
	}
	if a == b {
		return 0
	}

	// row is the previous row in the LD table (contains top right at current index and top left at current index - 1)
	row := make([]uint8, len(a)+1)
	for i := range row {
		row[i] = uint8(i)
	}

	// first characters aren't the same
	var current uint8

	// bottom left, starts at 1
	var bl uint8
	// go through columns first
	for i := 1; i <= len(b); i++ {
		// previous top left - used for if letters are the same
		ptl := uint8(i - 1)
		// set first value of previous row equal to ptl
		row[0] = ptl
		current = 0

		// top left
		var tl uint8
		// top right
		var tr uint8
		bl = uint8(i)

		// go through each character in the row
		for j := 1; j <= len(a); j++ {
			// set top right equal to the value at
			tr = row[j]
			tl = ptl

			// in first row of array, so top values should be equal to index of item (e.g. [0 1 2 3 4 5])
			// value of top right should then be the value of the array at the index in the current loop
			if i == 1 {
				tr = uint8(j)
			}

			// characters are the same - use previous top left value
			if a[j-1] == b[i-1] {
				current = tl
			} else {
				// characters are different - take minimum of the three, add one operation
				current = min(tl, tr, bl) + 1
			}

			// set the previous top left value equal to
			ptl = row[j]
			row[j] = current
			bl = current
		}
	}

	return uint8(current)
}
