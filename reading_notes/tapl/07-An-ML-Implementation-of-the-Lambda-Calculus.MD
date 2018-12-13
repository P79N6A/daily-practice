# Chapter 7. An ML Implementation of the Lambda-Calculus

## 7.1 Terms and Contexts

the type info recording the file position where that term was originally found:
```
    type term =
          TmVar of info * int * int
        | TmAbs of info * string * term
        | TmApp of info * term * term
```

The printing routine itself looks like this:
```ml
    let rec printtm ctx t = match t with
          TmAbs(fi,x,t1) ➝
            let (ctx',x') = pickfreshname ctx x in
            pr "(lambda "; pr x'; pr ". "; printtm ctx' t1; pr ")"
        | TmApp(fi, t1, t2) ➝
            pr "("; printtm ctx t1; pr " "; printtm ctx t2; pr ")"
        | TmVar(fi,x,n) ➝
            if ctxlength ctx = n then
                pr (index2name fi ctx x)
            else
                pr "[bad index]"
```

It uses the datatype context,
```ml
    type context = (string * binding) list
```

## 7.2 Shifting and Substitution

The definition of shifting (6.2.1) can be translated almost symbol for symbol into OCaml:
```ml
    let termShift d t =
        let rec walk c t = match t with
              TmVar(fi,x,n) ➝ if x>=c then TmVar(fi,x+d,n+d)
                            else TmVar(fi,x,n+d)
            | TmAbs(fi,x,t1) ➝ TmAbs(fi, x, walk (c+1) t1)
            | TmApp(fi,t1,t2) ➝ TmApp(fi, walk c t1, walk c t2)
        in walk 0 t
```

The substitution [j ➝ s]t of term s for the variable numbered j in term t is written as termSubst j s t here:
```ml
    let termSubst j s t =
        let rec walk c t = match t with
              TmVar(fi,x,n) ➝ if x=j+c then termShift c s else TmVar(fi,x,n)
            | TmAbs(fi,x,t1) ➝ TmAbs(fi, x, walk (c+1) t1)
            | TmApp(fi,t1,t2) ➝ TmApp(fi, walk c t1, walk c t2)
        in walk 0 t
```

In the operational semantics of the lambda-calculus, the only place where substitution is used is in the beta-reduction rule. As we noted before, this rule actually performs several operations: the term being substituted for the bound variable is first shifted up by one, then the substitution is made, and then the whole result is shifted down by one to account for the fact that the bound variable has been used up. The following definition encapsulates this sequence of steps:
```ml
    let termSubstTop s t =
        termShift (-1) (termSubst 0 (termShift 1 s) t)
```

## 7.3 Evaluation

the evaluation function depends on an auxiliary predicate isval:
```
    let rec isval ctx t = match t with
          TmAbs(_,_,_) ➝ true
        | _ ➝ false
```

The single-step evaluation function is a direct transcription of the evaluation rules, except that we pass a context ctx along with the term. This argument is not used in the present eval1 function, but it is needed by some of the more complex evaluators later on.
```ml
    let rec eval1 ctx t = match t with
          TmApp(fi,TmAbs(_,x,t12),v2) when isval ctx v2 ➝
            termSubstTop v2 t12
        | TmApp(fi,v1,t2) when isval ctx v1 ➝
            let t2' = eval1 ctx t2 in
            TmApp(fi, v1, t2')
        | TmApp(fi,t1,t2) ➝
            let t1' = eval1 ctx t1 in
            TmApp(fi, t1', t2)
        | _ ➝
            raise NoRuleApplies
```

The multi-step evaluation function is the same as before, except for the ctx argument:
```ml
    let rec eval ctx t =
        try let t' = eval1 ctx t
            in eval ctx t'
        with NoRuleApplies ➝ t
```