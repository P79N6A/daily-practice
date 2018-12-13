# Chapter 2. λ calculus

## SUMMARY
* Abstraction has a central role in programming languages, and λ calculus is a language based on pure abstraction.
* λ calculus syntax has been introduced and used to analyse the structure of some simple expressions.
* Normal order β reduction is used to reduce simple expressions. Note that not all reductions terminate.
* Notations have been introduced for defining functions and simplifying familiar reduction sequences.
* Functions may be constructed from other functions.
* Functions have been built to construct pairs of values and select values from them.
* Normal order β reduction can be formalized in terms of substitution for free variables.
* α conversion is a way of removing name clashes in expressions.
* η reduction is a way of simplifying expressions.

Some of these topics are summarized below.

#### ◉ λ calculus syntax
    <expression> ::= <name> | <function> | <expression>
    <name> ::= non-blank character sequence
    <function> ::= λ<name>.<body>
        <body> ::= <expression>
    <application> ::= (<function expression>
                        <argument expression>)
        <function expression> ::= <expression>
        <argument expression> ::= <expression>

#### ◉ Free variables
    • <name> is free in <name>.
    • <name> is free in λ<name1>.<body>
        if <name1> is not <name>
        and <name> is free in <body>.
    • <name> is free in <function expression>
                        <argument expression>)
        if <name> is free in <function expression>
        or <name> is free in <argument expression>.

#### ◉ Bound variables
    • <name> is bound in λ<name1>.<body>
        if <name> is <name1>
        or <name> is bound in <body>.
    • <name> is bound in (<function expression>
                          <argument expression>)
        if <name> is bound in <function expression>
        or <name> is bound in <argument expression>.

#### ◉ Normal order β reduction
For (<function expression> <argument expression>), containing no free variables:
    (1) Normal order β reduce <function expression> to <function value>.
    (2) If <function value> is λ<name>.<body>
        then replace all free occurrences of <name> in <body> with
            <argument expression>
        and β reduce the new <body>.
    or
    (3) If <function value> is not a function
        then normal β reduce <argument expression> to
            <argument value>
        and return (<function value> <argument value>).

#### ◉ Normal order reduction notation
    =>          - normal order β reduction
    => ... =>   - multiple normal order β reduction

#### ◉ Definitions
    def <name> = <expression>
Replace all subsequent occurrences of <name> with <expression> before evaluation.

#### ◉ Replacement notation
    ==          - defined name replacement

#### ◉ α conversion
    To rename <name1> as <name2> in λ<name1>.<body>
    if <name2> is not free in λ<name1>.<body>
    then replace all free occurrences of <name1> in <body> with
        <name2>
    and replace <name1> in λ<name1>

#### ◉ η reduction
    (λ<name>.(<expression> <name>) <argument>)  =>
    <expression> <argument>
