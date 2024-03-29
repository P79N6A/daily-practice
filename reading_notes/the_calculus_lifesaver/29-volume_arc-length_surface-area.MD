# 体积、弧长和表面积

## 一、旋转体的体积

#### 1 圆盘法
这是一个高 dx 个单位，半径为 y 个单位的薄盘，用它绕 x 轴旋转一个薄盘，将会得到一个球。
把它看作一个圆柱体，半径为 y 个单位，高为 dx 个单位，那么它的体积为 ```π*y^2*dx``` 立方单位。

在最大厚度（高度) dx 趋于 0 的积分，就是这个圆盘的体积：
```
∫[a,b] π*y^2 dx =  π * ∫[a,b] y^2 dx  
```

#### 2 壳法
这是一个厚度为 dx 单位，半径为 x 个单位的薄壳，用它绕 y 轴旋转，将会得到一个薄罐。

在薄罐厚度 dx 趋于 0 的积分，就是这一堆壳的体积：
```
∫[a,b] 2π * x*y dx = 2π * ∫[a,b] x*y dx
```

#### 3 总结和变式
(1) 原理：
1. 若将曲线 y=f(x) 下方，x=a 和 x=b 之间的区域绕 x 轴旋转，则可应用**圆盘法**且体积等于：
```
∫[a,b] π*y^2 dx 立方单位
```
2. 若将曲线 y=f(x) 下方，x=a 和 x=b 之间的区域绕 y 轴旋转，则可应用**壳法**且体积等于：
```
∫[a,b] 2π*x*y dx 立方单位
```

(2) 运用法则：在将区域切成小条后，有
* 若每个小条的短边平行于旋转轴，运用圆盘法
* 若每个小条的短边垂直于旋转轴，运用壳法

(3) 关于轴的准则：
```
若区域在曲线和 y 轴之间，交换 x 和 y
```

#### 4 两曲线间的区域
结论：
```
若区域在两曲线之间，求相应的旋转体体积之差
```

#### 5 绕平行于坐标轴的轴旋转
一般思想：
```
若旋转轴是 x=h，用(x-h)（若 x<h 则用 (h-x)）代换 x；
若旋转轴是 y=h，用(y-h)（若 y<h 则用 (h-y)）代换 y。
```

## 二、一般固体体积
解题思路：

1. 选定一个轴
2. 求一个在轴上点 x 处的典型横截面面积，称该面积为 A(x) 平方单位
3. 则，若 V 为固体的体积（立方单位），我们有
```
V = ∫[a,b] A(x)dx   其中 [a,b] 是完全覆盖固体的 x 取值范围
```

(l,x) 为横截面长度和高度，(L,h) 为底的长度和高度。底和横截面的面积之比是其长度之比的平方：
```
A/A(x) = (L/l)^2
化简可得
A(x) = A*x^2/h^2
```

## 三、弧长
一些公式：

1. 若 y=f(x)，且 x 在 a 到 b 间取值，则在上述被积函数中取因子 (dx)^2 并将其移出平方根得：
```
弧长 = ∫[a,b] sqrt(1+(dy/dx)^2 dx)  （标准形式）
关于 f 可另写为：
弧长 = ∫[a,b] sqrt(1+f'(x)^2) dx
```
2. 假定给定关于 y 的 x。若 x = g(y) 且 y 在 A 到 B 间取值，则取因子 (dy)^2 得：
```
弧长 = ∫[A,B] sqrt(1+(dx/dy)^2 dy)  （关于 y）
也可写为
弧长 = ∫[A,B] sqrt(1+g'(y)^2) dy
```
3. 参数形式呢？这意味着 x 和 y 是关于参数 t 的函数，t 在 t0 到 t1 间取值（参见27.1节）。我们将量 (dx)^2 看作 (dx/dt)^2(dt)^2；y 同理，然后我们将 (dt)^2 移出平方根得到一个有用的公式：
```
弧长 = ∫[t0,t1] sqrt((dx/dt)^2+(dy/dt)^2) dt (参数版)
```
4. 最后这个公式的特殊情况发生在极坐标情形。我们知道 ```x=r*cos(Θ)``` 且 ```y=r*sin(Θ)```，所以用 f(Θ) 代换 r，我们有 ```x=f(Θ)*cos(Θ)``` 和 ```y=f(Θ)*sin(Θ)```。这里 Θ 的角色是参数，我们可以使用上面参数版的弧长公式（t代换为Θ)。我们需要知道 dy/dΘ 和 dx/dΘ 是什么。根据乘积法则：
```
dx/dΘ = f'(Θ)*cos(Θ) - f(Θ)*sin(Θ)
dy/dΘ = f'(Θ)*sin(Θ) - f(Θ)*cos(Θ)
```
对这两个式子取平方并相加，会发现所有项都消掉了。还有两套 ```sin^2(Θ)+cos^2(Θ)```，这些项可用 1 代换。综上，可得公式：
```
弧长 = ∫[Θ0,Θ1] sqrt(f(Θ)^2 + f'(Θ)^2) dΘ    (极坐标函数，r=f(Θ))
```

#### ◉ 参数化和速率
参数情况下求弧长的积分：
```
速率 = sqrt((dx/dt)^2 + (dy/dt)^2)
```

## 四、旋转体的表面积
关于 x 轴旋转的原型公式：
```
表面积 = ∫[?,?] 2π*y * sqrt((dx)^2 + (dy)^2) (关于 x 轴旋转)
```
关于 y 轴旋转的原型公式：
```
表面积 = ∫[?,?] 2π*x * sqrt((dx)^2 + (dy)^2) (关于 y 轴旋转)
```

应用：

(1) 假设我们关于 x 轴旋转曲线 y=f(x)，其中 x 取值范围为 a 到 b。我们从第一个原型公式中的被积函数中提出因子 (dx)^2 并将其提到根号之外，就像在讨论弧长所做的一样，得：
```
表面积 = ∫[a,b] 2π*y * sqrt(1+(dy/dx)^2) dx    (关于 x 轴)
```

(2) 若我们关于 y 轴旋转相同的曲线，可对另一个原型公式采用相同的处理，得出：
```
表面积 = ∫[a,b] 2π*x * sqrt(1+(dy/dx)^2) dx    (关于 y 轴)
或写成
表面积 = ∫[a,b] 2π*x * sqrt(1+(f'(x))^2) dx
```

(3) 当然也有参数形式，若 x 和 y 是参数 t 的函数，其中 t 的取值范围为 t0 到 t1，则可推导：
```
表面积 = ∫[t0,t1] 2π*y * sqrt((dx/dt)^2 + (dy/dt)^2) dt   （参数版，关于 x 轴）
和
表面积 = ∫[t0,t1] 2π*x * sqrt((dx/dt)^2 + (dy/dt)^2) dt   （参数版，关于 y 轴）
```
