include("./goodness_of_fit_test.jl")
using GoodnessOfFitTest

fo_arr = [374, 344]
fe_arr = [565, 153]
println("观察值频数:", fo_arr)
println("期望值频数:", fe_arr)
println("自由度:", df(fo_arr))
println(x2(fo_arr, fe_arr))
