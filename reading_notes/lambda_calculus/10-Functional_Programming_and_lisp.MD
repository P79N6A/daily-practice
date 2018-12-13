# Chapter 10. Functional Programming And Lisp

The form is the basic LISP construct and consists of an atom or a left bracket followed by zero or more atoms or forms ending with a right bracket:

```lisp
    <form>  ::= <atom> | ( <forms> ) | ()
    <forms> ::= <form> | <form> <forms>
```

## 10.7. Global definitions
```lisp
    (defun <name> (<bound variable list>) <body>)
```

## 10.8. Conditional expressions
```lisp
    (cond   (<test1> <result1>)
            (<test2> <result2>)
            ...
            (t <resultN>))
```

## 10.9. Quoting

In LISP, a mechanism known as **quoting** is used to delay form evaluation.
```lisp
    (quote <argument>) -> <argument>
```
special notation:
```lisp
    ’<argument>
```

## 10.10. Lists

The empty list is the primitive:
```lisp
    nil
```
Lists may be constructed explicitly with the:
```lisp
(cons <head> <tail>)
```

## 10.11. List selection
head:
```lisp
    car
```
tail:
```lisp
    cdr
```

## 10.16. Functions as values and arguments

It creates a function value in which free variables are associated with bound variables in the defining scope. As with quote, an equivalent special notation:
```lisp
    #’<function>
```