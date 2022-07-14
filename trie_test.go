package txt

import (
	"fmt"
	"strings"
	"testing"
)

func printKids(t *Node, parent string, d int) {
	if t.id == 0 {
		fmt.Println("* (root, depth: 0)")
	}

	for k, v := range t.Kids {
		fmt.Printf("%v%v (parent: %v, depth: %v): %v\n", strings.Repeat(" ", d*2), string(k), parent, d, v)

		if len(v.Kids) != 0 {
			printKids(v, string(k), d+1)
		}
	}
}

func TestExactContains(t *testing.T) {
	trie := NewTrie()
	trie.Insert("testing", "original", "tertiary")

	present := trie.ExactContains("original")
	not_present := trie.ExactContains("test")

	if !present || not_present {
		t.Fail()
	}
}

func TestPartialContains(t *testing.T) {
	trie := NewTrie()
	trie.Insert("testing", "original", "tertiary")

	partial := trie.PartialContains("test", 3)
	full := trie.PartialContains("testing", -1)
	not_present := trie.PartialContains("oil", -1)
	if !partial || !full || not_present {
		t.Fatalf("%v,%v,%v", partial, full, not_present)
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("testing", "original", "tertiary")

	printKids(trie, "", 1)
}
