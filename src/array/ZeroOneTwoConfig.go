package array

func ZeroOneTwoConfig(arr *[]int) {
	len := len(*arr)
	if *arr == nil || len < 2 {
		return
	}
	left := -1
	right := len
	index := 0
	for {
		if index >= right {
			break
		} else {
			if (*arr)[index] == 0 {
				left++
				SwapConfig(index, left, arr)
				index++
			} else if (*arr)[index] == 2 {
				right--
				SwapConfig(index, right, arr)
			} else {
				index++
			}
		}
	}
}

func SwapConfig(i int, j int, arr *[]int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
