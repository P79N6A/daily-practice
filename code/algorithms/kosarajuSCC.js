const dfo = require('./depthFirstOrder')

// Kosaraju强连通分量
class KosarajuSCC {
	constructor(g) {
		this.graph = g
		this.count = 0
		this.id = new Array(this.graph.V())
		this.marked = new Array(this.graph.V())
		const order = new dfo.DepthFirstOrder(this.graph.reverse())
		const reversePost = order.ReversePost()
		for (let s = reversePost.pop(); !isNaN(s); s = reversePost.pop()) {
			if (!this.marked[s]) {
				this.dfs(s)
				this.count++
			}
		}
	}

	dfs(v) {
		this.marked[v] = true
		this.id[v] = this.count
		this.graph.Adj(v).forEach((adj) => {
			if (!this.marked[adj]) {
				this.dfs(adj)
			}
		})
	}

  stronglyConnected(v, w) {
    return this.id[v] === this.id[w]
  }

	components() {
		const components = new Array(this.count)
		for (let v = 0; v < this.graph.V(); v++) {
			const c = this.id[v]
			components[c] = components[c] || []
			components[c].push(v)
		}
		return components
	}
}

module.exports = {
	KosarajuSCC,
}
