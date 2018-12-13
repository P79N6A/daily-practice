include("../../lib/collection.jl")
include("../../lib/matrix.jl")

# 双因素方差分析
#   数据结构 dims 是矩阵:
#   [x11 x12 x13]
#   [x21 x22 x23]
module TwoFactor
    using Distributions
    using Matrix

    export sst, ssr, ssc, sse, n_of_sst, n_of_ssr, n_of_ssc,
            n_of_sse, msr, msc, mse, FR, FC, R;

    # 总误差平方和
    function sst(dims)
        result = 0
        x_ = mean_all(dims)
        row, col = size(dims)
        for i in Array(1:row)
            for j in Array(1:col)
                x_ij = value_by_xy(dims, i, j)
                result += (x_ij - x_)^2
            end
        end
        result
    end

    # 行误差平方和
    function ssr(dims)
        result = 0
        x_ = mean_all(dims)
        row, col = size(dims)
        for i in Array(1:row)
            for j in Array(1:col)
                x_i = mean_row(dims, i)
                result += (x_i - x_)^2
            end
        end
        result
    end

    # 列误差平方和
    function ssc(dims)
        result = 0
        x_ = mean_all(dims)
        row, col = size(dims)
        for i in Array(1:row)
            for j in Array(1:col)
                x_j = mean_col(dims, j)
                result += (x_j - x_)^2
            end
        end
        result
    end

    # 随机误差平方和
    function sse(dims)
        result = 0
        x_ = mean_all(dims)
        row, col = size(dims)
        for i in Array(1:row)
            for j in Array(1:col)
                x_j = mean_col(dims, j)
                x_i = mean_row(dims, i)
                x_ij = value_by_xy(dims, i, j)
                result += (x_ij - x_i - x_j + x_)^2
            end
        end
        result
    end

    # 总平方和自由度
    function n_of_sst(dims)
        length(dims) - 1
    end

    # 行误差平方和自由度
    function n_of_ssr(dims)
        row, _ = size(dims)
        row - 1
    end

    # 列误差平方和自由度
    function n_of_ssc(dims)
        _, col = size(dims)
        col - 1
    end

    # 随机误差平方和自由度
    function n_of_sse(dims)
        row, col = size(dims)
        (row-1) * (col-1)
    end

    # 行因素的均方
    function msr(dims)
        row, _ = size(dims)
        ssr(dims) / (row-1)
    end

    # 列因素的均方
    function msc(dims)
        _, col = size(dims)
        ssc(dims) / (col-1)
    end

    # 随机误差项的均方
    function mse(dims)
        row, col = size(dims)
        sse(dims) / ((row-1)*(col-1))
    end

    # 行因素对因变量是否有影响的统计量
    function FR(dims)
        msr(dims) / mse(dims)
    end

    # 列因素对因变量是否有影响的统计量
    function FC(dims)
        msc(dims) / mse(dims)
    end

    # 行列联合效应与因变量之间的关系强度
    function R(dims)
        sqrt((ssr(dims) + ssc(dims))/sst(dims))
    end
end
