package txt

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []string{
		"Simple text string",
		"a MoRe comPlicated  string   with odd  Spacing ... ",
		"  $testing str@ing !@#$%^&*()",
	}
	expected := [][]string{
		{"simple", "text", "string"},
		{"complicated", "string", "odd", "spacing"},
		{"testing", "str", "ing"},
	}

	for index, v := range tests {
		tokens := Tokenize(v, SplitNonAlphanumeric, TokenizerStopwords)
		for idx, tok := range tokens {
			expect := expected[index][idx]
			if expect != tok {
				t.Fatalf("expected '%v', got '%v' for input '%v'", expect, tok, v)
			}
		}
	}
}

func TestNormalizeColors(t *testing.T) {
	tests := []string{
		"normalize the color HotPink",
		"gold is a color",
	}
	expected := [][]string{
		{"normalize", "color", "#FF69B4"},
		{"#FFD700", "color"},
	}

	for index, v := range tests {
		tokens := Tokenize(v, SplitNonAlphanumeric, TokenizerStopwords, TokenizerColors)
		for idx, tok := range tokens {
			expect := expected[index][idx]
			if expect != tok {
				t.Fatalf("expected '%v', got '%v' for input '%v'", expect, tok, v)
			}
		}
	}
}
