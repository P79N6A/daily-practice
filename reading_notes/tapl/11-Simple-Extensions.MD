# Chapter 11. Simple Extensions

## 11.2 The Unit Type

```lambda
⟼ Unit
------------------------------------------------------------
New syntactic forms                      
t ::=   ...                                      terms:
        unit                              constant unit

v ::=   ...                                     values:
        unit                              constant unit

T ::=   ...                                      types:
        Unit                                  unit type
------------------------------------------------------------
New typing rules                             [𝜞 ⊢ t: T]
        𝜞 ⊢ unit: Unit                         (T-UNIT)

New derived forms
        t1;t2 ≝ (ℷx:Unit.t2)t1
                where x ∉ FV(t2)
------------------------------------------------------------
```

This use of Unit is similar to the role of the void type in languages like C and Java. The name void suggests a connection with the empty type Bot, but the usage of void is actually closer to our Unit.

## 11.3 Derived Forms: Sequencing and Wildcards
```lambda
evaluation rules:
            t1 ⟼ t1'
        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯                      (E-SEQ)
        t1;t2 ⟼ t1';t2

        unit;t2 ⟼ t2                     (E-SeqNext)

typing rule:
        𝜞 ⊢ t1:Unit     𝜞 ⊢ t2:T2
        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯             (T-SEQ)
             𝜞 ⊢ t1;t2: T2
```

### 11.3.1 Theorem [Sequencing is a derived form]

Let ```e ∈ ℷE➝ℷI``` be the elaboration function that translates from the external to the internal language by replacing every occurrence of t1;t2 with ```(ℷx:Unit.t2)t1```, where x is chosen fresh in each case. Now, for each term t of E, we have:
```
•       t ➝E t' iff e(t) ➝I e(t')

•       𝜞 ⊢E t:T iff 𝜞 ⊢I e(t):T
```

## 11.4 Ascription

```lambda
⟼ as
------------------------------------------------------------
New syntactic forms
t ::=                                  terms:
        t as T                      acription

New evaluation rules                [t ⟼ t']
       v1 as T ⟼ v1              (E-Ascribe)
        
             t1 ⟼ t1'
       ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯       (E-Ascribe1)
       t1 as T ⟼ t1' as T
------------------------------------------------------------
New typing rules                     [𝜞 ⊢ t:T]

           𝜞 ⊢ t1:T
        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯              (E-Ascribe)
        𝜞 ⊢ t1 as T:T
------------------------------------------------------------       
```

## 11.4 Let Bindings

```lambda
⟼ let
------------------------------------------------------------
New syntactic forms
t ::=                                                 terms:
        let x=t in t                             let binding

New evaluation rules                               [t ⟼ t']
        let x=v1 in t2 ⟼ [x⟼v1]t2                (E-LetV)
------------------------------------------------------------
                  t1 ⟼ t1'
        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯            (E-Let)
        let x=t1 in t2 ⟼ let x=t1' in t2


New typing rules                                    [𝜞 ⊢ t:T]

           𝜞 ⊢ t1:T1    𝜞,x:T1 ⊢ t2:T2  
        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯            (T-Let)
              𝜞 ⊢ let x=t1 in t2:T2
------------------------------------------------------------       
```
