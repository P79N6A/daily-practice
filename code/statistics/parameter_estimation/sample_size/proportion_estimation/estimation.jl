

# 总体比例的样本量估计
#   p 总概率
#   E 误差
#   z 置信水平
function n(p, E, z)
    z^2 * p * (1-p) / E^2
end
