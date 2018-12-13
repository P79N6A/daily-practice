# Chapter 7. COMPOSITE VALUES AND TREES

## 7.5. Generalised structure matching

```
    def FORENAME [F,S] = F

    def SURNAME [F,S] = S

    [F,S] == F::S::NIL

    rec NCOUNT N [] = 0
    or NCOUNT N ([F,S]::T) =
        IF EQUAL N F
        THEN 1 + (NCOUNT N T)
        ELSE (NCOUNT N T)

    rec DINSERT R [] = [R]
    or DINSERT [N1,A1,P1] ([N2,A2,P2]::T) =
        IF LESS P1 P2
        THEN [N1,A1,P1]::[N2,A2,P2]::T
        ELSE [N2,A2,P2]::(DINSERT [N1,A1,P1] T)

    rec DSORT [] = []
    or DSORT (H::T) = DINSERT H (DSORT T)
```

## 7.7. Matching composite value results

```
    rec SPLIT [] = []::[]
    or SPLIT ([F,S]::L) =
        let (FLIST::SLIST) = SPLIT L
        in ((F::FLIST)::(S::SLIST))
```

## 7.9. Trees

```
    def EMPTY   = NIL
    def ISEMPTY = ISNIL

    def NODE ITEM L R = [ITEM,L,R]

    def TREE_ERROR = LIST_ERROR

    def ITEM EMPTY = TREE_ERROR
    or ITEM [I,L,R] = I
    
    def LEFT EMPTY = TREE_ERROR
    or LEFT [I,L,R] = L
    
    def RIGHT EMPTY = TREE_ERROR
    or RIGHT [I,L,R] = R

    rec TADD I EMPTY = NODE I EMPTY EMPTY
    or TADD I [NI,L,R] =
        IF LESS I NI
        THEN NODE NI (TADD I L) R
        ELSE NODE NITEM L (TADD I R)
```

## 7.11. Binary tree traversal

```
    rec TRAVERSE EMPTY = []
    or TRAVERSE [I,L,R] = APPEND (TRAVERSE L) (I::(TRAVERSE R))
```

## 7.12. Binary tree search

```
    rec TFIND V EMPTY = ""
    or TFIND V [NV,L,R] =
        IF EQUAL V NV
        THEN TRUE
        ELSE
            IF LESS V NV
            THEN TFIND V L
            ELSE TFIND V R
```

## 7.13. Binary trees of composite values

```
    rec CTADD N EMPTY = [N,EMPTY,EMPTY]
    or CTADD [F,S] [[NF,NS],L,R] =
        IF STRING_LESS S NS
        THEN [[NF,NS],(CTADD [F,S] L),R]
        ELSE [[NF,NS],L,(CTADD [F,S] R)]

    rec CTADDLIST [] TREE = TREE
    or CTADDLIST (H::T) TREE = CTADDLIST T (CTADD H TREE)
```
