# 总体均值的区间估计
#   [x - t * s / sqrt(n), x + t * s / sqrt(n)]

sample = [1510, 1450, 1480, 1460, 1520, 1480, 1490, 1460,
1480, 1510, 1530, 1470, 1500, 1520, 1510, 1470];

t = 2.131;

function sum(arr)
    result = 0
    for i in eachindex(arr)
        result += arr[i]
    end
    result
end

function average(arr)
    sum(arr) / length(arr)
end

function standard_deviation(arr)
    x = average(arr)
    sum = 0
    for i in eachindex(arr)
        sum += (arr[i] - x)^ 2
    end
    sqrt(sum / (length(arr) - 1))
end

function confidence_interval(arr, t)
    n = length(sample)
    x = average(sample)
    s = standard_deviation(sample)
    [x - t * s / sqrt(n), x + t * s / sqrt(n)]
end

println("样本：", sample)
println("X平均：", average(sample))
println("t分布临界值：", t)
println("置信区间：", confidence_interval(sample, t))
