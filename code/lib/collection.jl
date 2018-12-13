module Collection
    export average, s, t, total_length, total_sum, total_average

    # 平均值
    function average(arr)
        mean(arr)
    end

    # 标准差
    function s(arr)
        x = average(arr)
        sum = 0
        for i in eachindex(arr)
            sum += (arr[i] - x)^ 2
        end
        sqrt(sum / (length(arr) - 1))
    end

    # t统计量
    #       u0 假设均值
    function t(arr, u0)
        x_ = average(arr)
        s_ = s(arr);
        n  =length(arr)
        (x_ - u0) / (s_ / sqrt(n))
    end

    # 多维数组的总元素个数
    function total_length(arr)
        len = 0
        for i in eachindex(arr)
            len += length(arr[i])
        end
        len
    end

    # 多维数组的和
    function total_sum(arr)
        result = 0
        for i in eachindex(arr)
            for j in eachindex(arr[i])
                result += arr[i][j]
            end
        end
        result
    end

    # 多维数组的平均
    function total_average(arr)
        s = total_sum(arr)
        l = total_length(arr)
        s / l
    end
end
