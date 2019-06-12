package string

import (
	"container/list"
)

func LongestParenthesis(str string) int {
	len := len(str)
	dp := make([]int, len)
	result := 0
	for i := range str {
		if i == 0 {
			continue
		}
		if str[i] == ')' {
			pre := i - dp[i-1] - 1
			if str[pre] == '(' {
				if pre > 0 {
					dp[i] = dp[i-1] + 2 + dp[pre-1]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func LongestParenthesisUpdate(str string) int {
	result := 0
	last := -1
	list := list.New()
	for i := range str {
		if str[i] == '(' {
			list.PushBack(i)
		} else {
			if list.Len() == 0 {
				last = i
			} else {
				element := list.Back()
				list.Remove(element)
				if list.Len() == 0 {
					if result < i-last {
						result = i - last
					}
				} else {
					if result < i-list.Back().Value.(int) {
						result = i - list.Back().Value.(int)
					}
				}
			}
		}
	}
	return result
}
