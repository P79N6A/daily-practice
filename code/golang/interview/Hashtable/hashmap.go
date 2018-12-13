package Hashtable

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

func CalcHashcode(data interface{}) uint64 {
	buf := new(bytes.Buffer)
	hash := fnv.New64a()
	switch v := data.(type) {
	case int:
		binary.Write(buf, binary.BigEndian, v)
	case string:
		buf = bytes.NewBufferString(v)
	}
	hash.Write(buf.Bytes())
	return hash.Sum64()
}

type Dict struct {
	distribution map[uint64]uint
	M            uint64
}

func (d *Dict) Insert(k uint64) {
	p := k % d.M
	d.distribution[p]++
}
func (d *Dict) Print() {
	fmt.Printf("%v \n", d.distribution)
}

func NewDict(m uint64) *Dict {
	return &Dict{
		distribution: make(map[uint64]uint),
		M:            m,
	}
}
