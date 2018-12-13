## z统计量和t统计量用于 均值和比例的检验
## x2卡方统计量用于 方差的检验

## Z值（统计值）检验：    |Z| > |Za/2|，则拒绝
## P值（概率值）检验：    P < Pz，则拒绝

# 样本量n 大的情况下，样本统计量服务正态分布
#   总体标准差已知
#   z统计量可用于与置信水平对应的统计量进行对比
function z_o(x, u0, o, n)
    (x - u0)/o/sqrt(n)
end
#   总体标准差未知
function z_s(x, u0, s, n)
    (x - u0)/s/sqrt(n)
end
# 样本量n 小的情况下
#   总体标准差o 已知，样本统计量服从正态z分布；
#   总体标准差未知，样本统计量服从t分布
function t(x, u0, s, n)
    (x - u0)/s/sqrt(n)
end

# 校验比例，使用统计量z
#   p 为样本比例， pi0为总体比例pi的假设值
function z_p(p, pi0, n)
    (p - pi0) / sqrt(pi0*(1-pi0)/n)
end

# 校验方差，使用卡方分布统计量X2
function x2(s, o, n)
    (n-1) * s^2 / o^2
end
