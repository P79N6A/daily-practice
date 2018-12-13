datatype 'a tree = empty | node of 'a * ('a tree) * ('a tree);

fun tadd _ (v:'a) empty = node(v, empty, empty) |
    tadd (less:'a -> ('a -> bool))
         (v:'a)
         (node(nv:'a, l:'a tree, r:'a tree)) =
            if less v nv
            then node(nv, tadd less v l, r)
            else node(nv, l, tadd less v r);

fun append l r = l @ r;

fun traverse empty = [] |
    traverse (node(v:'a, l:'a tree, r:'a tree)) =
        append (traverse l) (v::traverse r);

val less = fn x => fn y => x < y;

val root = empty;
val root = tadd less 5 root;
val root = tadd less 3 root;
val root = tadd less 7 root;
val root = tadd less 2 root;
val root = tadd less 4 root;
val root = tadd less 9 root;

traverse root;

(* 
- datatype 'a tree = empty | node of 'a * 'a tree * 'a tree
val tadd = fn : ('a -> 'a -> bool) -> 'a -> 'a tree -> 'a tree
val append = fn : 'a list -> 'a list -> 'a list
val traverse = fn : 'a tree -> 'a list
val less = fn : int -> int -> bool
val root = empty : 'a tree
val root = node (5,empty,empty) : int tree
val root = node (5,node (3,empty,empty),empty) : int tree
val root = node (5,node (3,empty,empty),node (7,empty,empty)) : int tree
val root = node (5,node (3,node #,empty),node (7,empty,empty)) : int tree
val root = node (5,node (3,node #,node #),node (7,empty,empty)) : int tree
val root = node (5,node (3,node #,node #),node (7,empty,node #)) : int tree
val it = [2,3,4,5,7,9] : int list
*)