# Chapter 2. λ calculus

## 2.7 β reduction

Formally, the replacement of a bound variable with an argument in a function body is called **beta reduction(β reduction)**.
In future, instead of detailing each normal order β reduction we will introduce the notation:
```
    (<function> <argument>) => <expression>
```
to mean that the <expression> results from the application of the <function> to the unevaluated <argument>.
When we have seen a sequence of reductions before, or we are familiar with the functions involved, we will omit the reductions and write:
```
    => ... =>
```
to show where they should be.

## 2.10 free and bound variables

#### (1) bound:

1. The expression is an application:
```
    (<function> <argument>)
```
and the variable is bound in <function> or <argument>. For example, **convict** is bound in
```
    (λconvict.convict fugitive)
```
and in:
```
    (λprison.prison λconvict.convict)
```
2. The expression is a function:

#### (2) free:

1. The expression is a single name:
```
    <name>
```
and the variable's name is <name>. For example, **truant** is free in:
```
    <truant>
```
2. The expression is an application:
```
    (<function> <argument>)
```
and the variable is free in <function> or in <argument>. For example, **escaper** is free in:
```
    (λprisoner.prisoner escaper)
```
and in:
```
    (escaper λjailor.jailor)
```
3. The expression is a function:
```
    λ<name>.<body>
```
and the variable's name is not <name> and the variable is free in <body>. For example, **fugitive** is free in:
```
    λprison.(prison fugitive)
```
and in:
```
    λshort.λsharp.λshock.fugitive
```

## 2.11 name clashes and α conversion

#### (1) name clash:
A name clash arises when a β reduction places an expression with a free variable in the scope of a bound variable with the same name as the free variable.

#### (2) α conversion:
Consistent renaming, which is known as **alpha conversion(α conversion)**, remove the name clash.

For a function:
```
    λ<name1>.<body>
```
the name <name1> and all free occurrences of <name1> in <body> may be replaced by a new name <name2> provided <name2> is not the name of a free variable in λ<name1>.<body>. Note that replacement include the name at:
```
    λ<name1>
```

## 2.12 Simplification through η reduction

Consider an expression of the form:
```
    λ<name>.(<expression> <name>)
```
This is similar to the function application function above, after application to a function expression only. This is equivalent to:
```
    <expression>
```
because the application of this expression to an arbitrary argument:
```
    <argument>
```
gives:
```
    (λ<name>.(<expression> <name>) <argument>) =>
    (<expression> <argument>)
```
This Simplification of:
```
    λ<name>.(<expression> <name>)
```
to:
```
    <expression>
```
is called **eta reduction(η reduction)**.
