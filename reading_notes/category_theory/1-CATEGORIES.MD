# CATEGORIES

### Definition

1. identity law: for any arrows f, g such that cod(f) = b = dom(g)
```
idb ° f = f
g ° idb = g
```
2. associative law: for any arrows f, g, h such that dom(f) = cod(g) and dom(g) = cod(h)
```
( f ° g ) ° h = f ° ( g ° h )
```

### Functor

A functor
```
F : C → D
```
between categories C and D is a mapping of objects to objects and arrows to arrows, in such a way that:
* (a) F(f : A → B) = F(f) : F(A) → F(B),
* (b) F(g ◦ f) = F(g) ◦ F(f),
* (c) F(1A) = 1F(A).

### Applicative

```haskell
Applicative m => class Monad m where
  return :: a -> m a
  return = pure

  join :: m (m a) -> m a
```

```haskell
(>>=) :: m a -> (a -> m b) -> m b
x >>= f = join $ fmap f x
```
