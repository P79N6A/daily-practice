class Queue {
	constructor() {
		this.q = []
	}
	enqueue(v) {
		this.q.push(v)
		return this.q
	}
	dequeue() {
		return this.q.shift()
	}
	length() {
		return this.q.length
	}
	toString() {
		return this.q
	}
}

class Stack {
	constructor() {
		this.q = []
	}
	push(v) {
		this.q.push(v)
		return this.q
	}
	pop() {
		return this.q.pop()
	}
	length() {
		return this.q.length
	}
	toString() {
		return this.q
	}
}

class Node {
	constructor(item, next) {
		this._item = item
		this._next = next
	}

	next() { return this._next }
}

class Bag {
	constructor(node) {
		this._first = node
	}

	add(item) {
		const oldFirst = this.first
		this._first = new Node(item, oldFirst)
	}

	first() { return this._first }
}

module.exports = {
	Queue,
	Stack,
	Bag,
}
