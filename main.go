package main

import "fmt"

type list []any

func (l list) len() int {
	return len()
}

func main() {
	var ml list
	ml = []any{"1", "2", "3"}
	fmt.Println(ml.len())
}
