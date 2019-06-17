package main

import (
	"array"
	"dp"
	"fmt"
	"uf"
)

func main() {
	arr := []int{1, 3, 4, 5, 7, 8, 9, 10}
	fmt.Println(arr)
	fmt.Println(uf.ContinousSebsequence(&arr))
	fmt.Println(dp.StockI(&arr))
	fmt.Println(dp.StockII(&arr))
	fmt.Println(dp.StockIII(&arr, 1))
	fmt.Println(array.LongestIncreasingSubsequence(&arr))
	fmt.Println(dp.IsMatcher("abcd", "abc*"))
	gold := []int{3, 9, 5, 2, 4, 4}
	fmt.Println(array.DivideGold(&gold))
	fmt.Println(array.FindTopKs(&gold, 2))
	fmt.Println(array.FindTopKth(&gold, 6))
}
