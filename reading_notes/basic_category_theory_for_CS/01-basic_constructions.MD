# 1 Basic Constructions

## 1.1 Categories

### 1.1.1 Definition 

A category C comprises:

1. a collection of **objects**;
2. a collection of **arrows**(ofen called **morphisms**);
3. operations assigning to each arrow _f_ an object _dom f_, its **domain(定义域)**, and an object _cod f_, its **codomain(值域)**(we write _f: A → B_ or _A f⥱ B_ to show that _dom f = A_ and _cod f = B_; the collection of all arrows with domain *A* and codomain _B_ is written C(_A,B_));
4. a composition operator assigning to each pair of arrows _f_ and _g_, with _cod f = dom g_, a **composite** arrow _g ∘ f: dom f → cod g_, satisfying the following **_associative law_**:
```
    for any arrow f: A → B, g: B → C, and h: C → D(with A,B,C and D not necessarily distinct),
        h∘(g∘f) = (h∘g)∘f
```
5. for each object _A_, an **identity** arrow _id𝑨: A → A_ satisfying the following **_identity law_**:
```
    for any arrow f: A → B,
          idʙ∘f = f   and    f∘idʙ = f
```

### 1.1.2 Remark

#### 1.1.2.1 Total Function: 
A function which is defined for all inputs of the right type, that is, for all of a domain. See also partial function. 

#### 1.1.2.2 Partial function:
a partial function from X to Y (sometimes written as f: X ↛ Y or f: X ⇸ Y) is a function f: X′ → Y, for some proper subset X′ of X. It generalizes the concept of a function f: X → Y by not forcing f to map every element of X to an element of Y (only some proper subset X′ of X).

#### 1.1.2.3 Note: 
Square (x²) is a total function. Reciprocal (1/x) is not, since 0 is a real number, but has no reciprocal.

### 1.1.5 一些概念

#### 1.1.5.1 reflexive 自反关系
对于集合X上的二元关系R，若满足： ```∀a ∈ X(aRa)```，则称二元关系R是自反的，或称R具有自反性，或称R为自反关系。

#### 1.1.5.2 transitive 传递关系
若对所有的a，b，c属于X，下述语句保持有效，则集合X上的二元关系R是传递的：“若a关系到b且b关系到c，则 a关系到c。”
```∀a,b,c ∈ X, aRb ∧ bRc ⇒ aRc```

#### 1.1.5.3 symmetric 对称关系
若对所有的 a 和 b 属于 X，下述语句保持有效，则集合 X 上的二元关系 R 是对称的：“若 a 关系到 b，则 b 关系到 a。”
```∀a,b ∈ X, aRb ⇒ bRa```

#### 1.1.5.4 anti-symmetric 反对称关系
若对所有的 a 和 b 属于 X，下述语句保持有效，则集合 X 上的二元关系 R 是反对称的：“若 a 关系到 b 且 b 关系到 a，则 a = b。”
```∀a,b ∈ X, aRb ∧ bRa ⇒ a=b```

#### 1.1.5.5    组合
```
reflexive       + transitive    = preorder 预序关系
anti-symmetric  + preorder      = poset (partially ordered set) 偏序关系
symmetric       + preorder      = equivalence relation 等价关系
```

### 1.1.6 Monoid
A monoid _(M,・,e)_ is an underlying set M equipped with a binary operation・from pairs of elements of M into M such that _(x・y)・z = x・(y・z)_ for all _x,y,z ∈ M_ and a distinguished element _e_ such that _e・x = x = x・e_ for all _x ∈ M_.

A monoid homomorphism from _(M,・,e)_ to _(M',・',e')_ is a function _f: M → M'_ such that _f(e) = e'_ and _f(x・y) = f(x)・'f(y)_.

So, a **_monoid_** satisfies the following two axioms:
1. **Associativity**: For all a, b and c in S, the equation (a • b) • c = a • (b • c) holds.
2. **Identity element**: There exists an element e in S such that for every element a in S, the equations e • a = a • e = a hold.

#### 1.1.6.1 Homomorphism
In algebra, a **_homomorphism_** is a structure-preserving map between two algebraic structures of the same type (such as two groups, two rings, or two vector spaces).

codomain(值域)中被函数箭头打到的元素称为该函数的像（image）。符号 _Im(φ)_ 表示同态 _φ_ 的像。

### 1.1.7 Ω-algebra
given a signature Ω, an Ω-algebra is comprised of a "carrier" for the algebra and an "interpretation" for every operation symbol. 
The carrier of _A_ is a set (written |A|).
The interpretation of _ω ∈ Ω_ is a function _aω:|A|ar(ω)⟼|A|_, where _ar_ is the function associating to each operation symbol its arity.

#### 1.1.7.1 Ω-homomorphism
An Ω-homomorphism from Ω-algebra _A_ to an Ω-algebra _B_ is a function _h: |A|→|B|_ such that for each operator _ω ∈ Ω_ and tuple _X1,X2,...,Xar(ω)_ of elements of _|A|_, the following equation holds:
```
    h(aω(X1,X2,...,Xar(ω))) = bω(h(X1),h(X2),...,h(Xar(ω)))
```
The category **Ω-Alg** has Ω-algebras as objects and  Ω-homomorphisms as arrows.

#### 1.1.7.2 Disjoint Union
The disjoint union of two sets A and B is a binary operator that combines all distinct elements of a pair of given sets, while retaining the original set membership as a distinguishing characteristic of the union set. The disjoint union is denoted
```
    A ∪* B = (A×{0})∪(B×{1}) ≡ A* ∪ B*
```
where A×S is a Cartesian product. For example, the disjoint union of sets _A={1,2,3,4,5}_ and _B={1,2,3,4}_ can be computed by finding
```
    A* = {(1,0),(2,0),(3,0),(4,0),(5,0)}
    B* = {(1,1),(2,1),(3,1),(4,1)}
```
so
```
    A ∪* B = A* ∪ B*
        = {(1,0),(2,0),(3,0),(4,0),(5,0),(1,1),(2,1),(3,1),(4,1)}
```

### 1.1.14
By a twist of perspective, we can call the objects in an arbitrary category formulas and the arrows proofs.

An arrow _f: A → B_ is viewed as a proof of the logical implication _A ⊃ B_.

### 1.1.16
For each category C, the objects of the **dual category C°P** are the same as those of C; the arrows in C°P are the opposites of the arrows in C. That is, if _f : A - B_ is a arrow of C then _f : B - A_ is a arrow of C°P.

### 1.1.17
For any pair of categories Cand D, the **product category** _CxD_ has as objects pairs _(A, B)_ of a C-object A and a D-object B and as arrows pairs (f, g) of a C-arrow f and a D-arrow g. Composition and identity arrows are defined pairwise: (f, g) o (h, i) = (f o h, g o i) and id(A,B) = (idA, idB).


## 1.2 Diagrams

### 1.2.1 Definition

A **diagram** in a category C is a collection of vertices and directed edges, consistently labeled with objects and arrows of C, where "consistently" means that if an edge in the diagram is labeled with an arrow _f_ and _f_ has domain _A_ and codomain _B_, then the endpoints of the edge must be labeled with _A_ and _B_.

### 1.2.2 Definition

A diagram in a category C is said to **commute** if, for every pair of vertices _X_ and _Y_, all the paths in the diagram from _X_ to _Y_ are equal, in the sense that each path in the diagram determines an arrow and these arros are equal in C.

### 1.2.4 Proposition
If both inner squares of the following diagram commute, then so does the outer rectangle.

Proof:
    (g' o g) o a    =   g' o (g o a)    (associativity)
                    =   g' o (b o f)    (commutativity of first square)
                    =   (g o b) o f     (associativity)
                    =   (c o f') o f    (commutativity of second square)
                    =   c o (f' o f)    (associativity).
                                                                        (End of Proof)

## 1.3 Monomorphisms, Epimorphisms, and Isomorphisms

### 1.3.1 Definition 
An arrow _f : B → C_ in a category C is a **monomorphism**(or "is **monic**") if, for any pair of C-arrows _g : A → B_ and _h: A → B_, the equality _f o g = f o h_ implies that _g = h_.

### 1.3.4 Definition 
An arrow _f : A → B_ is an **epimorphism** (or "is **epic**") if, for any pair of arrows _g : B → C_ and _h : B → C_ , the equality _g o f = h o f_ implies that _g = h_.

### 1.3.7 Definition
An arrow _f : A → B_ is an **isomorphism** if there is an arrow _f⁻¹ : B → A_, called the inverse of f, such that _f⁻¹ o f = idᴬ_ and _f o f⁻¹ = idᴮ_. The objects _A_ and _B_ are said to be **isomorphic** if there is an isomorphism between them.

## 1.4 Initial and Terminal Objects

### 1.4.1 Definition 
An object _O_ is called an **initial object** if, for every object _A_, there is exactly one arrow from _O_ to _A_.

### 1.4.2 Definition
Dually, an object _1_ is called a **terminal** or **final object** if, for every object A, there is exactly one arrow from _A_ to _1_.

### 1.4.5 Example
In categorical terms, an arrow from a terminal object to an object S is called a **global element** or **constant** of S.


## 1.5 Products

The usual set-theoretic definition of the cartesian product of two sets A and B is:
```   A x B={(a, b)| a ∈ A and b ∈ B}.```

### 1.5.1 Definition
A **product** of two objects A and B is an object A x B, together with two **projection arrows** _π1 : AxB → A_ and _π2 : AxB → B_, such that for any object C and pair of arrows _f : C → A_ and _g: C → B_ there is exactly one **mediating arrow** _(f, g) : C → A x B_ making the diagram commute —— that is, such that _π1 o (f, g) = f_ and _π2 o (f, g) = g_.

### 1.5.3 Definition 
If _A x C_ and _B x D_ are product objects, then for every pair of arrows _f : A → B_ and _g : C → D_, the **product map** _f x g : A x C → B x D_ is the arrow _(f o π1, g o π2)_.

### 1.5.4 Definition
A **coproduct** of two objects A and B is an object _A + B_, together with two injection arrows _ι1 : A → A + B_ and _ι2 : B → A + B_, such that for any object C and pair of arrows _f: A → C_ and _g: B → C_ there is exactly one arrow _(f, g) : A + B → C_.

## 1.6 Universal Constructions
A **universal construction** describes a class of objects and accompanying arrows that share a common property and picks out the objects that are terminal when this class is considered as a category.

## 1.7 Equalizers

### 1.7.1 Definition
An arrow _e: X → A_ is an **equalizer** of a pair of arrows _f: A → B_ and _g: A → B_ if
1. f o e = g o e;
2. whenever _e' : X' → A_ satisfies _f o e' = g o e'_, there is a unique arrow _k : X' → X_ such that _e o k = e'_.

## 1.8 Pullbacks

### 1.8.1 Definition 
A **pullback** of the pair of arrows _f: A → C_ and _g: B → C_ is an object P and a pair of arrows _g': P → A_ and _f': P → B_ such that _f o g' = g o f'_:

and if _i: X → A_ and _j: X → B_ are such that _f o j = g o j_:

then there is a unique _k : X → P_ such that _i = g' o k_ and _j = f' o k_:

This situation is commonly described by saying that _f'_ is a pullback (or inverse image) of _f_ along _g_ and that _g'_ is a pullback of _g_ along _f_. The example that motivates this terminology comes, as usual, from **Set**:


## 1.9 Limits

## 1.9.1 Definition 
Let C be a category and D a diagram in C. A **cone** for D is a C-object X and arrows _f𝔦 : X → D𝔦_ (one for each object D𝔦 in D), such that for each arrow g in D, the diagram commutes. We use the notation _{f : X → D)_ for cones.

## 1.9.2 Definition 
A **limit** for a diagram D is a cone _{f𝔦, : X' → D𝔦}_ with the property that if _{f𝔦' : X' → D𝔦}_ is another cone for D then there is a unique arrow _k: X' → X_ such that the diagram commutes for every _D𝔦_ in D.

## 1.9.7 Theorem (Limit Theorem) 
Let D be a diagram in a category C, with sets V of vertices and E of edges. If every V-indexed and every E-indexed family of objects in C has a product and every pair of arrows in C has an equalizer, then D has a limit.


## 1.10 Exponentiation

### 1.10.1 Definition 
Let C be a category with all binary products and let _A_ and _B_ be objects of C. An object _Bᴬ_ is an **exponential object** if there is an arrow _evalᴬᴮ : (Bᴬ x A) → B_ such that for any object C and arrow _g: (C x A) → B_ there is a unique arrow _curry(g): C → Bᴬ_ making commute——that is, a unique _curry(g)_ such that
```
    evalᴬᴮ o (curry(g) x idᴬ) = g
```
If C has an exponential _Bᴬ_ for every pair of objects _A_ and _B_, then C is said to **have exponentiation**.

### 1.10.2 Definition 
A **cartesian closed category (CCC)** is a category with a terminal object, binary products, and exponentiation.