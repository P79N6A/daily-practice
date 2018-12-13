const dir = console.dir
const parser = require('@babel/parser')
const fs = require('fs')
const file = './sample.js'

let ast = null
const code = fs.readFileSync(file, "utf8")
ast = parser.parse(code)
for (node of ast.program.body) {
    dir(node)
}


