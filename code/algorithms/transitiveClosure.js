const directedDFS = require('./directedDFS')

// 顶点可达性
class TransitiveClosure {
	constructor(g) {
		this.graph = g
		this.all = new Array(this.graph.V())
		for (let v = 0; v < all.length; i++) {
			all[v] = new directedDFS.DirectedDFS(g)
      all[v].dfs(v)
		}
	}

	reachable(v, w) {
		return this.all.marked[w]
	}
}

module.exports = {
	TransitiveClosure,
}
