package main

import "fmt"

type point *[]Node

func main() {
	input := [][]int{
		{1, 3, 5, 7, 9},
		{2, 4, 6, 8, 10},
		{3, 7, 8, 9, 11},
	}
	findTopK(input, 15)
}

type Node struct {
	val   int
	index int
	count int
}

func findTopK(input [][]int, k int) {
	len := len(input)
	arr := inits(input)
	for i := 0; i < k; i++ {
		fmt.Println((*arr)[1].val)
		index := (*arr)[1].index
		count := (*arr)[1].count
		if index > 0 {
			(*arr)[1] = Node{
				val:   input[count][index-1],
				index: index - 1,
				count: count}
			sink(1, len, arr)
		} else {
			swap(1, len, arr)
			len = len - 1
			sink(1, len, arr)
		}
	}
}

func inits(input [][]int) point {
	N := len(input)
	len := len(input[0])
	arr := make([]Node, N+1)
	for i := range input {
		node := Node{
			val:   input[i][len-1],
			index: len - 1,
			count: i}
		arr[i+1] = node
		sink(1, i+1, &arr)
	}
	return &arr
}

func sink(left int, right int, arr point) {
	harf := right >> 1
	for {
		if left > harf {
			return
		} else {
			child := left << 1
			if child+1 <= right && less(child, child+1, arr) {
				child++
			}
			if less(left, child, arr) {
				swap(left, child, arr)
			}
			left = child
		}
	}
}

func swap(i int, j int, arr point) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func less(i int, j int, arr point) bool {
	return (*arr)[i].val < (*arr)[j].val
}
