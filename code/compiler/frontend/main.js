const tokenize = require('./tokenizer')
const psr = require('./parser')
const trans = require('./transformer')
const println = console.log
const inspect = (obj) => console.dir(obj, { depth: null })

function test(expression) {
	println('原始表达式: ', expression)

	const token = tokenize(expression)
	println('词法序列：')
	inspect(token)
	const parser = new psr.Parser(token)
	const tree = parser.parse()
	println('具体语法树：')
	inspect(tree)

	println('转换后:')
	println(trans.transfom(tree))
}

test('(10+20)*5')
