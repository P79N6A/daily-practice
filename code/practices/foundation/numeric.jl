println("typeof(π): ", typeof(π))
println(π)

println("字长：", Sys.WORD_SIZE)

for T = [Int8, Int16, Int32, Int64, Int128, UInt8, UInt16, UInt32, UInt64, UInt128]
    println("$(lpad(T, 7)): [$(typemin(T)), $(typemax(T))]")
end

x = Int128(2^32)
println(x*x)

x = 7
y = 5
z = x / y
println("$(x)/$(y)=", z, " ", typeof(z))

x = 0x3166
x <<= 1
println(x)

a = rand(1:100, 22)
push!(a, 666)
println("rand array:", a)

a = Array(Int64, 15);
a[1] = 0;
a[2] = 1;
[a[i] = a[i-1] + a[i-2] for i = 3:length(a)]
println(a)


A = [1 2 3; 4 5 6]
B = reshape(A, 3, 2)
C = transpose(B)
println("A:", A)
println("B:", B)
println("C:", C)

println("矩阵")
println("A + C = ", A + C)
println("A * B = ", A * B)
println("A / C = ", A / C)

try
    sqrt(-1)
catch e
    println("Error:", e)
end
