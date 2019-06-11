package main

import "fmt"

func main() {
	f := uniqueBinaryTree
	fmt.Println(f(4))
	fmt.Print(uniqueBinaryTree(3))
}

func uniqueBinaryTree(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
