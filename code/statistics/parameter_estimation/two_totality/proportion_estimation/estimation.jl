
# 两个总体比例之差的区间估计
function estimate(p1, n1, p2, n2, z)
    p = p1 - p2
    error_estimation = z * sqrt(p1*(1-p1)/n1 + p2*(1-p2)/n2)
    [p - error_estimation, p + error_estimation]
end
