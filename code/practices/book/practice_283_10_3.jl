include("../../lib/collection.jl")
include("../../statistics/analysis_of_variance/single_factor.jl")
include("../../lib/distribution.jl")

using Collection
using SingleFactor
using DistributionTable

data = [[57,66,49,40,34,53,44],
[68, 39, 29, 45, 56, 51],
[31, 49, 21, 34, 40],
[44, 51, 65, 77, 58]];

n = total_length(data)
sum_ = total_sum(data)
a = 0.05
k_1 = n_of_ssa(data)
n_k = n_of_sse(data)
i = 1
j = 3
println("n=", n)
println("sum=", sum_)
println("x_=", total_average(data))
println("总平方和  SST=", sst(data))
println("自由度 n=", n_of_sst(data))
println("组间平方和  SSA=", ssa(data))
println("自由度 k-1=", k_1)
println("组间均方 MSA=", msa(data))
println("组内平方和  SSE=", sse(data))
println("自由度 n-k=", n_k)
println("组内均方 MSE=", mse(data))
println("F=MSA/MSE= ", F(data))
println("显著性水平 a= ", a)
println("Fa(m,n)= ", FINV(k_1, n_k, a))
println("关系强度 R= ", R(data))
println("样本$(i)与样本$(j)的 LSD= ", LSD(data, i, j, a))
println("样本$(i)与样本$(j)的 均差= ", MD(data, i, j))
