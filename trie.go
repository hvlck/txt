package txt

import (
	"fmt"
	"unicode"
)

// Node is a node in a trie tree.
type Node struct {
	// Pointers to child branches
	Kids map[rune]*Node
	// Custom data, inserted into the last child ('*') of a string's tree.
	Data []byte
	// End of branch
	Done bool
	// Character representing current branch
	Character rune
	// Internal id. 0 is the id of the root node.
	Id uint32
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

func (n *Node) String() string {
	return fmt.Sprintf(`Node{
	done: %v,
	character: %v,
	id: %v
	}`, n.Done, string(n.Character), n.Id)
}

// ExactContains determines whether the provided string is entirely within the trie.
func (n *Node) ExactContains(s string) bool {
	// empty root node or string is empty
	if (len(n.Kids) == 0 && n.Id == 0) || len(s) == 0 {
		return false
	}

	rn := rune(s[0])
	if node, ok := n.Kids[rn]; ok {
		sl := s[1:]
		if len(sl) != 0 {
			return node.ExactContains(sl)
		} else if len(s) == 1 {
			// last character in string
			if len(node.Kids) == 1 {
				if final, ok := node.Kids['*']; ok {
					if final.Character == '*' && final.Done {
						return true
					}
				}
			}

			return node.ExactContains(s)
		}
	}

	return false
}

// PartialContains checks if the provided string is completely within the tree.
// `d` is an optional depth value that controls how many characters of the provided string must be present
// sequentially in the tree. For example, providing (`exam`, 3) for a trie that already has `example` will
// check to make sure that `e`, `x`, and `a` are in the trie as children of the previous character.
// Setting this value to -1 or a value greater than the length of `s` is equivalent to setting it to len(s),
// as well as the ExactContains() method. Note that this method only searches for substrings at the beginning of a
// word.
func (n *Node) PartialContains(s string, d int) bool {
	if d == -1 {
		d = len(s)
	} else if d > len(s) {
		d = len(s)
	}

	// root node is empty or string is empty
	if (len(n.Kids) == 0 && n.Id == 0) || len(s) == 0 {
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
			return true
		}
	}

	return false
}

// todo: may be issue with this in global scope and having multiple tries
// Global counter for giving trie nodes an id.
var c uint32

// Creates a new tree node.
func newNode(rn rune) *Node {
	c++
	return &Node{
		Kids:      make(map[rune]*Node),
		Done:      false,
		Character: rn,
		Id:        c,
	}
}

// Inserts a word into a trie.
func (n Node) Insert(s string, data []byte) {
	if len(s) == 0 {
		return
	}

	rn := rune(s[0])
	if !unicode.IsNumber(rn) && !unicode.IsLetter(rn) {
		return
	}

	// no kids
	if len(n.Kids) == 0 {
		n.Kids[rn] = newNode(rn)
	} else if node, ok := n.Kids[rn]; ok {
		// node is present, continue down branch
		sl := s[1:]
		if len(sl) != 0 {
			node.Insert(sl, data)
			return
		}
	} else if !ok {
		// character is not present on end of branch, create new node
		n.Kids[rn] = newNode(rn)
	}

	sl := s[1:]
	if len(sl) != 0 {
		n.Kids[rn].Insert(sl, data)
	}

	// last child
	if len(s) == 1 {
		n.Kids[rn].Kids['*'] = newNode('*')
		n.Kids[rn].Kids['*'].Done = true
		n.Kids[rn].Kids['*'].Data = data
	}
}

func (n Node) Delete(words ...string) bool {
	if len(n.Kids) == 0 && n.Id == 0 || len(words) == 0 {
		return false
	}

	for _, s := range words {
		if len(s) == 0 {
			continue
		}

		// todo: finish implementation
	}

	return false
}

func NewTrie() *Node {
	return &Node{
		Kids:      make(map[rune]*Node),
		Done:      true,
		Character: '*',
		Id:        0,
	}
}
