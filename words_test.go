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