# Chapter 5. TYPES

## 5.4. Representing typed objects

```
    def make_obj type value = λs.(s type value)

    def type obj = obj select_first

    def value obj = obj select_second

    def istype t obj = equal (type obj) t
```

## 5.5. Errors

```
    def error_type = zero

    def MAKE_ERROR = make_obj error_type

    def ERROR = MAKE_ERROR error_type

    def iserror = istype error_type

    def ISERROR E = MAKE_BOOL (iserror E)
```

## 5.6. Booleans

```
    def bool_type = one

    def MAKE_BOOL = make_obj bool_type

    def TRUE = MAKE_BOOL true

    def FALSE = MAKE_BOOL false

    def isbool = istype bool_type

    def ISBOOL B = MAKE_BOOL (isbool B)

    def BOOL_ERROR = MAKE_ERROR bool_type

    def NOT X =
        if isbool X
        then MAKE_BOOL (not (value X))
        else BOOL_ERROR

    def AND X Y =
        if and (isbool X) (isbool Y)
        then MAKE_BOOL (and (value X) (value Y))
        else BOOL_ERROR
```

## 5.7. Typed conditional expression

```
    def COND E1 E2 C =
        if isbool C
        then
            if value C
            then E1
            else E2
        else BOOL_ERROR
```

## 5.8. Numbers and arithmetic

```
    def numb_type = two

    def MAKE_NUMB = make_obj numb_type

    def NUMB_ERROR = MAKE_ERROR numb_type

    def isnumb = istype numb_type

    def ISNUMB N = MAKE_BOOL (isnumb N)

    def 0 = MAKE_NUMB zero

    def SUCC N =
        if isnumb N
        then MAKE_NUMB (succ (value N))
        else NUMB_ERROR

    def PRED N =
        if isnumb N
        then
            if iszero (value N)
            then NUMB_ERROR
            else MAKE_NUMB ((value N) select_second)
        else NUMB_ERROR

    def ISZERO N =
        if isnumb N
        then MAKE_BOOL (iszero (value N))
        else NUMB_ERROR

    def both_numbs X Y = and (isnumb X) (isnumb Y)

    def + X Y =
        if both_numbs X Y
        then MAKE_NUMB (add (value X) (value Y))
        else NUMB_ERROR

    def * X Y =
        if both_numbs X Y
        then MAKE_NUMB (mult (value X) (value Y))
        else NUMB_ERROR

    def / X Y =
        if both_numbs X Y
        then
            if iszero (value Y)
            then NUMB_ERROR
            else MAKE_NUMB (div1 (value X) (value Y))
        else NUMB_ERROR

    def EQUAL X Y =
        if both_numbs X Y
        then MAKE_BOOL (equal (value X) (value Y))
        else NUMB_ERROR
```

## 5.9. Characters

```
    def char_type = four

    def CHAR_ERROR = MAKE_ERROR char_type
    
    def ischar = istype char_type

    def ISCHAR C = MAKE_BOOL (ischar C)

    def MAKE_CHAR = make_obj char_type

    def ’0’ = MAKE_CHAR forty_eight
    def ’1’ = MAKE_CHAR (succ (value ’0’))
    ...
    def ’z’ = MAKE_CHAR (succ (value ’y’))    

    def CHAR_LESS C1 C2 =
        if and (ischar C1) (ischar C2)
        then MAKE_BOOL (less (value C1) (value C2))
        else CHAR_ERROR

    def ORD C =
        if ischar C
        then MAKE_NUMB (value C)
        else CHAR_ERROR

    def CHAR N =
        if isnumb N
        then MAKE_CHAR (value N)
        else NUMB_ERROR

    def CHAR_EQUAL C1 C2 =
        if and (ischar C1) (ischar C2)
        then MAKE_BOOL (equal (value C1) (value C2))
        else CHAR_ERROR
```

## 5.10. Repetitive type checking

```
    def NOT X =
        IF X
        THEN FALSE
        ELSE TRUE

    def AND X Y =
        IF ISBOOL Y
        THEN
            IF X
            THEN Y
            ELSE FALSE
        ELSE BOOL_ERROR
```

## 5.13. Case definitions and structure matching

```
    λ<names1>.<expression1>
    or <names2>.<expression2>
    or ...
```