# saint julia

## 中心极限定理

> 只要样本量足够大(n>=30)，不论原来的总体是否符合正态分布，样本均值的抽样分布都将趋于正态分布，其分布的数学期望为总体均值u，方差为总体方差的1/n。

## 符号

* ∃   It means "there exists"
* ∀   Which means "for all"
* ε   理解为 {String.Empty}, 读作epsilon
    * ε-closure 是闭包函数，表示是当前状态通过ε边可以到达的所有状态的集合
* Σ   读作 sigma
* γ   读作 gamma

> When used in an expression such as
```
∃x s.t. x > 0
```
It means "There exists a number x such that x is greater than 0."
Its counterpart is ∀, which means "for all". It's used like this:
```
∀x, x > 0
```
Which means "For any number x, it is greater than 0."

## 理论

1. [克莱尼星号](https://zh.wikipedia.org/wiki/%E5%85%8B%E8%8E%B1%E5%B0%BC%E6%98%9F%E5%8F%B7)
2. [确定有限状态机](https://zh.wikipedia.org/wiki/%E7%A1%AE%E5%AE%9A%E6%9C%89%E9%99%90%E7%8A%B6%E6%80%81%E8%87%AA%E5%8A%A8%E6%9C%BA)

## Y Combinator
```
Y := λf.(λx.(f (x x)) λx.(f (x x)))
```
[不动点组合子](https://zh.wikipedia.org/zh-hans/%E4%B8%8D%E5%8A%A8%E7%82%B9%E7%BB%84%E5%90%88%E5%AD%90)

## Context-Free Grammar
* grammar:
```
F → id
F → ( E )
T → T * F
T → F
E → E + T
E → T
```

* expansion:
```
a + b * c
--------->
E → E1 + T1

E1 → T2
T2 → F1 → (E2) → (a + b)

T1 → F2 → c
--------->
(a + b) * c
```

* CFG > LR(1) > LL(1) > RG
