package txt

import "testing"

func TestReadTime(t *testing.T) {
	words := map[string]int{
		"": 10,
	}

	for k, v := range words {
		if ReadTime(k) != v {
			t.Fatalf("'%s' read time not equal to %v", k, v)
		}
	}
}

func TestWords(t *testing.T) {
	examples := []string{
		"The   quick  brown fox jumped over the   sly dog",
		"the-rest-of-the-word",
		"This contains two sentences. This is the second!",
	}

	words := [][]string{
		{"The", "quick", "brown", "fox", "jumped", "over", "the", "sly", "dog"},
		{"the", "rest", "of", "the", "word"},
		{"This", "contains", "two", "sentences", "This", "is", "the", "second"},
	}

	for idx, v := range examples {
		for index, word := range Words(v) {
			if words[idx][index] != word {
				t.Fatalf("expected %v, got %v", words[idx][index], word)
			}
		}
	}
}
