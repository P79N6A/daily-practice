const println = console.log

function Btree(value, left, right) {
    this.value = value
    this.left = left
    this.right = right
}

function ReverseTree(tree) {
    if (!tree) {
        return
    }
    t = tree.left
    tree.left = tree.right
    tree.right = t

    ReverseTree(tree.left)
	ReverseTree(tree.right)
}

const tree = new Btree(10, new Btree(20, new Btree(30)), new Btree(30, null, new Btree(30)))
ReverseTree(tree)
println(tree)
