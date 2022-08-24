package txt

import "testing"

func TestNormalize(t *testing.T) {
	tests := []string{
		"A  ",
		"!",
		"The quick brown fox jumped over the sly  dog.",
	}
	tokens := [][]string{
		{"a"},
		{},
		{"the", "quick", "brown", "fox", "jumped", "over", "the", "sly", "dog"},
	}

	for idx, v := range tests {
		n := Normalize(v)
		if len(n) != len(tokens[idx]) {
			t.Fatalf("expected %v tokens, got %v: '%v'", len(tokens[idx]), len(n), v)
		}

		for index, tok := range n {
			if tok != tokens[idx][index] {
				t.Fatalf("expected %v, got %v in '%v'", tokens[idx][index], tok, v)
			}
		}
	}
}
