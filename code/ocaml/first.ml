let rec iseven a =
  if a = 0 then true
  else if a - 1 == 0 then false
  else iseven(a - 2);;


let a = 12;;
Printf.printf "iseven(%d) == %b\n" a (iseven 12) ;;
