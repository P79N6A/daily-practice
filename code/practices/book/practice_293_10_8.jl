include("../../statistics/analysis_of_variance/two_factor.jl")
include("../../lib/distribution.jl")

using TwoFactor
using DistributionTable

data = [365 350 343 340 323;
345 368 363 330 333;
358 323 353 343 308;
288 280 298 260 298];
a = 0.05

sst_ = sst(data)
ssr_ = ssr(data)
ssc_ = ssc(data)
sse_ = sse(data)
FR_ = FR(data)
FC_ = FC(data)
n_k = n_of_ssr(data)
n_c = n_of_ssc(data)
n_r = n_of_sse(data)
println("矩阵：", data)
println("显著性水平：", a)
println("SST:", sst_)
println("SSR:", ssr_)
println("SSC:", ssc_)
println("SSE:", sse_)
println("总平方和自由度:", n_of_sst(data))
println("行误差平方和自由度:", n_of_ssr(data))
println("列误差平方和自由度:", n_of_ssc(data))
println("随机误差平方和自由度:", n_of_sse(data))
println("行因素的均方:", msr(data))
println("列因素的均方:", msc(data))
println("随机误差项的均方:", mse(data))
println("行因素对因变量是否有影响的统计量 Fr:", FR_)
println("行因素对因变量是否有影响的统计量 Fc:", FC_)
println("行的Fa(m, n)=", FINV(n_k, n_r, a))
println("列的Fa(m, n)=", FINV(n_c, n_r, a))
println("行列联合效应与因变量之间的关系强度 R: ", R(data))
