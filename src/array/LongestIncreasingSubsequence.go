package array

func LongestIncreasingSubsequence(arr *[]int) int {
	len := len(*arr)
	if len < 2 {
		return len
	}
	dp := make([]int, len+1)
	dp[1] = (*arr)[0]
	result := 1
	for i := 1; i < len; i++ {
		pos := help(&dp, 1, result, (*arr)[i])
		dp[pos] = (*arr)[i]
		result = max(result, pos)
	}
	return result
}

func help(dp *[]int, start, end int, value int) int {
	if (*dp)[end] < value {
		return end + 1
	}
	for {
		mid := start + (end-start)/2
		if (*dp)[mid] < value {
			start = mid + 1
		} else {
			end = mid
		}
		if start >= end {
			break
		}
	}
	return start

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
