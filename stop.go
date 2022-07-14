// utilities for managing stopwords
// stopword list courtesy of https://dev.mysql.com/doc/refman/8.0/en/fulltext-stopwords.html#fulltext-stopwords-stopwords-for-myisam-search-indexes
// our version is somewhat modified
package txt

import (
	"bytes"
	"io/ioutil"
	"strings"
)

var stopwords, stopErr = loadStopwords()

// loads the stopwords list
func loadStopwords() (map[string]bool, error) {
	b, err := ioutil.ReadFile("./dicts/stop.txt")
	if err != nil {
		return nil, nil
	}

	bytes := bytes.Split(b, []byte(","))
	m := make(map[string]bool, len(bytes))
	for _, v := range bytes {
		m[string(v)] = true
	}

	return m, nil
}

func ContainsStopwords(str string) (bool, error) {
	if stopErr != nil {
		return false, stopErr
	}

	m := false
	w := strings.Split(str, " ")

	for _, v := range w {
		if _, ok := stopwords[v]; ok {
			return true, nil
		}
	}

	return m, nil
}

func removeAtIndex(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)
}

func RemoveStopwords(str string) (string, error) {
	str = strings.ReplaceAll(strings.ToLower(str), "\n", "")
	if stopErr != nil {
		return "", stopErr
	}

	w := strings.Split(str, " ")

	for i := 0; i < len(w); i++ {
		v := w[i]
		if stopwords[v] {
			w = removeAtIndex(w, i)
			i--
		}
	}

	return strings.Join(w, " "), nil
}
