include("../../lib/collection.jl")
using Collection

data    = [159, 280, 101, 212, 224, 379, 179, 264,
        222, 362, 168, 250, 149, 260, 485, 170];
ta2     = 2.131;
u0      = 225;
n       = length(data)
x_      = average(data)
s_      = s(data);
t_      = t(data, u0)

println("n=", n, ", x平均=", x_, ", u0=", u0)
println("T0.025($(n-1))=", ta2)
println("标准差=", s_)
println("T(u0=$u0)=", t_)
