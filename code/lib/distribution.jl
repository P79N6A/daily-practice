module DistributionTable
    using Distributions

    export FINV, Ta2

    # F概率分布函数的反函数值
    function FINV(m, n, a)
        d = FDist(m, n)
        quantile(d, (1-a))
    end

    # T概率分布临界值
    #   a/2
    #   v 自由度
    function Ta2(a, v)
        abs(quantile(TDist(v), a/2))
    end
end
