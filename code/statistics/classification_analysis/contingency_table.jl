include("../../lib/collection.jl")
include("../../lib/matrix.jl")

module ContingencyTable
    using Collection
    using Matrix

    export x2, fe_by_xy, degree_of_freedom, correlation_coefficient, phi, C, V;

    # 计算卡方统计量
    #   x2 = ∑(fo-fe)^2/fe
    #   dims 矩阵
    function x2(dims)
        n = sum(dims)
        row, col = size(dims)
        result = 0

        for i in Array(1:row)
            for j in Array(1:col)
                fo = value_by_xy(dims, i, j)
                fe = fe_by_xy(dims, i, j, row, col, n)
                result += (fo - fe)^2/fe
            end
        end

        result
    end


    # 计算指定坐标的频数期望
    #   row 总行数
    #   col 总列数
    #   n 总合计
    function fe_by_xy(dims, x, y, row, col, n)
        rt = ct = 0

        for j = Array(1:col)
            rt += value_by_xy(dims, x, j)
        end

        for i = Array(1:row)
            ct += value_by_xy(dims, i, y)
        end
        fe_by_rt_ct_n(rt, ct, n)
    end

    # 频数期望：
    #   fe = RT/n * CT/n * n
    #   fe = RT*CT/n
    #       RT: 给定单元所有行的合计
    #       CT: 给定单元所有列的合计
    #       n:  观察值总的合计
    function fe_by_rt_ct_n(rt, ct, n)
        rt * ct / n
    end

    # 自由度
    #   r 行　c　列
    function degree_of_freedom(dims)
        r, c = size(dims)
        (r-1) * (c - 1)
    end

    # φ相关系数
    #   φ 读做phi
    #   适胜于 2X2 列联表
    #   其中, x2（卡方）= ∑(fo-fe)^2/fe
    function phi(x2, n)
        sqrt(x2/n)
    end

    # 列联系数, c系数
    function C(x2, n)
        sqrt(x2/(x2+n))
    end

    # V系数
    function V(x2, n, r, c)
        sqrt(x2 / (n*min(r-1, c-1)))
    end
end
