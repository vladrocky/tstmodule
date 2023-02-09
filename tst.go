package tstmodule

import "fmt"

func Add(a, b int) int {
	fmt.Println(a + b + a)
	return a + b + a
}
