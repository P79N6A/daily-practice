# Chapter 4. Recursion and arithmetic

## SUMMARY

### Applicative order β reduction
For (&lt;function expression&gt; &lt;argument expression&gt;)

1. applicative order β reduce &lt;argument expression&gt; to &lt;argument value&gt;
2. applicative order β reduce &lt;function expression&gt; to &lt;function value&gt;
3. if &lt;function value&gt; is λ&lt;name&gt;.&lt;body&gt;
then replace all free occurences of &lt;name&gt; in &lt;body&gt; with &lt;argument value&gt; and applicative order β
reduce the new &lt;body&gt;
4. if &lt;function value&gt; is not a function
then return (&lt;function value&gt; &lt;argument value&gt;)

### Recursion function
```
    def recursive f = λs.(f (s s)) λs.(f (s s))
```

### Recursion notation
```
    rec <name> = <expression using ‘<name>’> ==
        def <name> = recursive λf.<expression using ‘f’>
```