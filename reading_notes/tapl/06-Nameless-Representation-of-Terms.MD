## 6.2 Shifting and Substitution

When a substitution goes under a λ-abstraction, as in ```[1 → s](λ.2) (i.e., [x → s](λy.x)```, assuming that 1 is the index of x in the outer context), the context in which the substitution is taking place becomes one variable longer than the original; we need to increment the indices of the free variables in s so that they keep referring to the same names in the new context as they did before. 

So, when calculating "↑dc(t), we know that the term t comes from inside c-many binders in the original argument to "d. Therefore all identifiers k < c in t are bound in the original argument and should not be shifted, while identifiers k ≥ c in t are free and should be shifted.

## 6.3 Evaluation

The only slightly subtle point is that reducing a redex “uses up” the bound variable: when we reduce ```((λx.t12) v2) to [x → v2]t12```, the bound variable x disappears in the process. Thus, we will need to renumber the variables of the result of substitution to take into account the fact that x is no longer part of the context. For example:
```
    (λ.1 0 2) (λ.0)     →   0 (λ.0) 1             (not 1 (λ.0) 2).
```

Similarly, we need to shift the variables in v2 up by one before substituting into t12, since t12 is defined in a larger context than v2. Taking these points into account, the beta-reduction rule looks like this:
```
    (λ.t12) v2      →     ↑−1 ([0 → ↑1 (v2)]t12)            (E-AppAbs)
```