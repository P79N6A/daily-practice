
module GoodnessOfFitTest
    export x2, df

    # x2卡方统计量可用于测定两个分类变量之间的相关程度
    #   fo_arr 包含 fo（观察值频数)
    #   fe_arr 包含 fe（期望值频数）
    function x2(fo_arr, fe_arr)
        result = 0
        for i in eachindex(fo_arr)
            result += (fo_arr[i] - fe_arr[i])^2 / fe_arr[i]
        end
        result
    end

    # 自由度 = R - 1
    #   R为分类变量类型的个数
    function df(fo_arr)
        length(fo_arr) - 1
    end

    # 二项分布转化为正态分布
    #   np>5, nq>5
    function z_p(p, pi0)
        (p - pi0) / sqrt(pi0 * (1-pi0) / n)
    end
end
