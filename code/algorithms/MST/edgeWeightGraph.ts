import { Bag } from './bases/bag'
import { Edge } from './edge'

// 加权无向图
export class EdgeWeightGraph {
	private v: number
	private e: number
	private adj: Bag<Edge>[]

	constructor(v) {
		this.v = v
		this.e = 0
		this.adj = new Array<Bag<Edge>>(v)
		for (let i = 0; i < v; i++) {
			this.adj[i] = new Bag<Edge>()
		}
	}

	V() { return this.v }
	E() { return this.e }
	addEdge(e) {
		const v = e.either()
		const w = e.other(v)
		this.adj[v].add(e)
		this.adj[w].add(e)
	}
	Adj(v) { return this.adj[v] }
}
