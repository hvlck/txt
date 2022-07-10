package txt

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := MakeTrie()
	trie.Add("test")
}
