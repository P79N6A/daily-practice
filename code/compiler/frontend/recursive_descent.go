package frontend

import (
	"fmt"
	"unicode"
)

func newNode(label rune, left, right *Node) *Node {
	return &Node{Label: label, LeftChild: left, RightChild: right}
}
func NewParser(expression string) *binaryTreeParser {
	return &binaryTreeParser{expression, 0}
}

func (node *Node) String() string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf("%c(%s,%s)", node.Label, node.LeftChild, node.RightChild)
}

type Node struct {
	Label      rune
	LeftChild  *Node
	RightChild *Node
}

//  N → a(N,N)
//  N → ε
type binaryTreeParser struct {
	inputString string
	index       int
}

func (tree *binaryTreeParser) peekChar() rune {
	return rune(tree.inputString[tree.index])
}
func (tree *binaryTreeParser) incrIndex() {
	tree.index++
}
func (tree *binaryTreeParser) ParseNode() *Node {
	peekChar := tree.peekChar()
	if unicode.IsLetter(peekChar) { // N → a(N,N)
		tree.incrIndex() // (
		tree.incrIndex() // N
		left := tree.ParseNode()
		tree.incrIndex() // ,
		right := tree.ParseNode()
		tree.incrIndex() // N
		return newNode(peekChar, left, right)
	} else if peekChar == ',' || peekChar == ')' {
		return nil
	} else {
		panic("syntax error.")
	}
}
