

# 两个总体方差比的区间估计
function estimate(x1, s1, x2, s2, f)
    f2 = 1 / f
    s = s1^2 / s2^2
    [s/f, s/f2]
end
