# Chapter 9. Functional Programming In Standard ML

## 9.2. Types

When representing objects, SML always displays types along with values:
```
    <value> : <type>
```
In particular, function values are displayed as:
```
    fn : <type>
```

## 9.13. Functions
```
    fn <bound variables> => <expression>
```

## 9.15. Definitions
```
    val <name> = <expression>
```

## 9.16. Conditional expressions
```
    if <expression1>
    then <expression2>
    else <expression3>
```

## 9.17. Recursion and function definitions
```
    fun <name> <bound variables> = <expression> ==
    val rec <name> = fn <bound variables> => <expression>
```

## 9.23. New types
A new **concrete** type may be introduced by a **datatype binding**.

```
    datatype <constructor> = <constructor1> |
                             <constructor2> |
                             ...
                             <constructorN>
```

## 9.25. λ calculus in ML

```sml
    id:
    fn x => x;
    
    apply:
    fn f => fn x => (f x);
```