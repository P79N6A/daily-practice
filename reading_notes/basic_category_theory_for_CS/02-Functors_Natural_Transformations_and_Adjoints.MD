# 2 Functors, Natural Transformations, and Adjoints

## 2.1 Functors

### 2.1.1 Definition 
Let C and D be categories. A **functor** _F: C → D_ is a map taking each C-object A to a D-object F(A) and each C-arrow _f: A → B_ to a D-arrow _F(f): F(A) → F(B)_, such that for all C-objects A and composable C-arrows f and g 
1. F(idᴬ) = id𝘍(ᴬ) 
2. F(g o f) = F(g) o F(f).

### 2.1.5 forgetful functor 
The **forgetful functor** _U: Mon → Set_ sends each monoid _(M,・, e)_ to the set M and each monoid homomorphism _h: (M,・, e) → (M',・', e')_ to the corresponding function _h: M - M'_ on the underlying sets.

### 2.1.8 The category Cat 
The category **Cat** has categories as objects and functors as arrows.

### 2.1.11 homfunctor
Given a category C, each C-object A determines a functor _C(A,－): C → Set_. This functor takes each C-object B to the set C(A, B) of arrows from A to B and each C-arrow _f: B → C_ to the function _C(A, f): C(A, B) → C(A, C)_ defined by
```
    C(A, f)(g: A → B) = fog
```
where the composition on the right is in _C. C(A,－)_ is called a **homfunctor**. (The origin of the term is in the frequent use of arrows in categories to model homomorphisms of various sorts. The set C(A, B) is often called a horn-set.)


## 2.2 F-Algebras

### 2.2.2 Definition
Let K be a category and _F: K → K_ a functor. An _F-algebra_ is a pair _(A, a)_, where A is a K-object and _a: F(A) → A_ is a K-arrow. An F-homomorphism _h: (A, a) → (B, b)_ is a K-arrow _h: A → B_ such that the followmg diagram commutes.