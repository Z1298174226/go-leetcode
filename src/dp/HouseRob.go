package dp

func MaxRob(arr *[]int) int {
	len := len(*arr)
	if len == 0 {
		return 0
	}
	preYes := (*arr)[0]
	preNo := 0
	temp := 0
	for i := 1; i < len; i++ {
		temp = preNo
		preNo = max(preYes, preNo)
		preYes = max(temp+(*arr)[i], preYes)
	}
	return max(preYes, preNo)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
