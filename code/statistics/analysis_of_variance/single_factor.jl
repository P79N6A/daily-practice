include("../../lib/collection.jl")

# 单因素方差分析
#   数据结构arr是二维数组:
#   [[x11, x12], [x21, x22],...]
module SingleFactor
    using Collection
    using Distributions

    export sst, ssa, sse, n_of_sst, n_of_ssa, n_of_sse, msa, mse,
            F, R, LSD, MD;

    # 总平方和
    function sst(arr)
        result = 0
        x_ = total_average(arr)
        for i in eachindex(arr)
            for j in eachindex(arr[i])
                result += (arr[i][j] - x_)^2
            end
        end
        result
    end

    # 组间平方和
    function ssa(arr)
        result = 0
        x_ = total_average(arr)
        for i in eachindex(arr)
            x_i = average(arr[i])
            n = length(arr[i])
            result += n*(x_i-x_)^2
        end
        result
    end

    # 组内平方和
    function sse(arr)
        result = 0
        for i in eachindex(arr)
            x_i = average(arr[i])
            for j in eachindex(arr[i])
                result += (arr[i][j] - x_i)^2
            end
        end
        result
    end

    # 总平方和的自由度
    function n_of_sst(arr)
        total_length(arr) - 1
    end

    # 组间平方和的自由度
    function n_of_ssa(arr)
        length(arr) - 1
    end

    # 组内平方和的自由度
    function n_of_sse(arr)
        n = total_length(arr)
        k = length(arr)
        n - k
    end

    # 组间均方
    function msa(arr)
        ssa(arr) / n_of_ssa(arr)
    end

    # 组内均方
    function mse(arr)
        sse(arr) / n_of_sse(arr)
    end

    # F=MSA/MSE
    # 当H0为真时，二者的比值服从分子自由度为k-1,分母自由度为n-k的F分布
    function F(arr)
        msa(arr)/mse(arr)
    end

    # 关系强度
    function R(arr)
        sqrt(ssa(arr) / sst(arr))
    end

    # 最小显著性差异方法
    #   a 显著性水平
    function LSD(arr, i, j, a)
        ni = 1/length(arr[i])
        nj = 1/length(arr[j])
        ta2 = abs(quantile(TDist(n_of_sse(arr)), a/2))
        ta2 * sqrt(mse(arr)*(ni+nj))
    end

    # 样本i与样本j的均差绝对值
    function MD(arr, i, j)
        abs(mean(arr[i]) - mean(arr[j]))
    end
end
