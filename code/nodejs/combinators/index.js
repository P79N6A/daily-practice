const println = console.log

function z(fn) {
	const recurFunc = recur => recur(recur)
	return recurFunc(r => x => fn(r(r))(x))
}

const factorial = z(fn => n => 2 >= n ? n : n * fn(n - 1))

println(factorial(2))
println(factorial(3))
println(factorial(4))
println(factorial(5))
