package main

import (
	"fmt"
	"uf"
)

func main() {
	arr := []int{1, 3, 4, 5, 7, 8, 9, 10}
	fmt.Println(arr)
	fmt.Println(uf.ContinousSebsequence(&arr))
}
