# 总体均值的区间估计
#   [x - z * s / sqrt(n), x + z * s / sqrt(n)]

sample = [23, 35, 39, 27, 36, 44,
36, 42, 46, 43, 31, 33,
42, 53, 45, 54, 47, 24,
34, 28, 39, 36, 44, 40,
39, 49, 38, 34, 48, 50,
34, 39, 45, 48, 45, 32];

level = 0.90;

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
    x = average(sample)
    sum = 0
    for i in eachindex(arr)
        sum += (arr[i] - x)^ 2
    end
    sqrt(sum / (length(arr) - 1))
end

function level2z(level)
    if level == 0.90
        1.64
    elseif level == 0.95
        1.96
    elseif level == 0.99
        2.58
    else
        1.96
    end
end

function confidence_interval(arr, level)
    n = length(sample)
    x = average(sample)
    s = standard_deviation(sample)
    z = level2z(level)
    [x - z * s / sqrt(n), x + z * s / sqrt(n)]
end

println("样本：", sample)
println("X平均：", average(sample))
println("置信水平：", level)
println("置信区间：", confidence_interval(sample, level))
