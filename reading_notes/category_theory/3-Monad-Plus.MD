# MonadPlus

```haskell
class (Alternative m, Monad m) => MonadPlus m where
  mzero :: m a
  mplus :: m a -> m a -> m a
```

和单位单群、选择应用函子类似，上面两个值需要满足如下条件：

* 单位元： ```mzero `mplus` ma ≣ ma `mplus` mzero ≣ ma```
* 结合律： ```ma `mplus` (mb `mplus` mc) ≣ (ma `mplus` mb) `mplus` mc```
