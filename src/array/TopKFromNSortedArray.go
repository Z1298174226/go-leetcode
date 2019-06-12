package array

import "fmt"

type point *[]Node

type Node struct {
	val   int
	index int
	count int
}

func FindTopK(input [][]int, k int) {
	len := len(input)
	arr := Inits(input)
	for i := 0; i < k; i++ {
		fmt.Println((*arr)[1].val)
		index := (*arr)[1].index
		count := (*arr)[1].count
		if index > 0 {
			(*arr)[1] = Node{
				val:   input[count][index-1],
				index: index - 1,
				count: count}
			Sink(1, len, arr)
		} else {
			Swap(1, len, arr)
			len = len - 1
			Sink(1, len, arr)
		}
	}
}

func Inits(input [][]int) point {
	N := len(input)
	len := len(input[0])
	arr := make([]Node, N+1)
	for i := range input {
		node := Node{
			val:   input[i][len-1],
			index: len - 1,
			count: i}
		arr[i+1] = node
		Sink(1, i+1, &arr)
	}
	return &arr
}

func Sink(left int, right int, arr point) {
	harf := right >> 1
	for {
		if left > harf {
			return
		} else {
			child := left << 1
			if child+1 <= right && Less(child, child+1, arr) {
				child++
			}
			if Less(left, child, arr) {
				Swap(left, child, arr)
			}
			left = child
		}
	}
}

func Swap(i int, j int, arr point) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func Less(i int, j int, arr point) bool {
	return (*arr)[i].val < (*arr)[j].val
}
