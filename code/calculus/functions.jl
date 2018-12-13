module Functions
    export solution, m

    # 斜率
    #   直线方程的点斜式
    function m(x1, y1, x2, y2)
        (y2-y1)/(x2-x1)
    end

    # 二次函数的解
    #   返回一个 Tuple{Float64,Float64}
    function solution_of_x2(a, b, c)
        p = sqrt(b^2-4*a*c)
        (-b+p)/(2*a), (-b-p)/(2*a)
    end
end
