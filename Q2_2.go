package main

import "fmt"

func main() {
	i := 0
L:
	fmt.Printf("%d\n", i)
	i++
	if i < 10 {
		goto L
	}
}
