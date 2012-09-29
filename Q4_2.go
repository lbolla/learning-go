package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s string = "asSASA ddd dsjkdsjs dk"
	fmt.Println(len([]byte(s)), utf8.RuneCount([]byte(s)))
}
