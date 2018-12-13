class DepthFirstSearch {
	constructor(g) {
		this.graph = g
		this.hasCycle = false
	}

	dfs(v, marked) {
		marked[v] = true
		this.graph.Adj(v).forEach((adj) => {
			if (!marked[adj]) {
				this.dfs(adj, marked)
			}
		})
	}

	connected(s) { // 是否为连通图
		const marked = new Array(this.graph.V())
		this.dfs(s, marked)

		for (let v = 0; v < marked.length; v++) {
			if (!marked[v]) {
				return false
			}
		}
		return true
	}

	dfs2(v, w, traces, found) {
		traces.push(v)
		if (v === w) {
			found.push(w)
		}
		if (found.length) {
			return
		}
		for (let adjs = this.graph.Adj(v), i = 0; i < adjs.length; i++) {
			const adj = adjs[i]
			if (adj === w) {
				if (traces.indexOf(adj) === -1) {
					traces.push(adj)
					found.push(adj)
				}
				return
			}
			if (!found.length && traces.indexOf(adj) === -1) {
				this.dfs2(adj, w, traces, found)
			}
		}
	}

	pathway(v, w) { // 是否有通路
		const traces = []
		const found = []
		this.dfs2(v, w, traces, found)

		if (found.length) {
			return traces
		}
	}

	dfs3(v, u, marked) {
		marked[v] = true
		this.graph.Adj(v).forEach((adj) => {
			if (!marked[adj]) {
				this.dfs3(adj, u, marked)
			} else if (adj !== u) {
				this.hasCycle = true
			}
		})
	}

	cycle() {
		this.hasCycle = false
		const marked = new Array(this.graph.V())

		for (let s = 0; s < this.graph.V(); s++) {
			if (!marked[s]) {
				this.dfs3(s, s, marked)
			}
		}

		return this.hasCycle
	}
}

module.exports = {
	DepthFirstSearch,
}
