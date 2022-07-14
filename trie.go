package txt

type Node struct {
	Kids      map[rune]*Node
	Done      bool
	Character rune
	id        uint32
}

// func (n *Node) String() string {
// 	return fmt.Sprintf(`Node{
// 	done: %v,
// 	character: %v,
// 	id: %v
// 	}`, n.Done, string(n.Character), n.id)
// }

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
// Setting this value to -1 is equivalent to setting it to len(s), as well as the ExactContains() method.
// `cd` is used internally for recursion. Do not set it yourself.
func (n *Node) PartialContains(s string, d int, cd ...int) bool {
	if d == -1 {
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

		if len(cd) == 1 {
			cd[0]++
		} else {
			cd = make([]int, 1)
			cd[0] = 1
		}

		if len(cd) == 1 && cd[0] == d {
			return true
		}

		if len(sl) != 0 {
			return node.PartialContains(sl, d, cd...)
		} else if len(s) == 1 {
			// last character in string
			cd[0] -= 1
			return node.PartialContains(s, d, cd...)
		}

	}

	// current depth > max depth
	if len(cd) == 1 && cd[0] > d {
		return false
	}

	// reached end of string, exact match
	if n.Done {
		return true
	}

	return false
}

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
			n.Kids[rn].Done = true
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
