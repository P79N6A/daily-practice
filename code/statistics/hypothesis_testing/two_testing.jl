## 两个样本的均值检验
##   o1, o2已知，使用z统计量
function z_o(x1, o1, u1, n1, x2, o2, u2, n2)
    (x1-x2)-(u1-u2)/sqrt(o1^2/n1+o2^2/n2)
end
##  o1, o2未知但相等，且n较小，使用t统计量
function t_s(x1, u1, n1, x2, u2, n2, sp)
    (x1-x2)-(u1-u2)/sp*sqrt(1/n1+1/n2)
end
##  o1，o2未知且不相等，且n较小，使用f统计量
function f(s1, n1, s2, n2)
    sn1 = s1^2/n1
    sn2 = s2^2/n2
    (sn1+sn2)^2 / ((sn1)^2/(n1-1) + (sn2)^2/(n2-1))
end
