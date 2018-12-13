# category-theory-for-programmers-the-preface

## 1. Category: The Essence of Composition

### 1.1 Properties of Composition

1. Composition is associative. If you have three morphisms, f, g, and h, that can be composed (that is, their objects match end-to-end), you don’t need parentheses to compose them. In math notation this is expressed as:
```haskell
h∘(g∘f) = (h∘g)∘f = h∘g∘f
```

2. For every object A there is an arrow which is a unit of composition. This arrow loops from the object to itself. Being a unit of composition means that, when composed with any arrow that either starts at A or ends at A, respectively, it gives back the same arrow. The unit arrow for object A is called idA (identity on A). In math notation, if f goes from A to B then
```haskell
f∘idA = f
```
and
```haskell
idB∘f = f
```

## 2. Types and Functions

### 2.1 What Are Types?

a function declared as:
```haskell
f :: Bool -> Bool
```
may return True, False, or _|_; the latter meaning that it would never terminate.

## 3. Categories Great and Small

### 3.1 Orders
* preorder:
  1. have identity morphisms
  2. every object is less than or equal to itself
  3. If a <= b and b <= c then a <= c
  4. composition associative
* partial order:
  1. preorder
  2. if a <= b and b <= a then a == b

### 3.2 Monoid as Set
a monoid is defined as a set with a binary operation.

They show up as:
  strings, lists, foldable data structures, futures in concurrent programming, events in functional reactive programming

* the values of arguments are sometimes called points
* Function equality without specifying the arguments is described as **point-free**

### 3.3 Monoid as Category

Every monoid can be described as a single object category with a set of morphisms that follow appropriate rules of composition.

## 4. Products and Coproducts

### 4.1 Isomorphisms (同构)

An isomorphism is an invertible morphism;
or a pair of morphisms, one being the inverse of the other.

### 4.2 Product

A product of two objects a and b is the object c equipped with two projections such that for any other object c’ equipped with two projections there is a unique morphism m from c’ to c that factorizes those projections.

A (higher order) function that produces the factorizing function m from two candidates is sometimes called the factorizer. In our case, it would be the function:
```
factorizer :: (c -> a) -> (c -> b) -> (c -> (a, b))
factorizer p q = \x -> (p x, q x)
```

### 4.3 CoProduct

A coproduct of two objects a and b is the object c equipped with two injections such that for any other object c’ equipped with two injections there is a unique morphism m from c to c’ that factorizes those injections.

Just as we’ve defined the factorizer for a product, we can define one for the coproduct. Given a candidate type c and two candidate injections i and j, the factorizer for Either produces the factoring function:
```haskell
factorizer :: (a -> c) -> (b -> c) -> Either a b -> c
factorizer i j (Left a)  = i a
factorizer i j (Right b) = j b
```

## 5. Simple Algebraic Data Types

### 5.1 Product Types

The canonical implementation of a product of two types in a programming language is a pair. In Haskell, a pair is a primitive type constructor.

Programming with tuples and multi-argument constructors can get messy and error prone — keeping track of which component represents what. It’s often preferable to give names to components. A product type with named fields is called a record in Haskell, and a struct in C.

### 5.2 Records

### 5.3 Sum Types

### 5.4 Algebra of Types

Taken separately, product and sum types can be used to define a variety of useful data structures, but the real strength comes from combining the two. Once again we are invoking the power of composition.

Here’s a translation table with some entries of interest:

|  Numbers  |   Types                        |
|:----------|:-------------------------------|
| 0         | Void                           |
| 1         | ()                             |
| a + b     | Either a b = Left a | Right b  |
| a * b	    | (a, b) or Pair a b = Pair a b  |
| 2 = 1 + 1 | data Bool = True | False       |
| 1 + a	    | data Maybe = Nothing | Just a  |

Logical and and or also form a semiring, and it too can be mapped into type theory:

|  Logic    |   Types                        |
|:----------|:-------------------------------|
| false     | Void                           |
| true      | ()                             |
| a \|\| b  | Either a b = Left a \| Right b |
| a && b    | (a, b)                         |


## 5. Functors

A functor is a mapping between categories. Given two categories, C and D, a functor F maps objects in C to objects in D — it’s a function on objects.
If a is an object in C, we’ll write its image in D as F a (no parentheses). But a category is not just objects — it’s objects and morphisms that connect them. A functor also maps morphisms — it’s a function on morphisms.
But it doesn’t map morphisms willy-nilly — it preserves connections. So if a morphism f in C connects object a to object b,
```haskell
f :: a -> b
```
the image of f in D, F f, will connect the image of a to the image of b:
```haskell
F f :: F a -> F b
```

### 5.1 Functors in Programming

#### 5.1.1 The Maybe Functor

The definition of Maybe is a mapping from type a to type Maybe a:

```haskell
data Maybe a = Nothing | Just a
```
#### 5.1.2 Equational Reasoning
#### 5.1.3 The List Functor

Generic containers are parameterized by the type of the elements they store, so let’s look at a very simple container, the list:
```haskell
data List a = Nil | Cons a (List a)
```
We have the type constructor List, which is a mapping from any type a to the type List a. To show that List is a functor we have to define the lifting of functions: Given a function a->b define a function List a -> List b:
```haskell
fmap :: (a -> b) -> (List a -> List b)
```
Putting it all together, here’s the instance declaration for the list functor:
```haskell
instance Functor List where
    fmap _ Nil = Nil
    fmap f (Cons x t) = Cons (f x) (fmap f t)
```

#### 5.1.4 The Reader Functor
So here’s the implementation of our fmap:
```haskell
instance Functor ((->) r) where
  fmap f g = f . g
```
It just works! If you like terse notation, this definition can be reduced further by noticing that composition can be rewritten in prefix form:
```haskell
fmap f g = (.) f g
```
and the arguments can be omitted to yield a direct equality of two functions:
```haskell
fmap = (.)
```
This combination of the type constructor (->) r with the above implementation of fmap is called the reader functor.

#### 5.1.5 Functors as Containers
#### 5.1.6 Functor Composition

## 6. Functoriality

### 6.1 Bifunctors
On objects, a bifunctor maps every pair of objects, one from category C, and one from category D, to an object in category E.

That’s pretty straightforward. But functoriality means that a bifunctor has to map morphisms as well. This time, though, it must map a pair of morphisms, one from C and one from D, to a morphism in E.

### 6.2 Product and Coproduct Bifunctors

** A monoidal category defines a binary operator acting on objects, together with a unit object. **

### 6.3 Functorial Algebraic Data Types

We’ve seen several examples of parameterized data types that turned out to be functors — we were able to define ```fmap``` for them. Complex data types are constructed from simpler data types. In particular, algebraic data types (ADTs) are created using sums and products. We have just seen that sums and products are functorial. We also know that functors compose. So if we can show that the basic building blocks of ADTs are functorial, we’ll know that parameterized ADTs are functorial too.

### 6.4 The Writer Functor

Morphisms in Kleisli category were represented as “embellished” functions returning the Writer data structure.
```haskell
type Writer a = (a, String)
```

### 6.5 Covariant and Contravariant Functors

Here’s the typeclass defining a contravariant functor (really, a contravariant endofunctor) in Haskell:
```haskell
class Contravariant f where
    contramap :: (b -> a) -> (f a -> f b)
```
Our type constructor Op is an instance of it:
```haskell
instance Contravariant (Op r) where
    -- (b -> a) -> Op r a -> Op r b
    contramap f g = g . f
```

The definition of contramap for Op may be made even terser, if you notice that it’s just the function composition operator with the arguments flipped. There is a special function for flipping arguments, called flip:
```haskell
flip :: (a -> b -> c) -> (b -> a -> c)
flip f y x = f x y
```
With it, we get:
```haskell
contramap = flip (.)
```

### 6.6 Profunctors

Because a contravariant functor is equivalent to a covariant functor from the opposite category, a profunctor is defined as:
```haskell
Cop × D → Set
```
Now we can assert that the function-arrow operator is an instance of a Profunctor:
```haskell
instance Profunctor (->) where
  dimap ab cd bc = cd . bc . ab
  lmap = flip (.)
  rmap = (.)
```

## 7. Function Types

A set of morphisms between two objects in any category is called a hom-set.

### 7.1 Universal Construction
Formally:

A function object from a to b is an object a⇒b together with the morphism
```haskell
eval :: ((a⇒b) × a) -> b
```
such that for any other object z with a morphism
```haskell
g :: z × a -> b
```
there is a unique morphism
```haskell
h :: z -> (a⇒b)
```
that factors ```g``` through ```eval```:
```haskell
g = eval ∘ (h × id)
```

### 7.2 Currying

Currying is essentially built into the syntax of Haskell. A function returning a function:
```haskell
a -> (b -> c)
```
is often thought of as a function of two variables. That’s how we read the un-parenthesized signature:
```haskell
a -> b -> c
```

It’s trivial to convert between the two representations, and the two (higher-order) functions that do it are called, unsurprisingly, curry and uncurry:
```haskell
curry :: ((a, b)->c) -> (a->b->c)
curry f a b = f (a, b)
```
and
```haskell
uncurry :: (a->b->c) -> ((a, b)->c)
uncurry f (a, b) = f a b
```

### 7.3 Exponentials

In mathematical literature, the function object, or the internal hom-object between two objects ```a``` and ```b```, is often called the exponential and denoted by ```bª```. Notice that the argument type is in the exponent. This notation might seem strange at first, but it makes perfect sense if you think of the relationship between functions and products. We’ve already seen that we have to use the product in the universal construction of the internal hom-object, but the connection goes deeper than that.

This is best seen when you consider functions between finite types — types that have a finite number of values, like Bool, Char, or even Int or Double. Such functions, at least in principle, can be fully memoized or turned into data structures to be looked up. And this is the essence of the equivalence between functions, which are morphisms, and function types, which are objects.

### 7.4 Cartesian Closed Categories

Although I will continue using the category of sets as a model for types and functions, it’s worth mentioning that there is a larger family of categories that can be used for that purpose. These categories are called cartesian closed, and ```Set``` is just one example of such a category.

A cartesian closed category must contain:

1. The terminal object,
2. A product of any pair of objects, and
3. An exponential for any pair of objects.

The terminal object and the product have their duals: the initial object and the coproduct. A cartesian closed category that also supports those two, and in which product can be distributed over coproduct

```haskell
a × (b + c) = a × b + a × c
(b + c) × a = b × a + c × a
```

### 7.5 Exponentials and Algebraic Data Types

The interpretation of function types as exponentials fits very well into the scheme of algebraic data types. It turns out that all the basic identities from high-school algebra relating numbers zero and one, sums, products, and exponentials hold pretty much unchanged in any bicartesian closed category theory for, respectively, initial and final objects, coproducts, products, and exponentials. We don’t have the tools yet to prove them (such as adjunctions or the Yoneda lemma), but I’ll list them here nevertheless as a source of valuable intuitions.

#### 7.5.1 Zeroth Power

```haskell
a⁰ = 1
```

In Haskell, we replace 0 with ```Void```; 1 with the unit type ```()```; and the exponential with function type. The claim is that the set of functions from ```Void``` to any type ```a``` is equivalent to the unit type — which is a singleton. In other words, there is only one function ```Void->a```. We’ve seen this function before: it’s called ```absurd```.

#### 7.5.2 Powers of One

```haskell
1ª = 1
```

In Haskell, there is only one function from any type ```a``` to unit. We’ve seen this function before — it’s called ```unit```. You can also think of it as the function ```const``` partially applied to ```()```.

#### 7.5.3 First Power

```haskell
a¹ = a
```

#### 7.5.4 Exponentials of Sums

```haskell
aᵇ⁺ᵓ = aᵇ × aᵓ
```

### 7.6 Curry-Howard Isomorphism

Let’s take the function eval we have introduced in the definition of the function object. Its signature is:

```haskell
eval :: ((a -> b), a) -> b
```

It takes a pair consisting of a function and its argument and produces a result of the appropriate type. It’s the Haskell implementation of the morphism:

```haskell
eval :: (a⇒b) × a -> b
```

which defines the function type a⇒b (or the exponential object ba). Let’s translate this signature to a logical predicate using the Curry-Howard isomorphism:

```haskell
((a ⇒ b) ∧ a) ⇒ b
```

Here’s how you can read this statement: If it’s true that b follows from a, and a is true, then b must be true. This makes perfect intuitive sense and has been known since antiquity as modus ponens. We can prove this theorem by implementing the function:

```haskell
eval :: ((a -> b), a) -> b
eval (f, x) = f x
```
