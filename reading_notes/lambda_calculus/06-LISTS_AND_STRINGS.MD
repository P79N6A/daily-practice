# Chapter 6. LISTS AND STRINGS

## 6.2. List representation

```
    def list_type = three

    def islist = istype list_type

    def ISLIST L = MAKE_BOOL (islist L)

    def LIST_ERROR = MAKE_ERROR list_type

    def MAKE_LIST = make_obj listtype

    def CONS H T =
        if islist T
        then MAKE_LIST λs.(s H T)
        else LIST_ERROR

    def NIL = MAKE_LIST λs.(s LIST_ERROR LIST_ERROR)

    def HEAD L =
        if islist L
        then (value L) select_first
        else LIST_ERROR

    def TAIL L =
        if islist L
        then (value L) select_second
        else LIST_ERROR

    def isnil L =
        if islist L
        then iserror (HEAD L)
        else false

    def ISNIL L =
        if islist L
        then MAKE_BOOL (iserror (HEAD L))
        else LIST_ERROR
```

## 6.4. Linear length of a list

```
    rec LENGTH L =
        IF ISNIL L
        THEN 0
        ELSE SUCC (LENGTH (TAIL L))
```

## 6.5. Appending lists

```
    rec APPEND L1 L2 =
        IF ISNIL L1
        THEN L2
        ELSE CONS (HEAD L1) (APPEND (TAIL L1) L2)
```

## 6.6. List notation

```
    <expression1>::<expression2> == CONS <expression1> <expression2>

    [<first item>,<second item>] == CONS <first item> (CONS <second item> NIL)
```

## 6.8. Deletion from a list

```
    rec DELETE V L =
        IF ISNIL L
        THEN NIL
        ELSE
            IF EQUAL V (HEAD L)
            THEN TAIL L
            ELSE (HEAD L)::(DELETE V (TAIL L))
```

## 6.9. List comparison

```
    rec LIST_EQUAL L1 L2 =
        IF AND (ISNIL L1) (ISNIL L2)
        THEN TRUE
        ELSE
            IF OR (ISNIL L1) (ISNIL L2)
            THEN FALSE
            ELSE
                IF EQUAL (HEAD L1) (HEAD L2)
                THEN LIST_EQUAL (TAIL L1) (TAIL L2)
                ELSE FALSE
```

## 6.10. Strings

```
    rec ISSTRING S =
        IF ISNIL S
        THEN TRUE
        ELSE AND (ISCHAR (HEAD S)) (ISSTRING (TAIL S))
```

## 6.11. String comparison

```
    rec STRING_EQUAL S1 S2 =
        IF (ISNIL S1) AND (ISNIL S2)
        THEN TRUE
        ELSE
            IF (ISNIL S1) OR (ISNIL S2)
            THEN FALSE
            ELSE
                IF CHAR_EQUAL (HEAD S1) (HEAD S2)
                THEN STRING_EQUAL (TAIL S1) (TAIL S2)
                ELSE FALSE

    rec STRING_LESS S1 S2 =
        IF ISNIL S1
        THEN NOT (ISNIL S2)
        ELSE
            IF ISNIL L2
            THEN FALSE
            ELSE
                IF CHAR_LESS (HEAD S1) (HEAD S2)
                THEN TRUE
                ELSE (CHAR_EQUAL (HEAD S1) (HEAD S2)) AND (STRING_LESS1 (TAIL S1) (TAIL S2))
```

## 6.12. Numeric string to number conversion

```
    def digit_value d = sub (value d) (value ’0’)

    rec string_val v L =
        IF ISNIL L
        THEN v
        ELSE string_val (add (mult v ten) (digit_value (HEAD L)))
                        (TAIL L)

    def STRING_VAL S = MAKE_NUMB (string_val zero S)
```

## 6.13. Structure matching with lists

```
    rec <name> [] = <expression1>
    or <name> (<head>::<tail>) = <expression2 using <head>
                                and <tail>>

    rec LENGTH L =
        IF ISNIL L
        THEN 0
        ELSE SUCC (LENGTH (TAIL L))

    rec FLAT [] = []
    or FLAT (H::T) =
        IF NOT (ISLIST H)
        THEN H::(FLAT T)
        ELSE APPEND (FLAT H) (FLAT T)
```

## 6.14. Ordered linear lists, insertion and sorting

```
    rec INSERT S [] = [S]
    or INSERT S (H::T) =
        IF STRING_LESS S H
        THEN S::H::T
        ELSE H::(INSERT H T)

    rec SORT [] = []
    or SORT (H::T) = INSERT H (SORT T)

    rec IFIND N [] = []
    or IFIND 0 (H::T) = H
    or IFIND (SUCC N) (H::T) = IFIND N T

    rec IDELETE N [] = []
    or IDELETE 0 (H::T) = H
    or IDELETE (SUCC N) (H::T) = H::(IDELETE N T)

    rec IBEFORE N E [] = []
    or IBEFORE 0 E L = E::L
    or IBEFORE (SUCC N) E (H::T) = H::(IBEFORE N E T)

    rec IREPLACE N E [] = []
    or IREPLACE 0 E (H::T) = E::T
    or IREPLACE (SUCC N) E (H::T) = H::(IREPLACE N E T)
```

## 6.16. Mapping functions

CAR mapping:
```
    rec MAPCAR FUNC [] = []
    or MAPCAR FUNC (H::T) = (FUNC H)::(MAPCAR FUNC T)

    rec MAPCARS FUNC [] [] = []
    or MAPCARS FUNC (H1::T1) (H2::T2) = (FUNC H1 H2)::(MAPCARS FUNC T1 T2)
```