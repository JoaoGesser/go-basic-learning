package main

import "fmt"

func main() {
	i, h := 10, 20

	p := &i
	q := h
	*p = 15
	fmt.Println(i)
	fmt.Println(q)

}
