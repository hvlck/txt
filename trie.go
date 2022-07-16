package txt

import (
	"sort"
	"strings"
	"unicode"
)

// Node is a node in a trie tree.
type Node struct {
	// Pointers to child branches
	Kids map[rune]*Node
	// End of branch
	Done bool
	// Character representing current branch
	Character rune
	// Internal id. 0 is the id of the root node.
	id uint32
}

// A word correction. A copy of the original word is not stored.
type Correction struct {
	// Corrected word
	Word string
	// Levenshtein distance from original. Lower is closer.
	ld uint8
	// Number of characters that both words share at the beginning.
	// For example, grace and grant have a prefix_len of 3 as they both share `gra` at the beginning.
	// Higher is better.
	prefix_len uint8
	// Sum of the distance between each character in the original and corrected word. Lower is better.
	key_len uint8
	// Weight of word correction. Higher values mean the correction is closer to the original word.
	Weight float32
}

// Searches for all words in the trie within a fixed `limit` edit distance away from the original string `s`.
func (n *Node) search_lev(s, b string, limit uint8, prev ...Correction) []Correction {
	if n.id == 0 {
		for rn, v := range n.Kids {
			prev = append(prev, v.search_lev(s, string(rn), limit)...)
		}
		return prev
	} else {
		for rn, v := range n.Kids {
			lev := Ld(b, s)
			// if lev > limit {
			// 	continue
			// }

			if v.Done && len(v.Kids) == 0 {
				if lev <= limit {
					prev = append(prev, Correction{ld: lev, Word: b, Weight: 0})
				}

				continue
			} else {
				prev = append(prev, v.search_lev(s, b+string(rn), limit)...)
			}
		}
	}

	return prev
}

// PrefixLength calculates the number of same characters at the beginning of both strings.
func PrefixLength(o, t string) uint8 {
	var n uint8 = 0
	for i, v := range t {
		if len(o)-1 < i {
			return n
		}

		if v == rune(o[i]) {
			n++
		} else {
			break
		}
	}

	return n
}

var keys = [][]rune{
	{'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
	{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']', '\\'},
	{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\'', ' ', ' '},
	{'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/', ' ', ' ', ' '},
}

var all_keys = make([]rune, 0, 13*4)

func abs(x int) int {
	y := 0
	if x < y {
		return y - x
	}
	return x - y
}

func max(x, y int) uint8 {
	if x > y {
		return uint8(x)
	} else {
		return uint8(y)
	}
}

// Returns the number of keys away `t` is from `o`.
// This is used as a measure of accidental typos, e.g. `jat` when the intention was `hat`.
func KeyProximity(o, t rune) uint8 {
	if o == t {
		return 0
	}

	if len(all_keys) == 0 {
		for _, v := range keys {
			all_keys = append(all_keys, v...)
		}
	}

	rO := 0
	cO := 0

	rT := 0
	cT := 0
	for idx, v := range all_keys {
		idx += 1
		if v == o {
			cO = idx / 13
			rO = idx - cO*13
		}

		if v == t {
			cT = idx / 13
			rT = idx - cT*13
		}
	}

	rowDiff := abs(rT - rO)
	colDiff := abs(cT - cO)

	// largest value, no trig
	return max(colDiff, rowDiff)
}

func (n *Node) PartialMatch(s string, target uint8, max int) []Correction {
	f := n.search_lev(strings.ToLower(s), "", target)

	var lim float32 = 0
	res := make([]Correction, max)

	last := 0
	for _, v := range f {
		v.weigh(s)
		if lim == 0 {
			lim = v.Weight
		}

		if v.Weight >= lim {
			// res is filled
			if n := res[last]; len(n.Word) != 0 && n.ld != 0 {
				// search for element with lowest weight, replace it
				for i, k := range res {
					// levenshtein and weight of word being examined is less than word currently in final results
					if v.Weight > k.Weight && v.ld <= target {
						res[i] = v
						sort.Slice(res, func(i, j int) bool {
							return res[i].Weight < res[j].Weight
						})
						break
					}
				}
			} else if last < max {
				res[last] = v
				if last+1 < max {
					last++
				}
				sort.Slice(res, func(i, j int) bool {
					return res[i].Weight < res[j].Weight
				})
			}
		}
		// ignore words with weights smaller than limit
	}

	return res
}

// NodeAt returns the node at the last character of the provided string.
func (n *Node) NodeAt(s string) *Node {
	for _, char := range s {
		if _, ok := n.Kids[char]; ok {
			return n.Kids[char].NodeAt(s[1:])
		}
	}

	return n
}

// func (n *Node) String() string {
// 	return fmt.Sprintf(`Node{
// 	done: %v,
// 	character: %v,
// 	id: %v
// 	}`, n.Done, string(n.Character), n.id)
// }

// ExactContains determines whether the provided string is entirely within the trie.
func (n *Node) ExactContains(s string) bool {
	// root node
	if len(n.Kids) == 0 && n.id == 0 {
		return false
	}

	rn := rune(s[0])
	if node, ok := n.Kids[rn]; ok {
		sl := s[1:]
		if len(sl) != 0 {
			return node.ExactContains(sl)
		} else if len(s) == 1 {
			// last character in string
			return node.ExactContains(s)
		}
	}

	if n.Done {
		return true
	}

	return false
}

// PartialContains checks if the provided string is completely within the tree.
// `d` is an optional depth value that controls how many characters of the provided string must be present
// sequentially in the tree. For example, providing (`exam`, 3) for a trie that already has `example` will
// check to make sure that `e`, `x`, and `a` are in the trie as children of the previous character.
// Setting this value to -1 or a value greater than the length of `s` is equivalent to setting it to len(s),
// as well as the ExactContains() method.
func (n *Node) PartialContains(s string, d int) bool {
	if d == -1 {
		d = len(s)
	} else if d > len(s) {
		d = len(s)
	}

	// root node
	if len(n.Kids) == 0 && n.id == 0 {
		return false
	}

	rn := rune(s[0])
	// next character is present as child of current character node
	if node, ok := n.Kids[rn]; ok {
		sl := s[1:]

		if d == 0 {
			return true
		}

		d--
		if len(sl) != 0 {
			return node.PartialContains(sl, d)
		} else if len(s) == 1 {
			// last character in string
			return node.PartialContains(s, d)
		}

	}

	// current depth > max depth
	if d == -1 {
		return false
	}

	// reached end of string, exact match
	if n.Done {
		return true
	}

	return false
}

// todo: may be issue with this in global scope and having multiple tries
var c uint32

func NewNode(rn rune) *Node {
	c++
	return &Node{
		Kids:      make(map[rune]*Node),
		Done:      false,
		Character: rn,
		id:        c,
	}
}

// inserts words into a trie
func (n Node) Insert(words ...string) {
	for _, s := range words {
		if len(s) == 0 {
			break
		}

		rn := rune(s[0])
		if !unicode.IsLetter(rn) {
			continue
		}

		// no kids
		if len(n.Kids) == 0 {
			n.Kids[rn] = NewNode(rn)
		} else if node, ok := n.Kids[rn]; ok {
			// node is present, continue down branch
			sl := s[1:]
			if len(sl) != 0 {
				node.Insert(sl)
				continue
			}
		} else if !ok {
			// character is not present on end of branch, create new node
			n.Kids[rn] = NewNode(rn)
		}

		sl := s[1:]
		if len(sl) != 0 {
			n.Kids[rn].Insert(sl)
		}

		// last child
		if len(s) == 1 {
			n.Kids[rn].Kids['*'] = NewNode('*')
			n.Kids[rn].Kids['*'].Done = true
		}
	}
}

func NewTrie() *Node {
	return &Node{
		Kids:      make(map[rune]*Node),
		Done:      true,
		Character: '*',
		id:        0,
	}
}
