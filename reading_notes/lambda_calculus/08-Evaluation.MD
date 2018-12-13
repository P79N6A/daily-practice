# Chapter 8. Evaluation

## 8.1 Termination and normal form
A lambda expression which cannot be reduced any further is said to be in **normal form**.

## 8.2 Normal Order
Normal order β reduction requires the evaluation of the leftmost redex in a expression.

(argument evaluation is delayed)

## 8.3 Applicative Order
Applicative order β reduction of an application requires the evaluation ofs both the function and the argument expressions.

## 8.6 Evaluation termination, the halting problem, evaluation equivalance and the Church-Rosser theorems
1. Church-Rosser theorem shows that every expression has a unique normal form. Thus, if an expression is reduced using two different evaluation orders and both reductions terminate then they both lead to the same normal form.
2. Church-Rosser theorem shows that if an expression has a normal form then it may be reached by normal order evaluation. In other words, if any evaluation order will terminate then normal order evaluation is guaranteed to terminate.

## 8.9. Lazy evaluation
Lazy evaluation is a method of delaying expression evaluation which avoids multiple evaluation of the same expression. Hence, lazy evaluation is also known as call by need.

1. number every bound pair
2. To lazy evaluate (&lt;function expression&gt; &lt;argument expression&gt;)i
    1. lazy evaluate &lt;function expression&gt; to &lt;function value&gt;
    2. if &lt;function value&gt; is λ&lt;name&gt;.&lt;body&gt; then replace all free occurences of &lt;name&gt; in &lt;body&gt; with &lt;argumentexpression&gt; and renumber consistently all surrounding bound pairs and replace all occurences of (&lt;function expression&gt; &lt;argument expression&gt;)i with the new &lt;body&gt; and lazy evaluate the new &lt;body&gt;

        or

        if &lt;function value&gt; is not a function then lazy evaluate &lt;argument expression&gt;to &lt;argument value&gt; and replace all occurences of (&lt;function expression&gt; &lt;argument expression&gt;)i with (&lt;function value&gt; &lt;argument value&gt;) and return (&lt;function value&gt; &lt;argument value&gt;)