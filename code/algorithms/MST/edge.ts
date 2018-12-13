import { Comparable } from './bases/interfaces'

export class Edge implements Comparable<Edge> {
	constructor(public v: number, public w:number, public weight:number) {
	}

	either() { return this.v }

	other(vertex) {
    if (vertex  === this.v) {
      return this.w
    } else if (vertex === this.w) {
      return this.v
    } else {
      throw 'Inconsistent edge'
    }
	}

	compareTo(that: Edge): number {
    if (this.weight < that.weight) {
      return -1
    } else if (this.weight > that.weight) {
      return 1
    } else {
      return 0
    }
	}

	toString() {
    return `${this.v} ${this.w} ${this.weight}`
	}
}
