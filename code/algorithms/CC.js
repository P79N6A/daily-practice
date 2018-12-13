// 连通分量
class CC {
	constructor(g) {
		this.graph = g
		this.id = []
		this.marked = []
		this.count = 0
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

	components() {
		this.id = new Array(this.graph.V())
		this.marked = new Array(this.graph.V())
		this.count = 0
		for (let v = 0; v < this.graph.V(); v++) {
			if (!this.marked[v]) {
				this.dfs(v)
				this.count++
			}
		}

		if (this.count) {
			const components = new Array(this.count)
			for (let v = 0; v < this.graph.V(); v++) {
				const c = this.id[v]
				components[c] = components[c] || []
				components[c].push(v)
			}
			return components
		}
	}
}

module.exports = {
	CC,
}
