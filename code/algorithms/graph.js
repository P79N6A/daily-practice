class Graph {
	constructor(v) {
		this.v = v
		this.e = 0
		this.adj = new Array(v)
		for (let i = 0; i < v; i++) {
			this.adj[i] = new Array()
		}
	}

	V() { return this.v }
	E() { return this.e }
	AddEdge(v, w) {
		this.adj[v].push(w)
		this.adj[w].push(v)
		this.e++
	}
	Adj(v) { return this.adj[v] }

	toString() {
		const buf = []
		for (let i = 0; i < this.adj.length; i++) {
			buf.push(i + ':\t')
			buf.push(this.adj[i].join(','))
			buf.push('\n')
		}
		return buf.join('')
	}
}

function constructGraph(edges, v) {
	const g = new Graph(v)
	edges.forEach(([v, w]) => {
		g.AddEdge(v, w)
	})
	return g
}

module.exports = {
	Graph,
	constructGraph,
}
