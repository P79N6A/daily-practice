import { Iterator, Iterable } from './interfaces'

export class Bag<T> implements Iterable<T> {
  private first: Node<T>

	add(item: T) {
		const oldFirst = this.first
		this.first = new Node<T>(item, oldFirst)
	}

  iterator(): Iterator<T> {
    return new ListIterator<T>()
  }
}

class Node<T> {
  item: T
  next: Node<T>

  constructor(item: T, next: Node<T>|null) {
    this.item = item
    this.next = next
  }
}

class ListIterator<T> implements Iterator<T> {
  private current: Node<T>
  hasNext(): boolean {
    return !!this.current
  }
  next(): T {
    const item = this.current.item
    this.current = this.current.next
    return item
  }
  remove() { }
}
