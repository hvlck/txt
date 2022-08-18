package txt

import (
	"testing"
)

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

func TestWordOffsets(t *testing.T) {
	examples := []string{
		"The quick brown fox",
	}

	words := [][]int{
		{0, 4, 10, 16},
	}

	for idx, v := range examples {
		for index, word := range WordOffsets(v) {
			if words[idx][index] != word.offset {
				t.Fatalf("expected %v, got %v", words[idx][index], word)
			}
		}
	}
}

func TestWordFrequency(t *testing.T) {
	examples := []string{
		"The quick brown fox ran over the slow gray fox.",
	}

	words := [][]string{
		{"the", "fox", "ran"},
	}

	freqs := [][]uint{
		{2, 2, 1},
	}

	for idx, v := range examples {
		for index, word := range words[idx] {
			if WordFrequency(v, word) != freqs[idx][index] {
				t.Fatalf("expected %v, got %v", freqs[idx][index], WordFrequency(v, word))
			}
		}
	}
}
