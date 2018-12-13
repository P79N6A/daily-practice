package Hashtable

type Thing struct {
	Name string
	Qty  int
}

func (t *Thing) Hashcode() (hash uint64) {
	hash = 17
	hash = 31*hash + CalcHashcode(t.Name)
	hash = 31*hash + CalcHashcode(t.Qty)
	return
}

func NewThing(name string, qty int) *Thing {
	return &Thing{
		Name: name,
		Qty:  qty,
	}
}
