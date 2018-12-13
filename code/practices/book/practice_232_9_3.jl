include("../../statistics/classification_analysis/contingency_table.jl")
using ContingencyTable

data = [6 13 14; 12 16 8; 38 40 11; 21 22 9];
level = 0.05

println("矩阵：", data)
println("X2:", x2(data))
println("自由度：", degree_of_freedom(data))
println("显著性水平：", level)
