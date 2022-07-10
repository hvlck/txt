// tries
// https://en.wikipedia.org/wiki/Trie
package txt

import "fmt"

type Trie struct {
	root *Branch
}

func MakeTrie() *Trie {
	return &Trie{
		root: &Branch{
			Branches: make(map[byte]*Branch),
		},
	}
}

func (t *Trie) String() string {
	return fmt.Sprintf(`
Trie {
	root: %v
}
	`, t.root)
}

func (t *Trie) Add(s string) {
	b := []byte(s)
	rt := t.root

	if len(rt.Branches) == 0 {

	}

	t.root.Set(b)
}

type Branch struct {
	Branches    map[byte]*Branch
	Value       []byte
	EndOfBranch bool
	Id          uint32
}

func (b *Branch) String() string {
	return fmt.Sprintf(`Branch {
		branches: %v,
		value: %v,
		end: %v,
		id: %v
	}
	`, b.Branches, string(b.Value), b.EndOfBranch, b.Id)
}

func (b *Branch) Set(s []byte) {
	if len(b.Value) == 0 {
		b.Value = s
	}
}
