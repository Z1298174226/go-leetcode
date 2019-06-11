package main

import (
	"fmt"
)

type pointInt *[]int

func main() {
	arr := []int{1, 5, 2, 4, 7, 1, 2, 99, 4, 6}
	//quickSort(0, len(arr) - 1, &arr)
	heapSort(&arr)
	fmt.Println(arr)
}

/**
快排
*/
func quickSort(left int, right int, arr pointInt) {
	if left >= right {
		return
	} else {
		//j := partition(left, right, arr)
		j := partitionUpdate(left, right, arr)
		quickSort(left, j-1, arr)
		quickSort(j+1, right, arr)
	}
}

func partition(left int, right int, arr pointInt) int {
	i := left
	j := right + 1
	for {
		for {
			i = i + 1
			if lessInt(left, i, arr) {
				break
			} else {
				if i == right {
					break
				}
			}
		}
		for {
			j = j - 1
			if lessInt(j, left, arr) {
				break
			} else {
				if j == left {
					break
				}
			}
		}
		if i >= j {
			break
		}
		swapInt(i, j, arr)
	}
	swapInt(left, j, arr)
	return j
}

func partitionUpdate(left int, right int, arr pointInt) int {
	i := left - 1
	for j := left; j < right; j++ {
		if lessInt(j, right, arr) {
			i++
			swapInt(i, j, arr)
		}
	}
	i++
	swapInt(i, right, arr)
	return i
}

func swapInt(i int, j int, arr pointInt) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func lessInt(i int, j int, arr pointInt) bool {
	return (*arr)[i] < (*arr)[j]
}

/**
堆排序
*/

func swapHeap(i int, j int, arr pointInt) {
	(*arr)[i-1], (*arr)[j-1] = (*arr)[j-1], (*arr)[i-1]
}

func lessHeap(i int, j int, arr pointInt) bool {
	return (*arr)[i-1] < (*arr)[j-1]
}

func sinkOperation(left int, right int, arr pointInt) {
	harf := right >> 1
	for {
		if left > harf {
			return
		}
		child := left << 1
		if child < right && lessHeap(child, child+1, arr) {
			child++
		}
		if lessHeap(left, child, arr) {
			swapHeap(left, child, arr)
		}
		left = child
	}
}

func heapSort(arr pointInt) {
	len := len(*arr)
	for i := len / 2; i >= 1; i-- {
		sinkOperation(i, len, arr)
	}
	for {
		if len <= 1 {
			return
		} else {
			swapHeap(1, len, arr)
			len--
			sinkOperation(1, len, arr)
		}
	}
}
