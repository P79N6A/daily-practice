import { Comparable } from './interfaces'

export class Double implements Comparable<Double> {
  constructor(public value:number) {
  }

  compareTo(that: Double): number {
    if (this.value < that.value) {
      return -1
    } else if (this.value > that.value) {
      return 1
    } else {
      return 0
    }
  }
}
