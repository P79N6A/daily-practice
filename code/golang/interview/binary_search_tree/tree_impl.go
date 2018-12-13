package kata

type Key interface{}
type Value interface{}
type Node struct {
	key         Key
	value       Value
	left, right *Node
	count       uint
}

type Tree struct {
	root    *Node
	compare func(i, j Key) int
}

func (t *Tree) get(n *Node, k Key) Value {
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
func (t *Tree) put(n *Node, k Key, v Value) *Node {
	if n == nil {
		return node(k, v)
	}
	cmp := t.compare(k, n.key)
	if cmp < 0 {
		n.left = t.put(n.left, k, v)
	} else if cmp > 0 {
		n.right = t.put(n.right, k, v)
	} else {
		n.value = v
	}
	n.count = t.size(n.left) + t.size(n.right) + 1
	return n
}
func (t *Tree) min(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return t.min(n.left)
}
func (t *Tree) size(n *Node) uint {
	if n == nil {
		return 0
	}
	return n.count
}
func (t *Tree) deleteMin(n *Node) *Node {
	if n.left == nil {
		return n.right
	}
	n.left = t.deleteMin(n.left)
	n.count = t.size(n.left) + t.size(n.right) + 1
	return n
}
func (t *Tree) delete(n *Node, key Key) *Node {
	if n == nil {
		return nil
	}
	cmp := t.compare(key, n.key)
	if cmp > 0 {
		n.left = t.delete(n.left, key)
	} else if cmp < 0 {
		n.right = t.delete(n.right, key)
	} else {
		if n.right == nil {
			return n.left
		}
		if n.left == nil {
			return n.right
		}
		tmp := n
		n = t.min(tmp.right)
		n.right = t.deleteMin(tmp.right)
		n.left = tmp.left
	}
	n.count = t.size(n.left) + t.size(n.right) + 1
	return n
}
func (t *Tree) Put(k Key, v Value) *Node {
	return t.put(t.root, k, v)
}
func (t *Tree) Min() Key {
	return t.min(t.root)
}
func (t *Tree) Delete(key Key) {
	t.root = t.delete(t.root, key)
}

func node(k Key, v Value) *Node {
	return &Node{
		key:   k,
		value: v,
		count: 1,
	}
}
func InitTree(k Key, v Value, compare func(i, j Key) int) *Tree {
	t := &Tree{
		root:    node(k, v),
		compare: compare,
	}
	return t
}
