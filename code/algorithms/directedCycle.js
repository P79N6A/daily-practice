const structs = require('./structs')

// 寻找有向环
class DirectedCycle {
	constructor(g) {
		this.graph = g
		this.cycle = new structs.Stack()
		this.onStack = new Array(this.graph.V())
		this.edgeTo = new Array(this.graph.V())
		this.marked = new Array(this.graph.V())
		for (let v = 0; v < this.graph.V(); v++) {
			if (!this.marked[v]) {
				this.dfs(v)
			}
		}
	}

	dfs(v) {
		this.onStack[v] = true
		this.marked[v] = true
		this.graph.Adj(v).forEach((adj) => {
			if (this.hasCycle()) {
				return
			} else if (!this.marked[adj]) {
				this.edgeTo[adj] = v
				this.dfs(adj)
			} else if (this.onStack[adj]) {
				this.cycle = new structs.Stack()
				for (let x = v; x != adj; x = this.edgeTo[x]) {
					this.cycle.push(x)
				}
				this.cycle.push(adj)
				this.cycle.push(v)
			}
		})
		this.onStack[v] = false
	}

	hasCycle() {
		return !!this.cycle.length()
	}

	cycle() {
		return this.cycle
	}
}

module.exports = {
	DirectedCycle,
}
