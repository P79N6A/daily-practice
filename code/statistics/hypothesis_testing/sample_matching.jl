data    = [9.5, 11.5, 8.5, 7.5, 11, 8, 9.5, 7.5, 11, 14.5];
ta      = 1.833;
u0      = 8.5;

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

function s(arr)
    x = average(arr)
    sum = 0
    for i in eachindex(arr)
        sum += (arr[i] - x)^ 2
    end
    sqrt(sum / (length(arr) - 1))
end

function o(arr)
    s(arr) / sqrt(length(arr))
end

x_ = average(data)
o_ = o(data)
println("实际均值:", u0 + ta * o_)
println("假设均值:", x_)
