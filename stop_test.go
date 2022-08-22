package txt

import (
	"testing"
)

func TestContainsStopwords(t *testing.T) {
	c := ContainsStopwords("this is the best text")

	if !c {
		t.Log("returned doesn't contain stopwords")
		t.Fail()
	}
}

func TestRemoveStopwords(t *testing.T) {
	c := RemoveStopwords("this is a lot of text with many stopwords that should be removed")

	if c != "lot text stopwords removed" {
		t.Log("didn't properly remove all stopwords")
		t.Fail()
	}
}

func BenchmarkRemoveStopwords(b *testing.B) {
	b.SetParallelism(1)

	c := RemoveStopwords("this is a lot of text with many stopwords that should be removed")
	b.StopTimer()

	if c != "lot text stopwords removed" {
		b.Fail()
	}
}

func TestFilterStopwords(t *testing.T) {
	words := []string{"This", "is", "a", "test"}
	filtered := FilterStopwords(words)

	if len(filtered) != 1 {
		t.Fatal()
	}
}
