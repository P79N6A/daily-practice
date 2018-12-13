function d_avg(sample1, sample2)
    diff_sum = 0
    for i in eachindex(sample1)
        diff_sum += sample1[i] - sample2[i]
    end
    diff_sum / length(sample1)
end

function sd(sample1, sample2, d_avg)
    sum = 0
    n = length(sample1)
    for i in eachindex(sample1)
        sum += (sample1[i] - sample2[i] - d_avg)^2
    end
    sqrt(sum / (n - 1))
end

# 样本匹配
function estimate(sample1, sample2, t)
    d = d_avg(sample1, sample2)
    n = length(sample1)
    s = sd(sample1, sample2, d)
    error_estimation = t * s / sqrt(n)
    [d - error_estimation, d + error_estimation]
end

sample1 = [78, 63, 72, 89, 91, 49, 68, 76, 85, 55];
sample2 = [71, 44, 61, 84, 74, 51, 55, 60, 77, 39];
t = 2.2622
estimation = estimate(sample1, sample2, t)
println("样本1：", sample1)
println("样本2：", sample2)
println("T分布临界值：", t)
println("匹配样本：", estimation)
