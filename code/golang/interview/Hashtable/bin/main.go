package main

import . ".."

func main() {
	d := NewDict(16)
	d.Insert(NewThing("Jack", 21).Hashcode())
	d.Insert(NewThing("Tom", 12).Hashcode())
	d.Insert(NewThing("Chuck", 23).Hashcode())
	d.Insert(NewThing("Tim", 18).Hashcode())
	d.Insert(NewThing("Peter", 46).Hashcode())
	d.Insert(NewThing("Allen", 34).Hashcode())
	d.Insert(NewThing("Bob", 32).Hashcode())
	d.Insert(NewThing("Lily", 27).Hashcode())
	d.Insert(NewThing("Steven", 19).Hashcode())
	d.Print()
}
