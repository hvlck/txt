package txt

import "testing"

func TestNormalize(t *testing.T) {
	tests := []string{
		"A  ",
		"!",
		"The quick brown fox jumped over the sly  dog.",
		"!@#$%^&*()-+=_[]{}\"';:<>?/\\.,`~",
	}

	expected := []string{
		"a ",
		"",
		"the quick brown fox jumped over the sly dog.",
		"$ :/\\.",
	}

	for idx, v := range tests {
		n := Normalize(v)

		if n != expected[idx] {
			t.Fatalf("expected '%v', got '%v' in '%v'", expected[idx], n, v)
		}
	}
}
