package main

import (
	"bytes"
	"fmt"
)

// Longest Common Prefix
func LCP(strings []string) string {
	var prefix bytes.Buffer

	for i := 0; i < len(strings[0]); i++ {
		for j := 1; j < len(strings); j++ {
			if strings[0][i] != strings[j][i] {
				return prefix.String()
			}
		}
		prefix.WriteByte(strings[0][i])
	}

	return prefix.String()
}

func main() {
	strings := []string{
		"/holy",
		"/home",
		"/hot",
	}
	lcp := LCP(strings)
	fmt.Println(lcp)
}
