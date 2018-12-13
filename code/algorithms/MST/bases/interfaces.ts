export interface Comparable<T> {
  compareTo(that: T): number
}

export interface Iterator<T> {
  hasNext(): boolean
  next(): T
  remove()
}

export interface Iterable<T> {
  iterator(): Iterator<T>
}
