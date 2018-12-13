# Functors, Applicatives, And Monads

[http://adit.io/posts/2013-04-17-functors,_applicatives,_and_monads_in_pictures.html](Functors, Applicatives, And Monads In Pictures)

## 函子 Functor

```haskell
fmap :: (a -> b) -> f a -> fb
```
在 Haskell 中，我们把抽象出来的这个函数记作 fmap，而被抽象出来的类型（列表和 Maybe 等）叫作函子。

```haskell
fmap :: (a -> b) -> [a] -> [b]
fmap = map

fmap :: (a -> b) -> Maybe a -> Maybe b
fmap f (Just a) = Just (f a)
fmap _ Nothing = Nothing
```

函子这个类型类提供了一个很重要的抽象 ：容器类型。
也就是说，函子类型基本上都可以看作某种类型的容器。如：
列表和 Maybe、Vector、Array、Either 等。

## 概括解释

1. Functor: 一个容器中的值与一个运算符的结合，运算后的结果再放回容器中（不同范畴）
2. Applicative：一个容器中的值与一个容器中的运算符的结合，运算时将二者取出，结果放回容器返回
3. Monad: 一个Monad容器中的值与一个返回 Monad 容器的运算符的结合，运算后的结果还是 Monad 容器类型（同一范畴）
