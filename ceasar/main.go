package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	// fmt.Printf("length: %d\n", length)
	// fmt.Printf("input: %s\n", input)
	// fmt.Printf("delta: %d\n", delta)

	alphabet := []rune("abcdefghijklmnopqestuvwxyz")
	alphabetStr := "abcdefghijklmnopqrstuwxyz"
	ret := ""
	for _, ch := range input {
		if strings.IndexRune(alphabetStr, ch) >= 0 {
			ret = ret + string(rotate(ch, delta, alphabet))
		} else {
			ret = ret + string(ch)
		}
	}
	fmt.Println(ret)
}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
}
