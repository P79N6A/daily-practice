package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "雷吼，Mr.Smith."

    for i := 0; i < len(s); {
        r, size := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%d\t%c\n", i, r)
        i += size
    }

    for i, r := range s {
        fmt.Printf("%d\t%q\t%d\n", i, r, r)
    }

    // "program" in Japanese katakana
    s2 := "プログラム"
    fmt.Printf("% x\n", s2) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
    r := []rune(s2)
    fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
}
