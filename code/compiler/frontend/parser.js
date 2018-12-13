/*
CFG:
    Expr      ->  Term ExprTail
    ExprTail  ->  + Term ExprTail
              |   - Term ExprTail
              |   null
    Term      ->  Factor TermTail
	TermTail  ->  * Factor TermTail
              |   / Factor TermTail
              |   null
    Factor    ->  (Expr)
              |   number
*/

class Node {
	constructor(type, child) {
		this.type = type
		if (child) {
			this.append(child)
		}
	}
	append(child) {
		this.children = this.children || []
		this.children.push(child)
	}
}

class Parser {
	constructor(token) {
		this.token = token
		this.index = 0
		this.root = new Node('ROOT')
		this.current = this.root
		this.fetchNext()
	}

	parse() {
		if (this.Expr() && !this.word()) {
			return this.root
		} else {
			return false
		}
	}

	/*
    Term      ->  Factor TermTail
	*/
	Expr() {
		const node = new Node('Expr')
		this.current.append(node)
		this.current = node
		if (!this.Term()) {
			return false
		} else {
			this.current = node
			return this.ExprTail()
		}
	}
	/*
	Term      ->  Factor TermTail
	*/
	Term() {
		const node = new Node('Term')
		this.current.append(node)
		this.current = node
		if (!this.Factor()) {
			return false
		} else {
			this.current = node
			return this.TermTail()
		}
	}
	/*
	ExprTail  ->  + Term ExprTail
              |   - Term ExprTail
              |   null
	*/
	ExprTail() {
		const node = new Node('ExprTail')
		this.current.append(node)
		this.current = node
		if (this.word() === '+' || this.word() === '-') {
			this.current.operator = this.word()
			this.fetchNext()
			if (!this.Term()) {
				return false
			} else {
				this.current = node
				return this.ExprTail()
			}
		}
		return true
	}
	/*
	TermTail  ->  * Factor TermTail
              |   / Factor TermTail
              |   null
	*/
	TermTail() {
		const node = new Node('TermTail')
		this.current.append(node)
		this.current = node
		if (this.word() === '*' || this.word() === '/') {
			this.current.operator = this.word()
			this.fetchNext()
			if (!this.Factor()) {
				return false
			} else {
				this.current = node
				return this.TermTail()
			}
		}
		return true
	}
	/*
	Factor    ->  (Expr)
              |   number
	*/
	Factor() {
		const node = new Node('Factor')
		this.current.append(node)
		this.current = node
		if (this.word() === '(') {
			this.fetchNext()
			if (!this.Expr()) {
				return false
			} else if (this.word() !== ')') {
				return false
			}
			this.fetchNext()
			return true
		} else if (this.isNumber()) {
			this.current.number = this.word()
			this.fetchNext()
			return true
		}
		return false
	}

	word() {
		return this.lexeme ? this.lexeme.value : undefined
	}
	fetchNext() {
		this.lexeme = this.token[this.index++]
	}
	isNumber() {
		return /^\d+$/.test(this.word())
	}
}

module.exports.Parser = Parser
