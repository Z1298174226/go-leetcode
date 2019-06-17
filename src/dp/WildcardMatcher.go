package dp

func Wildcard(s, p string) bool {
	len1 := len(s)
	len2 := len(p)
	dp := CreateMatrixBool(len1+1, len2+1)
	dp[0][0] = true
	for i := 1; i <= len2; i++ {
		if p[i-1] == '*' && dp[0][i-1] {
			dp[0][i] = true
		}
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i-1][j-1]
			}
		}
	}
	return dp[len1][len2]
}

type MatrixBool [][]bool

func CreateMatrixBool(row, col int) MatrixBool {
	matrix := make([][]bool, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]bool, col)
	}
	return matrix
}
