// 有向图
class DiGraph {
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
		this.e++
	}
	Adj(v) { return this.adj[v] }

	reverse() {
		const r = new DiGraph(this.V())
		for (let v = 0; v < this.V(); v++) {
			this.Adj(v).forEach((w) => {
				r.AddEdge(w, v)
			})
		}
		return r
	}

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

function constructDiGraph(edges, v) {
	const g = new DiGraph(v)
	edges.forEach(([v, w]) => {
		g.AddEdge(v, w)
	})
	return g
}

module.exports = {
	DiGraph,
	constructDiGraph,
}
