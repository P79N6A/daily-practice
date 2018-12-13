package main

import (
	. ".."
)

func main() {
	lru := NewLRUCache(3)
	lru.Set("name", "Tom")
	lru.Set("age", 21)
	lru.Set("mail", "tom@gmail.com")
	lru.Set("title", "Coder")

	lru.Print()
}
