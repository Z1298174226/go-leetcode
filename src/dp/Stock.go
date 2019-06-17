package dp

type point *[]int

func StockI(arr point) int {
	buy := -10000
	sell := 0
	for i := range *arr {
		if buy < -(*arr)[i] {
			buy = -(*arr)[i]
		}
		if sell < buy+(*arr)[i] {
			sell = buy + (*arr)[i]
		}
	}
	return sell

}

func StockII(arr point) int {
	buyFirst := -10000
	sellFirst := 0
	buySecond := -10000
	sellSecond := 0
	for i := range *arr {
		if buyFirst < -(*arr)[i] {
			buyFirst = -(*arr)[i]
		}
		if sellFirst < buyFirst+(*arr)[i] {
			sellFirst = buyFirst + (*arr)[i]
		}
		if buySecond < -(*arr)[i]+sellFirst {
			buySecond = -(*arr)[i] + sellFirst
		}
		if sellSecond < buySecond+(*arr)[i] {
			sellSecond = buySecond + (*arr)[i]
		}
	}
	return sellSecond
}

func StockIII(arr point, fee int) int {
	len := len(*arr)
	if len == 0 {
		return 0
	}
	buy := make([]int, len)
	hold := make([]int, len)
	sell := make([]int, len)
	skip := make([]int, len)
	buy[0] = -(*arr)[0]
	hold[0] = -(*arr)[0]
	for i := 1; i < len; i++ {
		buy[i] = max(skip[i-1], sell[i-1]) - (*arr)[i]
		hold[i] = max(hold[i-1], buy[i-1])
		sell[i] = max(hold[i-1], buy[i-1]+(*arr)[i])
		skip[i] = max(skip[i-1], sell[i-1])
	}
	result := max(sell[len-1], skip[len-1])
	return max(result, 0)
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
