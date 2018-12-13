include("../../lib/collection.jl")

# 一元线性回归
#   数据结构arr是二维数组:
#   [[x11, x12], [x21, x22],...]
module UnaryRegression
    using Collection
    using Distributions

    # 方程的b1项，直线的斜率
    function b1(array, x, y)
        n       = length(array)
        size    = length(array[1])
        sum_x   = 0
        sum_x2  = 0
        sum_y   = 0
        sum_xy  = 0
        for i in Array(1:size)
            sum_x   += array[x][i]
            sum_x2  += array[x][i]^2
            sum_y   += array[y][i]
            sum_xy  += array[x][i] * array[y][i]
        end
        (n * sum_xy - sum_x * sum_y) / (n * sum_x2 - sum_x^2)
    end

    # 方程的b0项，y轴上的截矩
    function b0(array, x, y)
        x_      = mean(array[x])
        y_      = mean(array[y])
        y_ - b1(array, x, y) * x_
    end
end
