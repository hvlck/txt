package txt

import (
	"fmt"
	"strings"
	"testing"
)

func printKids(t *Node, parent string, d int) {
	if t.Id == 0 {
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
	words := []string{"testing", "original", "tertiary"}
	for _, v := range words {
		trie.Insert(v, nil)
	}

	answers := map[string]bool{
		"original": true,
		"test":     false,
		"turtle":   false,
	}

	for k, v := range answers {
		if trie.Contains(k) != v {
			t.Fatalf("%v is not %v", k, v)
		}
	}
}

func BenchmarkExactContains(b *testing.B) {
	b.StopTimer()
	trie := NewTrie()
	words := []string{"Antidisestablishmentarianism", "original", "tertiary"}
	for _, v := range words {
		trie.Insert(v, nil)
	}
	b.StartTimer()

	present := trie.Contains("Antidisestablishmentarianism")

	if !present {
		b.Fail()
	}
}

func TestFuzzyContains(t *testing.T) {
	trie := NewTrie()
	words := []string{"testing", "original", "tertiary"}
	for _, v := range words {
		trie.Insert(v, nil)
	}

	answers := map[string]bool{
		"original": true,
		"test":     true,
		"turtle":   false,
		"tert":     true,
		"nal":      false,
	}

	for k, v := range answers {
		if trie.FuzzyContains(k, -1) != v {
			t.Fatalf("%v is not %v", k, v)
		}
	}
}

func BenchmarkPartialContains(b *testing.B) {
	trie := NewTrie()
	words := []string{"testing", "original", "tertiary", "Antidisestablishmentarianism"}
	for _, v := range words {
		trie.Insert(v, nil)
	}

	if trie.FuzzyContains("Antidisestablishm", -1) != true {
		b.Fail()
	}
}

func TestInsert(t *testing.T) {
	trie := NewTrie()
	trie.Insert("testing", nil)
	if len(trie.Kids) != 1 {
		t.Fail()
	}

	if _, ok := trie.At("testing"); !ok {
		t.Fail()
	}
}

func BenchmarkInsert(b *testing.B) {
	trie := NewTrie()
	trie.Insert("testing", nil)
	if len(trie.Kids) != 1 {
		b.Fail()
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	words := []string{"testing", "original", "tertiary"}
	for _, v := range words {
		trie.Insert(v, nil)
	}

	// printKids(trie, "", 1)
}

func TestDelete(t *testing.T) {
	trie := NewTrie()
	trie.Insert("testing", nil)
	trie.Insert("tertiary", nil)
	trie.Insert("test", nil)
	trie.Insert("testy", nil)

	trie.Delete("testing")

	if _, ok := trie.At("testing"); ok {
		t.Fail()
	}

	if node, ok := trie.At("test"); !ok {
		t.Fatal(node)
	}

	if _, ok := trie.At("testy"); !ok {
		t.Fail()
	}

	if _, ok := trie.At("tertiary"); !ok {
		t.Fail()
	}
}
