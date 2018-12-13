# 3 数据的描述

## 3.2 如何用少量数字来描述数据

### 3.2.1 数据的“位置”
和位置有关的统计量称为 **位置统计量（location statistic）**：
* 均值（mean）/ 样本均值（sample mean）
* 中位数（median）
* 分位数（quantile）
* 众数（mode）

### 3.2.2 数据的“尺度”
数据是否“均”由**尺度统计量（scale statistic）** 来描述：
* 极差（range）
* 标准差（standard deviation）
* 标准误差（standard error）/ 均值误差（standard error for mean）

### 3.2.3 数据的标准得分
不同样本中均值和标准差不同的数据，需经过标准化后方可比较。一种标准化方法：将样本的原始观测值与样本均值之差除以样本标准差，得到 **标准复仇（standard score）**：
```
𝒛i = (𝓍i - ẋ)/s
```
将各样本的观测值转换成相应的标准复仇，就可以进行比较了。
