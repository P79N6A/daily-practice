# 3. Duality

## 3.1 The duality principle

axioms of category:
```haskell
dom(1A) = A 
cod(1A) = A
f ◦ 1dom(f) = f 
1cod(f) ◦ f = f
dom(g ◦ f) = dom(f) 
cod(g ◦ f) = cod(g)
h ◦ (g ◦ f) = (h ◦ g) ◦ f
```