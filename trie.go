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

// At returns the end node of the last provided string. If no node exists, then the second argument will be `false`.
func (n *Node) At(s string) (*Node, bool) {
	if len(s) == 0 {
		if node, ok := n.Kids['*']; ok {
			return node, true
		}
	}

	for _, char := range s {
		if node, ok := n.Kids[char]; ok {
			if len(s) == 1 {
				return node.At("")
			} else {
				return node.At(s[1:])
			}
		} else {
			return n, false
		}

	}

	return n, false
}

// Finds the first node in the last branch of a given string `s`. This is used by the delete function to find the earliest
// place that a branch can be removed.
// For instance, to remove the word "testing" from the following tree:
// t
//
//	e
//	 s
//	  t
//	   (i)
//	    n
//	     g
//	      *
//	   y
//	    *
//	   *
//
// The algorithm finds (i) is the first entry in the final unique branch of testing, and thus can be removed.
// This method also handles edge cases, like when a prefix of another word (e.g. `test` in this case) is being removed,
// in which case only the wordstop (`*`) child node needs to be removed. This is the meaning of the second boolean return
// parameter; a `true` value indicates only the `*` child should be removed, while a false value indicates that all
// remaining children (which should only be one) should be removed. The Node pointer return value indicates the parent
// node, whose children should be removed.
func (n *Node) find_last_unique_branch(s string, lastBranch *Node) (*Node, bool) {
	for _, char := range s {
		if node, ok := n.Kids[char]; ok {
			if len(node.Kids) == 1 {
				if _, ok := node.Kids['*']; ok {
					return lastBranch, false
				}
			} else if len(node.Kids) > 1 {
				if len(s) > 1 {
					if kid, ok := node.Kids[rune(s[1])]; ok {
						return node.find_last_unique_branch(s[1:], kid)
					}
				} else {
					if _, ok := node.Kids['*']; ok {
						return node, true
					}

					return lastBranch, false
				}
			}

			return node.find_last_unique_branch(s[1:], lastBranch)
		} else {
			return lastBranch, false
		}
	}

	return n, false
}

func (n *Node) String() string {
	return fmt.Sprintf(`Node{
	done: %v,
	character: %v,
	id: %v
}`, n.Done, string(n.Character), n.Id)
}

// ExactContains determines whether the provided string is entirely within the trie.
func (n *Node) Contains(s string) bool {
	// empty root node or string is empty
	if (len(n.Kids) == 0 && n.Id == 0) || len(s) == 0 {
		return false
	}

	rn := rune(s[0])
	if node, ok := n.Kids[rn]; ok {
		sl := s[1:]
		if len(sl) != 0 {
			return node.Contains(sl)
		} else if len(s) == 1 {
			// last character in string
			if len(node.Kids) == 1 {
				if final, ok := node.Kids['*']; ok {
					if final.Character == '*' && final.Done {
						return true
					}
				}
			}

			return node.Contains(s)
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
func (n *Node) FuzzyContains(s string, d int) bool {
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
			return node.FuzzyContains(sl, d)
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
func (n *Node) Insert(s string, data []byte) {
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
		n.Done = false
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

// Delete removes words from a trie.
func (n *Node) Delete(words ...string) bool {
	// Root node is empty or no words provided
	if len(n.Kids) == 0 && n.Id == 0 || len(words) == 0 {
		return false
	}

	for _, s := range words {
		if len(s) == 0 {
			continue
		}

		if node, deleteFinal := n.find_last_unique_branch(s, n); node != nil {
			if deleteFinal {
				delete(node.Kids, '*')
			} else {
				// # of kids is guaranteed to be 1
				for k := range node.Kids {
					delete(node.Kids, k)
				}
			}
		}
	}

	return false
}

// NewTrie creates a new trie. Note that this function only creates a root node.
func NewTrie() *Node {
	return &Node{
		Kids:      make(map[rune]*Node),
		Done:      true,
		Character: '*',
		Id:        0,
	}
}
