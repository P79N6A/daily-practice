/*
    DFA:
                  OPERATOR, PARENTHESES, NUMBER
                  ------------------------------>       ----\
        -> ((s0))                                 ((s1))    | NUMBER
                  <------------------------------       <---/
                  OPERATOR, PARENTHESES
*/
function tokenize(expression) {
	const result = []
	let token = ''
	let status
	expression.split('').forEach((char, index) => {
		if(!valid(char)) {
			throw new Error('invalid char:', char)
		}
		if(isNumber(char)) {
			if(status !== STATUS.NUM && token.length > 0) {
				appendLexeme(result, status, token)
				token = char
			} else {
				token += char
			}
			status = STATUS.NUM
		} else if(isOperator(char)) {
			if(status === STATUS.OP) {
				throw new Error(`can't appear continuous operator.`)
			}
			if(token.length > 0) {
				appendLexeme(result, status, token)
			}
			token = char
			status = STATUS.OP
		} else if(isParentheses(char)) {
			if(status != STATUS.PAREN && token.length > 0) {
				appendLexeme(result, status, token)
			}
			token = char
			status = STATUS.PAREN
		}
		if(index === expression.length - 1) {
			appendLexeme(result, status, token)
		}
	})
	return result
}

const STATUS = {
	NUM: 'NUMBER',
	OP: 'OPERATOR',
	PAREN: 'PARENTHESES',
}

function appendLexeme(results, type, value) {
	results.push({ type, value })
}

function isNumber(char) {
	return /^\d+$/.test(char)
}

function isParentheses(char) {
	return ['(', ')'].indexOf(char) > -1
}

function isOperator(char) {
	return ['+', '-', '*', '/'].indexOf(char) > -1
}

function valid(char) {
	return isNumber(char) || isParentheses(char) || isOperator(char) || char === ' '
}

module.exports = tokenize
