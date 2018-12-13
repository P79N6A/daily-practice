import { IndexMinPQ } from './bases/indexMinPQ'
import { Double } from './bases/double'
import { Edge } from './edge'
import { EdgeWeightGraph } from './edgeWeightGraph'

// 最小生成树 Prim 算法即时版
export class PrimMST {
  private edgeTo: Array<Edge>
  private distTo: Array<number>
  private marked: Array<boolean>
  private pq: IndexMinPQ<Double>

  constructor() {

  }
}
