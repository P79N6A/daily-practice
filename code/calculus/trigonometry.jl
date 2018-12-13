# 三角函数
#   在一个参考角度为θ的直角三角形中
#   x为横边，y为纵边，r为斜边
#   ----------------------------------
#   sin(θ)=y/r          cos(θ)=x/r
#   tan(θ)=y/x          cot(θ)=x/y
#   sec(θ)=r/x          csc(θ)=r/y
#   ------------------------------------
#   tan(θ)=sin/cos      cot(θ)=cos/sin
#   sec(θ)=1/cos        csc(θ)=1/sin
#   ------------------------------------
#   sin^2+cos^2=1       (毕达哥拉斯定理)
#   1+tan^2=sec^2
#   cot^2+1=csc^2
#   ------------------------------------
#   sin(θ)=cos(π/2-θ)   cos(θ)=sin(π/2-θ)
#   tan(θ)=cot(π/2-θ)   cot(θ)=tan(π/2-θ)
#   sec(θ)=csc(π/2-θ)   csc(θ)=sec(π/2-θ)
#   ------------------------------------
#   sin(A+B)=sin(A)cos(B)+sin(B)cos(B)
#   cos(A+B)=cos(A)cos(B)-sin(A)sin(B)
#   -> 倍角公式
#   sin(2X)=2sin(X)cos(X)
#   cos(2X)=2cos(X)^2-1 (根据毕达歌拉斯定理)
#   ->
#   sin(4X)=2(2sin(X)cos(X))(2cos(X)^2-1) = 8sin(X)cos(X)^3-8cos(X)^2+1
#   cos(4X)=2cos(2X)^2-1 = 8cos(X)^3-8cos(X)^2+1
module Trigonomegtry
    export angle_to_radian, radian_to_angle

    # 角度转弧度
    function angle_to_radian(angle)
        π/180 * angle
    end

    # 弧度转角度
    function radian_to_angle(radian)
        radian / (π/180)
    end
end
