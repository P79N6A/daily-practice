fun insert _ i [] = [i] |
    insert comp i (h::t) =
    if comp (i,h)
    then i::h::t
    else h::(insert comp i t);

val sinsert = insert (fn (s1:string,s2:string) => s1<s2);

sinsert "zelina" ["hello", "abc", "world"];