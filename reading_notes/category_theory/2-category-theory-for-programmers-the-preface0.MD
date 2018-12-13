# category-theory-for-programmers-the-preface

## 8. Natural Transformations

There may be many ways of embedding one category in another. Sometimes they are equivalent, sometimes very different. One may collapse the whole source category into one object, another may map every object to a different object and every morphism to a different morphism. The same blueprint may be realized in many different ways. Natural transformations help us compare these realizations. They are mappings of functors — special mappings that preserve their functorial nature.

Consider two functors F and G between categories C and D. If you focus on just one object a in C, it is mapped to two objects: F a and G a. A mapping of functors should therefore map F a to G a.

### 8.1 Polymorphic Functions

We talked about the role of functors (or, more specifically, endofunctors) in programming. They correspond to type constructors that map types to types. They also map functions to functions, and this mapping is implemented by a higher order function ```fmap``` (or ```transform```, ```then```, and the like in C++).

To construct a natural transformation we start with an object, here a type, a. One functor, F, maps it to the type F a. Another functor, G, maps it to G a. The component of a natural transformation alpha at a is a function from F a to G a. In pseudo-Haskell:

```haskell
alphaª :: F a -> G a
```

Keep in mind that it’s really a family of functions parameterized by a. This is another example of the terseness of the Haskell syntax. A similar construct in C++ would be slightly more verbose:

```haskell
template<class A> G<A> alpha(F<A>);
```

There is a more profound difference between Haskell’s polymorphic functions and C++ generic functions, and it’s reflected in the way these functions are implemented and type-checked. In Haskell, a polymorphic function must be defined uniformly for all types. One formula must work across all types. This is called parametric polymorphism.

C++, on the other hand, supports by default ad hoc polymorphism, which means that a template doesn’t have to be well-defined for all types. Whether a template will work for a given type is decided at instantiation time, where a concrete type is substituted for the type parameter. Type checking is deferred, which unfortunately often leads to incomprehensible error messages.

In C++, there is also a mechanism for function overloading and template specialization, which allows different definitions of the same function for different types. In Haskell this functionality is provided by type classes and type families.

Haskell’s parametric polymorphism has an unexpected consequence: any polymorphic function of the type:

```haskell
alpha :: F a -> G a
```

where F and G are functors, automatically satisfies the naturality condition. Here it is in categorical notation (f is a function f::a->b):

```haskell
G f ∘ αa = αb ∘ F f
```

The naturality condition is a pretty stringent requirement. For instance, if the morphism F f is invertible, naturality determines αb in terms of αa. It transports αa along f:

```haskell
αb = (G f) ∘ αa ∘ (F f)-1
```

If there is more than one invertible morphism between two objects, all these transports have to agree. In general, though, morphisms are not invertible; but you can see that the existence of natural transformations between two functors is far from guaranteed. So the scarcity or the abundance of functors that are related by natural transformations may tell you a lot about the structure of categories between which they operate. We’ll see some examples of that when we talk about limits and the Yoneda lemma.

Looking at a natural transformation component-wise, one may say that it maps objects to morphisms. Because of the naturality condition, one may also say that it maps morphisms to commuting squares — there is one commuting naturality square in D for every morphism in C.
