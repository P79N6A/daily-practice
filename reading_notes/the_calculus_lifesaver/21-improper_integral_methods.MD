# 反常积分：如何解题

## 一、如何开始
给定一个反常积分```∫[a, b] f(x)dx```，根据定义被积函数 f 在区间 [a,b] 上至少有一个瑕点。瑕点经常出现在 f 的破裂点，如有垂直渐近线的点处，还出现在 ∞ 和 -∞ 处。例如：
```
∫[∞, -∞] 1/(x^2-1) dx
```
在 ∞ 和 -∞ 处有瑕点（只要包含它们，就一定有瑕点），同样在 x=1 和 x=-1 处也有瑕点（因为函数在这些点无定义）

#### ◉ 1 拆分积分
基本对策：

1. **确定** 区间 [a,b] 上 **所有瑕点**；
2. **将积分拆分** 成若干积分之和，便利每一个积分至多有一个瑕点，并使这些瑕点作为相应积分的上限或下限；
3. **分别讨论每个积分，如果任一积分发散，则整个积分发散。** 原反常积分收敛的唯一情况是每个积分都收敛；
4. **如果没有瑕点，则积分收敛！**

#### ◉ 2 如何处理负值函数
方法：

1. 如果被积函数 f(x) 在区间 [a,b] 上既有正值又有负值，你应该考虑使用绝对收敛判别法：```如果 ∫[a,b] |f(x)| dx 收敛，那么 ∫[a,b] f(x) dx 也收敛```，需要注意：**绝对收敛判别法只能用于判断收敛，不能用于判断积分发散！**
2. 假设被积函数 f(x) 在区间 [a,b] 上恒负（或为0），即在 [a,b] 上 f(x) <= 0。则：```∫[a, b] f(x) dx = -∫[a, b] (-f(x)) dx```
3. 如果以上两种情形都不适用，可以使用反常积分的正式定义试一下，如 20.6 节的：```∫[a,∞] cos(x) dx```

## 二、积分判别法总结
可以支配的最有价值的工具是 **比较判别法、极限判别法 和 P判别法**。

在正面所有的判别法中，被积函数 f(x) 被假定为在积分区间上恒正：

* **比较判别法的发散情况：** 若认为 ∫[a,b] f(x) dx 发散，找一个积分也是发散的较小函数，即找一个使得在区间 (a,b) 上有 f(x) >= g(x) 的非负函数 g(x)，且∫[a, b] g(x) dx 发散，则有：```∫[a, b] f(x) dx >= ∫[a, b] g(x) dx = ∞```，因此 ∫[a,b] f(x) dx 发散。
* **比较判别法的收敛情况：** 若认为 ∫[a,b] f(x) dx 收敛，找一个积分也是收敛的较大函数，即找一个使得在区间 (a,b) 上有 f(x) <= g(x) 的函数 g(x)，且∫[a, b] g(x) dx 收敛，则有：```∫[a, b] f(x) dx <= ∫[a, b] g(x) dx < ∞```，因此 ∫[a,b] f(x) dx 收敛。

**极限比较判别法** 是比较判别法的替代形式。使用该判别法的重点，是找到一个和被积函数在瑕点附近散敛性一致的函数。如 20.4.1 的定义：
```
当 x~a 时，f(x)~g(x)，等价于 lim(x→a) f(x)/g(x) = 1
```
根据极限比较判别法，g 与 f 散敛性一致，更准确地，下面是该判别法针对瑕点的有限和无穷两种情况的判别：
* **极限比较判别法中瑕点为无穷的情形：** 找一个在区间 [a,∞) 上没有瑕点、形式较简单的非负函数 g，且有当 x→∞ 时，则
    * 若 ∫[a,b] g(x) dx 收敛，则 ∫[a,b] f(x) dx 收敛；
    * 若 ∫[a,b] g(x) dx 发散，则 ∫[a,b] f(x) dx 发散；
    * 将区间 [a,∞) 换为 (-∞,b] 也成立； 瑕点为积分区间左端处的有限值 a，也成立。
* **极限比较判别法中瑕点为有限值的情形：** 找一个在区间 (a,b] 上没有瑕点、形式较简单的非负函数 g，且有当 x→a 时，则
    * 若 ∫[a,b] g(x) dx 收敛，则 ∫[a,b] f(x) dx 收敛；
    * 若 ∫[a,b] g(x) dx 发散，则 ∫[a,b] f(x) dx 发散；
    * 针对唯一的瑕点在右端点 x=b，且有当 x→b 时（而不是a）时 f(x)~g(x)的情形，结论相同。

因此 ，需要选择一个合适的函数 g 来做比较。通过选择 g(x) 为 1/x^p 的形式，并选择合理的 p 值，能够解决很多问题。这类函数积分的敛散性可以准确地由 **P判别法** 描述：
* **P判别法，∫[,∞] 的情况：** 对于任何有限值 a > 0，积分 ```∫[a,∞] 1/x^p dx``` 在 p > 1 时是收敛的，在 p <= 1 时是发散的；
* **P判别法，∫[0,] 的情况：** 对于任何有限值 a > 0，积分 ```∫[0,a] 1/x^p dx``` 在 p < 1 时是收敛的，在 p >= 1 时是发散的；

## 三、常数函数在 ∞ 和 -∞ 附近的表现

#### ◉ 1 多项式和多项式型函数在 ∞ 和 -∞ 附近的表现


多项式在 x→∞ 和 x→-∞ 时，都是**最高次项**起决定作用。更准确地说，设 p 为多项式，则
```
若 p(x) 的最高次项是 a*x^n，则当 x→∞ 和 x→-∞ 时，有 p(x) ~ a*x^n
```

#### ◉ 2 三角函数在 ∞ 和 -∞ 附近的表现

我们能知道的有用结论是，对于任意实数 A 有
```
sin|A| <= 1
cos|A| <= 1
```

#### ◉ 3 指数在 ∞ 和 -∞ 附近的表现

一个非常有用的原理：**指数比多项式增长得快。** 表现形式：
```
lim(x→∞) x^n/e^x = 0    其中 n 是任意正数
```
以上极限有当 x→∞ 时 f(x)→0。当 x>=0 时，y = f(x) 的图像必然有最大高度，定义为 C。一个有用的不等式：
```
对所有的 x>0，e^-x <= C/x^n  (e也可用常数替代)
```

#### ◉ 4 对数在 ∞ 附近的表现

1. 我们不考虑对数在 -∞ 附近的情形，因为负数不能取对数；
2. **对数在 ∞ 处增长的很慢。** 事实上，对数比 x 的任何次幂增长都慢。用符号表示就是，若 a>0 是某个正数，不管多小都有：``` lim(x→∞) ln(x)/x^a = 0```。由 21.3.3 节的论述可得，必有一个常数 C 使得：
```
对所有 x>1，ln(x) <= C*x^a
```

## 四、常见函数在 0 附近的情形

#### ◉ 0附近的多项式和多项式型函数
对多项式，**最低次幂在 x→0 时起决定作用。** 这与 x→∞ 时的情况正好相反，更准确的，设 p 为多项式，则有
```
若 p(x) 的最低次项是 b*x^m，则当 x→0，p(x)~b*x^m
```

#### ◉ 0附近的三角函数
有用的结论（可用于替换）：
```
当 x→0，sin(x)~x，tan(x)~x，且 cos(x)~1
```
这些只是第7章讨论过的极限的另一种描述：
```
lim(x→0) sin(x)/x = 1, lim(x→0) tan(x)/x = 1, lim(x→0) cos(x) = 1
```

#### ◉ 0附近的指数函数
**指数函数对 0 没有作用。** 更准确的，```e^x ~ 1 和 e^-x ~ 1，当 x→0```，这其实是 ```lim(x→0) e^x = 1 和 lim(x→0) e^(-x) = 1``` 的另一种说法。

9.4.2 节中的经典极限：
```
lim(x→0) (e^x - 1)/x = 1    所以，当 x→0，e^x - 1 ~ x
```
#### ◉ 0附近的对数函数
无论 a > 0 有多小，都存在常数 C 使得
```
对所有 0 < x < 1 成立，|ln(x)| <= C/x^a
```
这是由 9.4.6 节的极限 ```lim(x→0) x^a * ln(x) = 0``` 推出来的。

#### ◉ 0附近的更一般函数
基本观点：若一个函数有在 0 附近收敛于该函数的麦克劳林级数，则函数在 x→0 时渐近于级数的最低次项，即
```
f(x) = a[n]*x^n + a[n+1]*x^(n+1) + ...，当 x→0，则 f(x) ~ a[n]*x^n
```
例子：
```
∫[0,1] dx/(1-cos(x))
```
我们知道当 x→0 时 cos(x)~1，求 1-cos(x) 就需要用到 cos(x) 的麦克劳林级数：
```
cos(x) = 1 - x^2/2! + x^4/4! - ....
则推出
1 - cos(x) = x^2/2! + x^4/4! - ....
```
由上边原理可知，右边最低次项起决定作用，有：当 x→0 时，1-cos(x) ~ x^2/2，则
```
∫[0,1] dx/(1-cos(x)) = 2 * ∫[0,1] dx/x^2
```

## 五、如何应对不在 0 或 ∞ 处的瑕点
若瑕点出现在有限值而非 0 处，做换元。具体情况如下：
* 若积分 ∫[a,b] f(x)dx 的唯一瑕点出现在 x=a 处，做代换 t=x-a，注意 dt=dx。新的积分则只有 0 一个瑕点
* 若积分 ∫[a,b] f(x)dx 的唯一瑕点出现在 x=b 处，做代换 t=b-x，注意 dt=-dx，用多出的符号做积分上下限交换。新的积分则只有 0 一个瑕点。
