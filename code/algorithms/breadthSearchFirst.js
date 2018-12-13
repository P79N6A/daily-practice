class BreadthFirstSearch {
	constructor(g) {
		this.graph = g
	}

	bfs(queue, marked, edgeTo) {
		if (!queue.length) {
			return
		}
		const oldQueue = [].concat(queue)
		queue = []
		oldQueue.forEach((v) => {
			if (!marked[v]) {
				marked[v] = true
				this.graph.Adj(v).forEach((adj) => {
					if (!marked[adj]) {
						edgeTo[adj] = v
						queue.push(adj)
					}
				})
			}
		})
		this.bfs(queue, marked, edgeTo)
	}

	pathTo(s, v) { // 最短通路
		const queue = [s]
		const marked = new Array(this.graph.V())
		const edgeTo = new Array(this.graph.V())
		this.bfs(queue, marked, edgeTo)

		if (marked[v]) {
			const path = []
			for (let x = v; x != s; x = edgeTo[x]) {
				path.push(x)
			}
			return path.length ? path.concat([s]).reverse() : null
		}
	}
}

module.exports = {
	BreadthFirstSearch,
}
