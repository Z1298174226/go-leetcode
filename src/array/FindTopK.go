package array

import "fmt"

func FindTopKs(arr *[]int, k int) *[]int {
	len := len(*arr)
	heap := initHeap(arr, k)
	fmt.Println(heap)
	for i := k >> 1; i >= 1; i-- {
		sinkHeap(i, k, heap)
	}
	for i := k; i < len; i++ {
		if (*arr)[i] > (*heap)[0] {
			(*heap)[0] = (*arr)[i]
			sinkHeap(1, k, heap)
		}
	}
	return heap
}

func initHeap(arr *[]int, k int) *[]int {
	heap := make([]int, k)
	for i := 0; i < k; i++ {
		heap[i] = (*arr)[i]
	}
	return &heap
}

func sinkHeap(left, right int, arr *[]int) {
	half := right >> 1
	for {
		if left > half {
			break
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

func swapHeap(a, b int, arr *[]int) {
	(*arr)[a-1], (*arr)[b-1] = (*arr)[b-1], (*arr)[a-1]
}

func lessHeap(a, b int, arr *[]int) bool {
	return (*arr)[a-1] > (*arr)[b-1]
}

func FindTopKth(arr *[]int, k int) int {
	len := len(*arr)
	return findTopKth(arr, 0, len-1, k)
}

func findTopKth(arr *[]int, left, right, k int) int {
	for {
		if left > right {
			break
		}
		partition := partition(arr, left, right)
		if partition == k-1 {
			return (*arr)[partition]
		} else if partition > k-1 {
			return findTopKth(arr, left, partition-1, k)
		} else {
			return findTopKth(arr, partition+1, right, k)
		}
	}
	return -1
}

func partition(arr *[]int, left, right int) int {
	i := left - 1
	for j := left; j < right; j++ {
		if (*arr)[j] < (*arr)[right] {
			i++
			swaps(i, j, arr)
		}
	}
	i++
	swaps(i, right, arr)
	return i
}

func swaps(a, b int, arr *[]int) {
	(*arr)[a], (*arr)[b] = (*arr)[b], (*arr)[a]
}
