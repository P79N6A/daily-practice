include("./contingency_table.jl")
using ContingencyTable

x2 = 19.82
n = 500
row = col = 3

println("n:", n)
println("卡方统计量:", x2)
println("φ系数:", round(phi(x2, n), 3))
println("c系数:", round(C(x2, n),3))
println("V系数:", round(V(x2, n, row, col),3))
