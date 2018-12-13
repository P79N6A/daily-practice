package main

import "fmt"

var (
	tree = &Tree{"a",
		&Tree{"b",
			&Tree{"d", nil, nil},
			nil},
		&Tree{"c",
			nil,
			&Tree{"e", nil, nil},
		},
	}
)

// Tree : binary tree
type Tree struct {
	value string
	left  *Tree
	right *Tree
}

func (tree *Tree) String() string {
	if tree == nil {
		return ""
	}
	return fmt.Sprintf("%s{%s,%s}",
		tree.value, tree.left, tree.right)
}

// Reverse binary tree
func Reverse(tree *Tree) {
	if tree == nil {
		return
	}

	t := tree.left
	tree.left = tree.right
	tree.right = t

	Reverse(tree.left)
	Reverse(tree.right)
}

func main() {
	fmt.Println("[ORIGINAL]:\t\t", tree)
	Reverse(tree)
	fmt.Println("[FLIPPED]:\t\t", tree)
}
