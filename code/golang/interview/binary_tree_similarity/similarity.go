package kata

type Value int

type Node struct {
	value       Value
	left, right *Node
}

func (self *Node) IsSimilarTo(another *Node) bool {
	if self == nil && another == nil {
		return true
	} else if self == nil || another == nil {
		return false
	} else {
		return self.left.IsSimilarTo(another.left) &&
			self.right.IsSimilarTo(another.right)
	}
}

func NewTree(v Value, l *Node, r *Node) *Node {
	return &Node{
		value: v,
		left:  l,
		right: r,
	}
}
