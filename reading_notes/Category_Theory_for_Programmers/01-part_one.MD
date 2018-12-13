# Types and Functions

## 2. Pure and Dirty Functions
In programming languages, functions that always produce the same result given the same input and have no side effects are called **pure functions**.

Functions to Bool are called **predicates**.


## 3. Categories Great and Small

Consider some set P and a binary relation ≤ on P. Then ≤ is a **preorder**:
* a ≤ a (reflexivity)
* if a ≤ b and b ≤ c then a ≤ c (transitivity)

If a preorder is also antisymmetric, that is, a ≤ b and b ≤ a implies a = b, then it is a **partial order**.(反对称性)

Any two objects are in a relation with each other, one way or another; and that gives you a **linear order** or **total order**.

A set of morphisms from object a to object b in a category C is called a **hom-set** and is written as C(a, b) (or, sometimes, **HomC(a, b)**).


## 4. Kleisli Categories

### repr
```haskell
a -> m b
```

### Writer in Haskell

```haskell
type Writer a = (a, String)
```


## 5. Products and Coproducts

The **initial object** is the object that has one and only one morphism going to any object in the category.

The **terminal object** is the object with one and only one morphism coming to it from any object in the category.

An **isomorphism** is an invertible morphism; or a pair of morphisms, one being the inverse of the other.

A **product** of two objects a and b is the object c equipped with two projections such that for any other object c’ equipped with two **projections** there is a unique morphism m from c’ to c that factorizes those projections.

A **coproduct** of two objects a and b is the object c equipped with two injections such that for any other object c’ equipped with two **injections** there is a unique morphism m from c to c’ that factorizes those injections.

### surjective, injective and bijections
1. functions that tightly fill their codomains are called **surjective** or onto (满射).
2. non-collapsing functions: they call them **injective** or one-to-one (单射).
3. some functions that are neither embedding nor collapsing. They are called **bijections** and they are truly symmetric (双射).


## 6. Simple Algebraic Data Types

### Algebra of Types

the **sum types** with ```Void``` as the neutral element (coproduct), 
and the **product types** with the unit type, ```()```, as the neutral element (product). 


## 7. Functors

A **functor** is a mapping between categories.
* translation table with some entries of interest:

| Numbers  | Types                                    |
|----------|------------------------------------------|
| 0        | ```Void```                               |
| 1        | ```( )```                                |
| a + b    | ```Either a b = Left a \| Right b```     |
| a * b    | ```(a, b)``` or ```Pair a b = Pair a b```|
| 2 = 1 + 1| ```data Bool = True \| False```          |
| 1 + a    | ```data Maybe = Nothing \| Just a```     |


* Logical _and_ and _or_ also form a semiring, and it too can be mapped into type theory:

| Logic    | Types                                    |
|----------|------------------------------------------|
| false    | ```Void```                               |
| true     | ```( )```                                |
| a || b   | ```Either a b = Left a | Right b```      |
| a && b   | ```(a, b)```                             |


We can talk about functors that map this category into itself — such functors are called **endofunctors**. 

### Typeclasses

A **typeclass** defines a family of types that support a common interface. 

### The Reader Functor

```haskell
instance Functor ((->) r) where
    fmap f g = f . g

⇒

fmap = (.)
```
This combination of the type constructor ```(->) r``` with the above implementation of ```fmap``` is called the reader functor.



## 8. Functoriality

### Bifunctors

A **bifunctor** maps every pair of objects, one from category C, and one from category D, to an object in category E.

The definition of the Bifunctor typeclass:
```haskell
class Bifunctor f where
    bimap :: (a -> c) -> (b -> d) -> f a b -> f c d
    bimap g h = first g . second h
    first :: (a -> c) -> f a b -> f c b
    first g = bimap g id
    second :: (b -> d) -> f a b -> f a d
    second = bimap id
```

### Product and Coproduct Bifunctors

Here’s the **Bifunctor** instance for a pair constructor — the simplest product type:
```haskell
instance Bifunctor (,) where
    bimap f g (x, y) = (f x, g y)
```

### The Writer Functor

```haskell
type Writer a = (a, String)
```

Let’ me remind you that the composition is given by the fish operator:
```haskell
(>=>) :: (a -> Writer b) -> (b -> Writer c) -> (a -> Writer c)
m1 >=> m2 = \x -> 
    let (y, s1) = m1 x
        (z, s2) = m2 y
    in (z, s1 ++ s2)
```

and the identity morphism by a function called return:
```haskell
return :: a -> Writer a
return x = (x, "")
```

combine them to produce a function with the right type signature to serve as **fmap**. Like this:
```haskell
fmap f = id >=> (\x -> return (f x))
```

You might wonder if the **fmap** we have just defined is the same **fmap** the compiler would have derived for us with **deriving Functor**. Interestingly enough, it is. This is due to the way Haskell implements polymorphic functions. It’s called _parametric polymorphism_, and it’s a source of so called theorems for free. One of those theorems says that, if there is an implementation of **fmap** for a given type constructor, one that preserves identity, then it must be unique.

### Covariant and Contravariant Functors

Here’s the typeclass defining a **contravariant functor** (really, a contravariant endofunctor) in Haskell:
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

There is a special function for flipping arguments, called flip:
```haskell
flip :: (a -> b -> c) -> (b -> a -> c)
flip f y x = f x y
```

With it, we get:
```haskell
contramap = flip (.)
```

### Profunctors

A profunctor is defined as:
```
Cop × D → Set
```

Here’s the appropriate typeclass taken from the Data.Profunctor library:
```haskell
class Profunctor p where
  dimap :: (a -> b) -> (c -> d) -> p b c -> p a d
  dimap f g = lmap f . rmap g
  lmap :: (a -> b) -> p b c -> p a c
  lmap f = dimap f id
  rmap :: (b -> c) -> p a b -> p a c
  rmap = dimap id
```

Now we can assert that the function-arrow operator is an instance of a **Profunctor**:
```haskell
instance Profunctor (->) where
  dimap ab cd bc = cd . bc . ab
  lmap = flip (.)
  rmap = (.)
```

### The Hom-Functor

A morphism in Cop×C is a pair of morphisms from C:
```haskell
f :: a'→ a
g :: b → b'
```

The lifting of this pair must be a morphism (a function) from the set C(a, b) to the set C(a', b'). Just pick any element h of C(a, b) (it’s a morphism from a to b) and assign to it:
```haskell
g ∘ h ∘ f
```
which is an element of ```C(a', b')```.

As you can see, the **hom-functor** is a special case of a profunctor.



## 9. Function Types

A set of morphisms between two objects in any category is called a **hom-set**.

### Universal Construction

A **function object** from ```a``` to ```b``` is an object ```a⇒b``` together with the morphism
```haskell
eval :: ((a⇒b) × a) -> b
```
such that for any other object ```z``` with a morphism
```haskell
g :: z × a -> b
```
there is a unique morphism
```haskell
h :: z -> (a⇒b)
```
that factors ```g``` through eval:
```haskell
g = eval ∘ (h × id)
```

### Currying

g as a function of two variables, z and a.
```haskell
g :: z × a -> b
```
h that maps z to a function object ```a⇒b```.
```haskell
h :: z -> (a⇒b)
```
In **Set**, this just means that h is a function that takes one variable of type z and returns a function from a to b. That makes h a higher order function. Therefore the universal construction establishes a one-to-one correspondence between functions of two variables and functions of one variable returning functions. This correspondence is called _currying_, and h is called the curried version of g.

### Exponentials
In mathematical literature, the function object, or the internal hom-object between two objects a and b, is often called the exponential and denoted by ```bª```.

### Cartesian Closed Categories
A cartesian closed category must contain:

1. The terminal object,
2. A product of any pair of objects, and
3. An exponential for any pair of objects.

### Exponentials and Algebraic Data Types

#### Zeroth Power
```
a^0 = 1
```