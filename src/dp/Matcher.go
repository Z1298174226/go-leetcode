package dp

func IsMatcher(s, p string) bool {
	len1 := len(s)
	len2 := len(p)
	dp := create(len1+1, len2+1)
	dp[0][0] = true
	for i := 1; i <= len2; i++ {
		if p[i-1] == '*' && i&1 == 0 && i > 1 && dp[0][i-1] {
			dp[0][i] = true
		}
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				if p[j-2] != s[i-1] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i][j-2]
				}
			}
		}
	}
	return dp[len1][len2]

}

type Matrix [][]bool

func create(rows, cols int) Matrix {
	matrix := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]bool, cols)
	}
	return matrix
}
