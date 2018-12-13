import { Comparable } from './interfaces'

// 索引优先队列
export class IndexMinPQ<T extends Comparable<T>> {
  private N: number
  private pq: Array<number>
  private qp: Array<number>
  private keys: Array<T>

  private less(i: number, j: number): boolean {
    return this.pq[i] < this.pq[j]
  }
  private exch(i: number, j: number) {
    const t = this.pq[i]
    this.pq[i] = this.pq[j]
    this.pq[j] = t
  }
  private swim(k: number) {
    while (k > 1 && this.less(k/2, k)) {
      this.exch(k/2, k)
      k = k/2
    }
  }
  private sink(k: number) {
    while (2*k <= this.N) {
      let j = 2*k
      if (j < this.N && this.less(j, j+1)) {
        j++
      }
      if (!this.less(k, j)) {
        break
      }
      this.exch(k, j)
      k = j
    }
  }

  constructor(maxN: number) {
    this.keys = new Array<T>(maxN + 1)
    this.pq = new Array<number>(maxN + 1)
    this.qp = new Array<number>(maxN + 1)
    for (let i = 0; i <= maxN; i++) {
      this.qp[i] = -1
    }
  }

  isEmpty() { return this.N === 0 }

  contains(k: number) { return this.qp[k] !== - 1 }

  insert(k: number, key: T) {
    this.N++
    this.qp[k] = this.N
    this.pq[this.N] = k
    this.keys[k] = key
    this.swim(this.N)
  }

  min(): T {
    return this.keys[this.pq[1]]
  }

  delMin(): number {
    const indexOfMin = this.pq[1]
    this.exch(1, this.N--)
    this.sink(1)
    this.keys[this.pq[this.N+1]] = null
    this.qp[this.pq[this.N+1]] = -1
    return indexOfMin
  }
}
