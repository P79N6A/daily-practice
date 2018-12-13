package kata

type Color uint
type Key interface{}
type Value interface{}

const (
	RED Color = iota
	BLACK
)

type Node struct {
	key   Key
	value Value
	left  *Node
	right *Node
	count uint
	color Color
}

func (n *Node) setColor(color Color) {
	if n != nil {
		n.color = color
	}
}

func (n *Node) isRed() bool {
	return n.color == RED
}

func (n *Node) flipColors() {
	n.setColor(RED)
	n.left.setColor(BLACK)
	n.right.setColor(BLACK)
}

func size(n *Node) uint {
	if n == nil {
		return 0
	}
	return n.count
}

type RedBlackTree struct {
	root    *Node
	compare func(i, j Key) int
}

func (t *RedBlackTree) rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	x.count = h.count
	h.count = size(h.left) + size(h.right) + 1
	return x
}
func (t *RedBlackTree) rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	x.count = h.count
	h.count = size(h.left) + size(h.right) + 1
	return x
}

func (t *RedBlackTree) put(h *Node, k Key, v Value) *Node {
	if h == nil {
		return &Node{
			key:   k,
			value: v,
			count: 1,
			color: RED,
		}
	}

	cmp := t.compare(k, h.key)
	if cmp < 0 {
		h.left = t.put(h.left, k, v)
	} else if cmp > 0 {
		h.right = t.put(h.right, k, v)
	} else {
		h.value = v
	}

	if h.right.isRed() && !h.left.isRed() {
		h = t.rotateLeft(h)
	} else if h.left != nil && h.left.isRed() && h.left.left.isRed() {
		h = t.rotateRight(h)
	} else if h.left.isRed() && h.right.isRed() {
		h.flipColors()
	}

	h.count = size(h.left) + size(h.right) + 1
	return h
}

func (t *RedBlackTree) get(n *Node, k Key) Value {
	if n == nil {
		return nil
	}
	cmp := t.compare(k, n.key)
	if cmp < 0 {
		return t.get(n.left, k)
	} else if cmp > 0 {
		return t.get(n.right, k)
	} else {
		return n.value
	}
}

func (t *RedBlackTree) Put(k Key, v Value) {
	t.root = t.put(t.root, k, v)
	t.root.color = BLACK
}

func (t *RedBlackTree) Get(k Key) Value {
	return t.get(t.root, k)
}

func NewTree(compare func(i, j Key) int) *RedBlackTree {
	return &RedBlackTree{
		compare: compare,
	}
}
