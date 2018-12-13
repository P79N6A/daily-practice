// 有向图的DFS
class DirectedDFS {
	constructor(g) {
		this.graph = g
    this.marked = new Array(this.graph.V())
    this.hasCycle = false
	}

  dfs(v, u) {
    this.marked[v] = true
    this.graph.Adj(v).forEach((adj) => {
      if (!this.marked[adj]) {
        this.dfs(adj, v)
      } else if (adj !== u) {
        this.hasCycle = true
      }
    })
  }

  cycle(s) {
    this.hasCycle = false
    this.marked = new Array(this.graph.V())

    for (let s = 0; s < this.graph.V(); s++) {
      if (!this.marked[s]) {
        this.dfs(s, s)
      }
    }

    return this.hasCycle
  }
}

module.exports = {
	DirectedDFS,
}
