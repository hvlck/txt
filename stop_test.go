package txt

import (
	"testing"
)

func TestContainsStopwords(t *testing.T) {
	c, err := ContainsStopwords("this is the best text")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if !c {
		t.Log("returned doesn't contain stopwords")
		t.Fail()
	}
}

func TestRemoveStopwords(t *testing.T) {
	c, err := RemoveStopwords("this is a lot of text with many stopwords that should be removed")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if c != "lot text stopwords removed" {
		t.Log("didn't properly remove all stopwords")
		t.Fail()
	}
}

func BenchmarkRemoveStopwords(b *testing.B) {
	b.SetParallelism(1)

	c, err := RemoveStopwords("this is a lot of text with many stopwords that should be removed")
	b.StopTimer()

	if err != nil || c != "lot text stopwords removed" {
		b.Fail()
	}
}
