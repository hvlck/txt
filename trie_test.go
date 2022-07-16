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

func TestPrefixLength(t *testing.T) {
	vals := []uint8{
		PrefixLength("tree", "trees"),
		PrefixLength("grant", "grace"),
		PrefixLength("hammer", "hankering"),
	}
	answers := []uint8{4, 3, 2}

	for i, v := range vals {
		if v != answers[i] {
			t.Fatal(v, answers[i], i)
		}
	}
}

func TestMax(t *testing.T) {

}

func TestAbs(t *testing.T) {
	nth := abs(10 - 23)
	th := abs(10 + 3)

	if nth != 13 || th != 13 {
		t.Fail()
	}
}

func TestKeyProximity(t *testing.T) {
	vals := []uint8{
		KeyProximity('r', 't'),
		KeyProximity('s', 'w'),
		KeyProximity('a', 'w'),
		KeyProximity('l', 'p'),
		KeyProximity('v', 'p'),
		KeyProximity('1', '.'),
	}
	answers := []uint8{1, 1, 1, 1, 6, 7}

	for i, v := range vals {
		if v != answers[i] {
			t.Fatal(v, answers[i])
		}
	}
}

func TestPartialMatch(t *testing.T) {
	trie, err := loadTrie()
	if err != nil {
		t.Fatal(err)
	}

	matches := trie.PartialMatch("tesk", 3, 15)
	if len(matches) != 15 {
		t.Fail()
	}
}

func BenchmarkPartialMatch(b *testing.B) {
	b.SetParallelism(1)
	b.StopTimer()
	trie, err := loadTrie()
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()

	matches := trie.PartialMatch("tesk", 3, 15)
	if len(matches) != 15 {
		b.Fail()
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

func BenchmarkInsert(b *testing.B) {
	trie := NewTrie()
	trie.Insert("testing")
	if len(trie.Kids) != 1 {
		b.Fail()
	}
}

func TestInsert(t *testing.T) {
	trie := NewTrie()
	trie.Insert("testing")
	if len(trie.Kids) != 1 {
		t.Fail()
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("testing", "original", "tertiary")

	printKids(trie, "", 1)
}
