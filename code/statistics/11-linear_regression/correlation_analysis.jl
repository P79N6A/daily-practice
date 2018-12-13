include("../../lib/collection.jl")

# 相关性分析
#   数据结构arr是二维数组:
#   [[x11, x12], [x21, x22],...]
module CorrelationAnalysis
    using Collection
    using Distributions

    export R, t

    # 相关系数R
    #   |r| >= 0.8，高度相关；
    #   0.5 =< |r| <0.8 中度相关；
    #   0.3 =< |r| < 0.5 低度相关；
    #   |r| < 0.3 相关极弱
    #       参数：第y行 和 第y行
    function R(arr, x, y)
        n = length(arr[x])
        sum_xy  = 0
        sum_x   = 0
        sum_y   = 0
        sum_x2  = 0
        sum_y2  = 0
        for j in Array(1:n)
            x_j     = arr[x][j]
            y_j     = arr[y][j]
            sum_xy  += x_j * y_j
            sum_x   += x_j
            sum_y   += y_j
            sum_x2  += x_j^2
            sum_y2  += y_j^2
        end
        (n*sum_xy-(sum_x*sum_y))/(sqrt(n*sum_x2 - sum_x^2) * sqrt(n*sum_y2-sum_y^2))
    end

    # 显著性检验的t统计量
    function t(arr, x, y)
        r = abs(R(arr, x, y))
        n = length(arr[x])
        r * sqrt((n-2) / (1-r^2))
    end
end
