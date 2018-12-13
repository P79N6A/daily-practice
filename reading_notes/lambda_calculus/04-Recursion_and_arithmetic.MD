# Chapter 4. Recursion and arithmetic

```
    def add x y =
        if iszero y
        then x
        else add (succ x) (pred y)

    def mult x y =
        if iszero y
        then zero
        else add x (mult x (pred y))
```

## 4.5 Recursion function
```
    def recursive f = λs.(f (s s)) λs.(f (s s))
```

## 4.7. Recursion notation
```
    rec <name> = <expression>
```

## 4.8. Arithemtic operations

### 4.8.1. Power
```
    rec power x y =
        if iszero y
        then one
        else mult x (power x (pred y))
```

### 4.8.2. Subtraction
```
    rec sub x y =
        if iszero y
        then x
        else sub (pred x) (pred y)
```

### 4.8.3. Comparison
```
    def abs_diff x y = add (sub x y) (sub y x)
    
    def equal x y = iszero (abs_diff x y)

    def greater x y = not (iszero (sub x y))

    def greater_or_equal x y = iszero (sub y x)

    def less x y = greater y x

    def less_or_equal x y = greater_or_equal y x
```

### 4.8.4. Division
```
    rec div1 x y =
        if greater y x
        then zero
        else succ (div1 (sub x y) y)

    def div x y =
        if iszero y
        then zero
        else div1 x y
```