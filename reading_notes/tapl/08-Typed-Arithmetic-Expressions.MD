# Chapter 8. Typed Arithmetic Expressions

## 8.2 The Typing Relation

Typing rules for booleans:
```
------------------------------------------------------------------------------------
New syntactic forms                             New typing rules              [t: T]
T ::=                         types:                 true:  Bool            (T-True)
    Bool            type of booleans                 false: Bool           (T-False)

                                                t1: Bool   t2: T   t3: T
                                                -------------------------     (T-If)
                                                if t1 then t2 else t3: T
------------------------------------------------------------------------------------
```

Typing rules for numbers:
```
------------------------------------------------------------------------------------
New syntactic forms                             t1: Nat
T ::= ...                      types:           ------------                (T-Succ)
    Nat       type of natural numbers           succ t1: Nat

New typing rules               [t : T]          t1: Nat
            0 : Nat           (T-Zero)          ------------                (T-Pred)
                                                pred t1: Nat

                                                t1 : Nat
                                                ---------------            (T-IsZero)
                                                iszero t1 : Bool
------------------------------------------------------------------------------------
```

### 8.2.2 Lemma [Inversion of the typing relation]:

1. If true : R, then R = Bool.
2. If false : R, then R = Bool.
3. If if t1 then t2 else t3 : R, then t1 : Bool, t2 : R, and t3 : R.
4. If 0 : R, then R = Nat.
5. If succ t1 : R, then R = Nat and t1 : Nat.
6. If pred t1 : R, then R = Nat and t1 : Nat.
7. If iszero t1 : R, then R = Bool and t1 : Nat.    

## 8.3 Safety = Progress + Preservation

### 8.3.2 Theorem [Progress]: Suppose t is a well-typed term (that is, t : T for some T). Then either t is a value or else there is some t' with t ➝ t'.

### 8.3.3 Theorem [Preservation]: If t : T and t ➝ t', then t' : T.
