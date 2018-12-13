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

function transfom(root) {
	return extractExpr(root.children[0])
}

function extractExpr(Expr) {
	let expr = Expr
	let result = extractTerm(expr.children[0])
	while (expr.children[1].children) {
		const operator = expr.children[1].operator
		const secondParam = extractTerm(expr.children[1].children[0])
		result = `(${operator} ${result} ${secondParam})`
		expr = expr.children[1]
	}
	return result
}

function extractTerm(Term) {
	let term = Term
	let result = extractFactor(term.children[0])
	while (term.children[1].children) {
		const operator = term.children[1].operator
		const secondParam = extractFactor(term.children[1].children[0])
		result = `(${operator} ${result} ${secondParam})`
		term = term.children[1]
	}
	return result
}

function extractFactor(node) {
	if (node.number) {
		return node.number
	} else if (node.children) {
		return extractExpr(node.children[0])
	}
}

module.exports.transfom = transfom
