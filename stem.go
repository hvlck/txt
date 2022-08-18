// Porter stemming
// this is a port of https://medium.com/analytics-vidhya/building-a-stemmer-492e9a128e84 as I am not familiar
// with the Porter stemming algorithmn
// some fixes were made as many of the test words failed at various stages
// see https://tartarus.org/martin/PorterStemmer/def.txt for the spec
package txt

import (
	"strings"
)

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'y': true,
}

// assumes input is lowercase
func isVowel(s rune) bool {
	return vowels[s]
}

// tests if the provided input is a list of consonants or vowels
// true is vowel, false is consonant
func stringIsVowel(s string) bool {
	return vowels[rune(s[0])]
}

func m(s string) uint8 {
	var m uint8 = 0
	enc := encode(groupWord(s))

	if len(enc) > 2 {
		if enc[0] == 'C' {
			enc = enc[1:]
		}

		if enc[len(enc)-1] == 'V' {
			enc = enc[:len(enc)-1]
		}

		// may technically be wrong, result should be floored
		if len(enc)/2 >= 1 {
			m = uint8(len(enc) / 2)
		}
	}

	return m
}

// groups a word by consonants and vowels
// "testing" -> ["t", "e", "st", "i", "ng"]
func groupWord(s string) []string {
	f := strings.ToLower(s)
	offset := 0
	lastWasVowel := false

	for i, v := range s {
		iv := isVowel(v)

		// last was consonant and current is vowel or last was vowel and current is consonant
		if (!lastWasVowel && iv) || (lastWasVowel && !iv) {
			lastWasVowel = iv
			r := i + offset
			if i == 0 {
				continue
			}

			if r < len(f) {
				f = f[:r] + "," + f[r:]
				offset += 1
			}
		}
	}

	return strings.Split(f, ",")
}

type cvList []rune
type stringList []string

func encode(s []string) cvList {
	r := make([]rune, len(s))
	for i, v := range s {
		vow := isVowel(rune(v[0]))
		if vow {
			r[i] = 'V'
		} else {
			r[i] = 'C'
		}
	}

	return r
}

// stem has vowel
func stemContainsVowel(stem string) bool {
	for _, v := range stem {
		if isVowel(v) {
			return true
		}
	}

	return false
}

// stem ends with double consonant
func stemEndsCC(s string) bool {
	o := s[len(s)-1]
	t := s[len(s)-2]

	if !stringIsVowel(string(o)) && !stringIsVowel(string(t)) {
		return true
	}

	return false
}

// stem ends with consonant-vowel-consonant
// last consonant cannot be W, X, or Y
func stemEndsCVC(s stringList) bool {
	if len(s) < 3 {
		return false
	}

	one := s[len(s)-1]
	two := s[len(s)-2]
	three := s[len(s)-3]

	if !stringIsVowel(one) && stringIsVowel(two) && !stringIsVowel(three) {
		if !strings.Contains(one, "w") && !strings.Contains(one, "x") && !strings.Contains(one, "y") {
			return true
		}
	}

	return false
}

func Stem(s string) string {
	// don't care, it's beautiful
	return porter_five(porter_four(porter_three(porter_two(porter_one(s)))))
}

// alias for strings.HasSuffix
// will only return true if any provided suffix is present
func suf(s string, suffix ...string) bool {
	for _, v := range suffix {
		if strings.HasSuffix(s, v) {
			return true
		}
	}

	return false
}

func porter_one(word string) string {
	enc := groupWord(word)
	s := word
	oneB := false

	if suf(s, "sses") {
		s = s[:len(s)-2]
	} else if suf(s, "ies") {
		s = s[:len(s)-2]
	} else if !suf(s, "ss") && suf(s, "s") {
		s = s[:len(s)-1]
	}

	if len(s) > 4 {
		if suf(s, "eed") && m(s) > 0 {
			s = s[:len(s)-1]
		} else if suf(s, "ed") {
			s = s[:len(s)-2]
			if !stemContainsVowel(s) {
				s = word
			} else {
				oneB = true
			}
		} else if suf(s, "ing") {
			s = s[:len(s)-3]
			if !stemContainsVowel(s) {
				s = word
			} else {
				oneB = true
			}
		}
	}

	if oneB {
		if suf(s, "at", "bl", "iz") {
			s += "e"
		} else if stemEndsCC(s) && !suf(s, "lsz") {
			s = s[:len(s)-1]
		} else if m(s) == 1 && stemEndsCVC(enc) {
			s += "e"
		}
	}

	if stemContainsVowel(s) && suf(s, "y") {
		s = s[:len(s)-1] + "i"
	}

	return s
}

var twoTerms = []string{
	"ational",
	"tional",
	"enci",
	"anci",
	"izer",
	"abli",
	"alli",
	"entli",
	"eli",
	"ousli",
	"ization",
	"ation",
	"ator",
	"alism",
	"iveness",
	"fulness",
	"ousness",
	"aliti",
	"iviti",
}

var twoSubs = []string{
	"ate",
	"tion",
	"ence",
	"ance",
	"ize",
	"able",
	"ali",
	"ent",
	"e",
	"ous",
	"ize",
	"ate",
	"ate",
	"al",
	"ive",
	"ful",
	"ous",
	"al",
	"ive",
}

func porter_two(word string) string {
	if m(word) > 0 {
		for i, v := range twoTerms {
			if suf(word, v) {
				r := word[:len(word)-len(v)]
				if m(r) > 0 {
					word = r + twoSubs[i]
				}
			}
		}
	}

	return word
}

var threeTerms = []string{
	"icate",
	"ative",
	"alize",
	"iciti",
	"ical",
	"ful",
	"ness",
}

var threeSubs = []string{
	"ic",
	"",
	"al",
	"ic",
	"ic",
	"",
	"",
}

func porter_three(word string) string {
	if m(word) > 0 {
		for i, v := range threeTerms {
			if suf(word, v) {
				return word[:len(word)-len(v)] + threeSubs[i]
			}
		}
	}

	return word
}

var fourSuffixes = []string{
	"al", "ance", "ence", "er", "ic", "able", "ible", "ant", "ement", "ment", "ent", "ou", "ism", "ate", "iti", "ous", "ive", "ize",
}

func porter_four(word string) string {
	if suf(word, "ion") {
		s := word[:len(word)-3]
		if suf(s, "s", "t") {
			return s
		}
	}

	for _, v := range fourSuffixes {
		if suf(word, v) {
			sl := word[:len(word)-len(v)]
			measure := m(sl)
			if measure > 1 {
				return sl
			}
		}
	}
	return word
}

func porter_five(word string) string {
	g := groupWord(word)
	mw := m(word)
	// 5a
	if mw > 1 && suf(word, "e") {
		word = word[:len(word)-1]
	} else if mw == 1 && !stemEndsCVC(g) && suf(word, "e") && len(word) > 4 {
		word = word[:len(word)-1]
	}
	// 5b
	if mw > 1 && stemEndsCC(word) && suf(word, "l") {
		word = word[:len(word)-1]
	}

	return word
}
