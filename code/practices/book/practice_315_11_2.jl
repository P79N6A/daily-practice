include("../../statistics/linear_regression/correlation_analysis.jl")
include("../../lib/distribution.jl")

using CorrelationAnalysis
using DistributionTable

data = [[0.9, 1.1, 4.8, 3.2, 7.8, 2.7, 1.6, 12.5,1.0,2.6, 0.3, 4.0, 0.8,
3.5, 10.2, 3.0, 0.2, 0.4, 1.0, 6.8, 11.6, 1.6, 1.2, 7.2, 3.2],
[67.3, 111.3, 173.0, 80.8, 199.7, 16.2, 107.4, 185.4, 96.1, 72.8, 64.2, 132.2, 58.6,
174.6, 263.5, 79.3, 14.8, 73.5, 24.7, 139.4, 368.2, 95.7, 109.6, 196.2, 102.2]];

x = 1
y = 2
a = 0.05
v = length(data[x]) - 2
println("样本$(x)：", data[x])
println("样本$(y)：", data[y])
println("样本$(x) 与样本$(y) 的相关系数 R=", R(data, x, y))
println("显著性检验的t统计量 = ", t(data, x, y))
println("t($(a/2))(n-2) = ", Ta2(a, v))
