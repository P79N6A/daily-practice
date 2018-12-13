# Basics
```lambda
⟼ (untyped)
------------------------------------------------------------
Syntax:

t ::=                                  terms:
        x                            variable
        ℷx.t                      abstraction
        t t                       application

v ::=                                 values:
        ℷx.t                abstraction value
------------------------------------------------------------
Evaluation:                         [t ⟼ t']

           t1 ⟼ t1'
        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯             (E-APP1)
        t1 t2 ⟼ t1' t2

           t2 ⟼ t2'
        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯             (E-APP2)
        v1 t2 ⟼ v1 t2'

   (ℷx.t12)v2 ⟼ [x ⟼ v2]t12      (E-APP ABS)
------------------------------------------------------------
```

# Church Booleans
```lambda
tru = ℷt.ℷf. t;
fls = ℷt.ℷf. f;
and = ℷb.ℷc. b c fls;
```
# Pairs
```lambda
pair = ℷf.ℷs.ℷb. b f s;
fst  = ℷp. p tru;
snd  = ℷp. p fls;
```

# Church Numerals
```lambda
c0 = ℷs.ℷz. z;
c1 = ℷs.ℷz. s z;
c2 = ℷs.ℷz. s (s z);
c3 = ℷs.ℷz. s (s (s z));
etc.

scc   = ℷn.ℷs.ℷz. s (n s z);
plus  = ℷm.ℷn.ℷs.ℷz. m s (n s z);
times = ℷm.ℷn. m (plus n) c0;

iszro = ℷm. m (ℷx. fls) tru;

zz  = pair c0 c0;
ss  = ℷp. pair (snd p) (plus c1 (snd p));
prd = ℷm. fst (m ss zz);
```

# Exercises:
* Write a function equal that tests two numbers for equality and returns a Church boolean. For example：
```lambda
equal c3 c3;
> (ℷt.ℷf. t)
equal c3 c2;
> (ℷt.ℷf. f)
```
Other common datatypes like lists, trees, arrays, and variant records can be encoded using similar techniques.
```lambda
equal = ℷm.ℷn. and (iszro (m prd n)) (iszro (n prd m))
```

* A list can be represented in the lambda-calculus by its fold function. (OCaml’s name for this function is fold_left; it is also sometimes called reduce .) For example, the list [x,y,z] becomes
a function that takes two arguments c and n and returns c x (c y (c z n))). What would the representation of nil be? Write a function cons that takes an element h and a list (that is, a fold function) t and returns a similar representation of the list formed by prepending h to t. Write isnil and head functions, each taking a list parameter. Finally, write a tail function for this representation of lists (this is quite a bit harder and requires a trick analogous to the one used to define prd for numbers).
```lambda
nil   = pair tru tru;
isnil = fst;
head  = ℷl. fst (snd l);
tail  = ℷl. snd (snd l);
cons  = ℷh.ℷt pair fls (pair h t);
```

# Enriching the Calculus
```lambda
realbool   = ℷb. b true false;
churchbool = ℷb. if b then tru else fls;
realeq     = ℷm.ℷn. (equal m n) true false;
realnat    = ℷm. m (ℷx. succ x) 0;
```

# Recursion
```lambda
omega = (ℷx. x x) (ℷx. x x);

fix = ℷf. (ℷx. f (ℷy. x x y)) (ℷx. f (ℷy. x x y));

g = ℷfct.ℷn. if realeq n c0 then c1 else (times n (fct (prd n)));
factorial = fix g;
```
