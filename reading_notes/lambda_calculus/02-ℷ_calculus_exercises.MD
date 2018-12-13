# Chapter 2. λ calculus

## Exercises

#### 2.4 Define a function:

make_triple:
```
def make_triple = λfunc.λfirst.λsecond.λthird(((func first) second) third)
```

so:
```
def triple_first  = (make_triple λx.y.z.x)
def triple_second = (make_triple λx.y.z.y)
def triple_third  = (make_triple λx.y.z.z)
```
