const structs = require('./structs')

// 顶点排序
class DepthFirstOrder {
	constructor(g) {
		this.graph = g
		this.marked = []
		this.pre = new structs.Queue()
		this.post = new structs.Queue()
		this.reversePost = new structs.Stack()

		for (let v = 0; v < this.graph.V(); v++) {
			if (!this.marked[v]) {
				this.dfs(v)
			}
		}
	}

	dfs(v) {
		this.pre.enqueue(v)
		this.marked[v] = true
		this.graph.Adj(v).forEach((adj) => {
			if (!this.marked[adj]) {
				this.dfs(adj)
			}
		})

		this.post.enqueue(v)
		this.reversePost.push(v)
	}

	Pre() { return this.pre }
	Post() { return this.post }
	ReversePost() { return this.reversePost }
}

module.exports = {
	DepthFirstOrder
}
