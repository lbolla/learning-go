package main

import (
	"fmt"
)

func main() {
	var s string = "asSASA ddd dsjkdsjs dk"
	fmt.Println(s)
	c := []rune(s)
	r := []rune("abc")
	for i := 0; i < len(r); i++ {
		c[4 + i] = r[i]
	}
	fmt.Println(string(c))
}
