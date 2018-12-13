# 总体比例的区间估计
#   [p - z * sqrt(p * (1 - p) / n), p + z * sqrt(p * (1 - p) / n)]
function proportion_estimate(p, n, z)
    error_estimation = z * sqrt(p * (1 - p) / n)
    [p - error_estimation, p + error_estimation]
end

function wrap(interval)
    [string(round(interval[1] * 100, 2), "%"), string(round(interval[2] * 100, 2), "%")]
end

p = 0.65
n = 100
z = 1.96
interval =  proportion_estimate(p, n, z)
println("0.95置信水平下，该城市下岗职工中女性比例为：", wrap(interval))
