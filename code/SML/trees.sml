datatype inttree = empty | node of int * inttree * inttree;

fun add (v:int) empty = node(v,empty,empty) |
    add (v:int) (node(nv:int,l:inttree,r:inttree)) =
        if v < nv
        then node(nv,add v l,r)
        else node(nv,l,add v r);

fun last l [] = l |
    last _ (h::t) = last h t;

fun front l [_] = l |
    front p (h::t) = front (p@[h]) t;

fun append l1 l2 = l1@l2;

fun traverse empty = [] |
    traverse (node(v:int,l:inttree,r:inttree)) =
        append (traverse l) (v::traverse r);

val root = empty;
val root = add 5 root;
val root = add 3 root;
val root = add 7 root;
val root = add 2 root;
val root = add 4 root;
val root = add 9 root;

traverse root;