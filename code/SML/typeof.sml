datatype wrapper = Int of int | Real of real | String of string;

fun typeof x = 
	case x of Int(x)      => "[Int]:" ^ Int.toString(x)
			| Real(x)     => "[Real]:" ^ Real.toString(x) 
			| String(x)   => "[String]:" ^ x;

val ty1 = Int(32);
typeof(ty1); 
val ty2 = Real(22.2);
typeof(ty2);
val ty3 = String("hello world");
typeof(ty3);
