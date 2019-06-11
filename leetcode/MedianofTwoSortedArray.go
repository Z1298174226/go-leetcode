package main

import "fmt"

func medianOfTwo(arr1, arr2 []int) float64 {
	len1 := len(arr1)
	len2 := len(arr2)
	len := len1 + len2
	if (len & 1) == 1 {
		result := help(arr1, arr2, len1, len2, 0, 0, len/2+1)
		return float64(result)
	} else {
		result := help(arr1, arr2, len1, len2, 0, 0, len/2+1) +
			help(arr1, arr2, len1, len2, 0, 0, len/2)
		return float64(result) / 2
	}

}

func help(arr1, arr2 []int, len1, len2 int, index1, index2 int, k int) int {
	if len1 > len2 {
		return help(arr2, arr1, len2, len1, index2, index1, k)
	} else {
		if k == 1 {
			if arr1[index1] > arr2[index2] {
				return arr2[index2]
			} else {
				return arr1[index1]
			}
		}
		if index1 == len1 {
			return arr2[index2+k]
		}
		p := k >> 1
		q := k - p
		if arr1[index1+p] < arr2[index2+q] {
			return help(arr1, arr2, len1, len2, index1+p, index2, k-p)
		} else {
			return help(arr1, arr2, len1, len2, index1, index2+q, k-q)
		}
	}
}

func main() {
	a := [...]int{1, 3, 5, 7, 9, 13}
	b := [...]int{2, 4, 6, 8, 16, 18}
	fmt.Println(medianOfTwo(a[:], b[:]))
}
