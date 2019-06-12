package array

import (
	"container/list"
	"fmt"
)

func ThreeSum(arr *[]int, target int) list.List {
	result := list.New()
	len := len(*arr)
	shellSort(arr)
	fmt.Println(arr)
	for i := 0; i < len-2; i++ {
		left := i + 1
		right := len - 1
		num := target - (*arr)[i]
		for {
			if left > right {
				break
			}
			if (*arr)[left]+(*arr)[right] == num {
				list := list.New()
				list.PushBack((*arr)[left])
				list.PushBack((*arr)[right])
				list.PushBack((*arr)[i])
				for ele := list.Front(); ele != nil; ele = ele.Next() {
					fmt.Print(ele.Value)
				}
				fmt.Println()
				result.PushBack(list)
				for {
					if (*arr)[left+1] != (*arr)[left] {
						break
					} else {
						left += 1
					}
				}
				for {
					if (*arr)[right-1] != (*arr)[right] {
						break
					} else {
						right -= 1
					}
				}
				left++
				right--
			} else if (*arr)[left]+(*arr)[right] < num {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return *result
}
