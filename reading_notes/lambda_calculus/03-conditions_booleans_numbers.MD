# Chapter 3. Conditions, booleans and numbers

## 3.2. Truth values and conditional expression
```
    def cond  = λe1.λe2.λc.((c e1) e2)
    
    select_first  = λfirst.λsecond. first
    select_second = λfirst.λsecond. second

    def true  = select_first
    def false = select_second   
```

## 3.3. NOT
```
    def not = λx.((x false) true)
```

## 3.4. AND
```
    def and = λx.λy.((x y) false)
```

## 3.5. OR
```
    def or = λx.λy.((x true) y)
```

## 3.6. Natural numbers
```
    def zero = identity
    def succ = λn.λs.((s false) n)

    def iszero = λn.(n select_first)

    def pred1 = λn.(n select_second)
    def pred  = λn.(((iszero n) zero) (n select_second))
```

## 3.7. Simplified notations
```
    def identity x = x
    def self_apply s = s s
    def apply func = λarg.(func arg)
```
rewrite:
```
    def pred n =
        if iszero n
        then zero
        else n select_second

    def not x =
        if x
        then false
        else true

    def and x y =
        if x
        then y
        else false

    def or x y =
        if x
        then true
        else y
```