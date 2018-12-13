module Matrix
    export value_by_xy, mean_all, mean_col, mean_row

    # 获取指定坐标点的值
    function value_by_xy(dims, x, y)
        row, col = size(dims)
        pos = (y - 1) * row + x
        dims[pos]
    end

    # 总平均值
    function mean_all(dims)
        mean(dims)
    end

    # 某列的平均值
    function mean_col(dims, j)
        row, _ = size(dims)
        total = 0
        for i in Array(1:row)
            total += value_by_xy(dims, i, j)
        end
        total / row
    end

    # 某行的平均值
    function mean_row(dims, i)
        _, col = size(dims)
        total = 0
        for j in Array(1:col)
            total += value_by_xy(dims, i, j)
        end
        total / col
    end
end
