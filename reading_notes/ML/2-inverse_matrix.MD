# 矩阵的逆

如果矩阵 A 可逆，则
```
A(-1) = A(*) / |A|，其中 A(*) 是 A 的伴随矩阵，|A| 是矩阵的行列式
```

**逆矩阵 还可以使用 高斯消除法来得出：即用一个单位矩阵，协同原矩阵进行同步操作，直到原矩阵变成单位矩阵，则协同的单位矩阵将变成逆矩阵**

## 伴随矩阵
### 2X2 矩阵 伴随矩阵
一个 2X2 矩阵 A
```
[a   b]
[c   d]
```
的伴随矩阵是
```
[d  -b]
[-c  a]
```
### 3X3 矩阵 伴随矩阵
```
[a11  a12  a13]
[a21  a22  a23]
[a31  a32  a33]
```
的伴随矩阵是
```
[+|a22 a23|  -|a12 a13|  +|a12 a13|]
[ |a32 a33|   |a32 a33|   |a22 a23|]
[                                  ]
[-|a21 a23|  +|a11 a13|  -|a11 a13|]
[ |a31 a33|   |a31 a33|   |a21 a23|]
[                                  ]
[+|a21 a22|  -|a11 a12|  +|a11 a12|]
[ |a31 a32|   |a31 a32|   |a21 a22|]
```
每一项都是转置后（沿对角线翻转）的当前行列的余子式，进一步的需要计算每一项的行列式，以第一项为例，即：
```
a22*a33 - a23*a32
```

#### 余子式
每一项的余子式就是去除当前项行和列，剩余的行列形成的矩阵

## 行列式
行列式（Determinant）是数学中的一个函数，将一个nxn的矩阵 A映射到一个标量，记作 |A|(或det(A))
### 二维向量的行列式
```
det(X, X') = |a b| = ad - bc
             |c d|  
```
### 三维向量的行列式
```
det(X, X', X'') = |x  x'  x''| = xy'z'' + x'y''z + x''yz' - xy''z' - x'yz'' - x''y'z
                  |y  y'  y''|
                  |z  z'  z''|
```
