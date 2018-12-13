function level2z(level)
    if level == 0.90
        1.64
    elseif level == 0.95
        1.96
    elseif level == 0.99
        2.58
    else
        1.96
    end
end

# 两个总体均值之差的区间估计（大样本）
function big_estimate(n1, x1, s1, n2, x2, s2, level)
    z = level2z(level)
    error_estimation = z * sqrt(s1^2/n1 + s2^2/n2)
    x = x1 - x2
    [x - error_estimation, x + error_estimation]
end

# 两个总体均值之差的区间估计（小样本）
#   两总体的方差o1和o2未知且相等时
function small_estimate_eq(n1, x1, s1, n2, x2, s2, t)
    sp = ((n1-1)*s1^2 + (n2-1)*s2^2)/(n1+n2-2)
    error_estimation = t * (n1+n2-2) * sqrt(sp * (1/n1 + 1/n2))
    x = x1 - x2
    [x - error_estimation, x + error_estimation]
end

# 两个总体均值之差的区间估计（小样本）
#   两总体的方差o1和o2未知且不相等时
function small_estimate_neq(n1, x1, s1, n2, x2, s2, t)
    v = (s1^2/n1 + s2^2/n2)^2 / ((s1^2/n1)^2/(n1-1) + (s2^2/n2)^2/(n2-1))
    error_estimation = t * v * sqrt(s1^2/n1 + s2^2/n2)
    x = x1 - x2
    [x - error_estimation, x + error_estimation]
end
