package array

import "fmt"

func DivideGold(arr *[]int) int {
	inits(arr)
	fmt.Println(arr)
	len := len(*arr)
	ans := 0
	for {
		if len == 1 {
			break
		}
		a := (*arr)[0]
		fmt.Println(a)
		swap(1, len, arr)
		len--
		sink(1, len, arr)
		b := (*arr)[0]
		ans = ans + a + b
		(*arr)[0] = a + b
		sink(1, len, arr)
	}
	return ans
}

func inits(arr *[]int) {
	len := len(*arr)
	for i := len / 2; i >= 1; i-- {
		sink(i, len, arr)
	}

}

func sink(left, right int, arr *[]int) {
	harf := right >> 1
	for {
		if left > harf {
			break
		}
		child := left << 1
		if child < right && (*arr)[child-1] > (*arr)[child] {
			child++
		}
		if (*arr)[left-1] > (*arr)[child-1] {
			swap(left, child, arr)
		}
		left = child
	}
}

func swap(a, b int, arr *[]int) {
	(*arr)[a-1], (*arr)[b-1] = (*arr)[b-1], (*arr)[a-1]
}
