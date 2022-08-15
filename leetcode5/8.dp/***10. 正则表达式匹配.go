func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	if n == 0 { //IM: 1
		if m == 0 {
			return true
		} else {
			return false
		}
	}
	if m == 0 && n == 1 { //IM:2
		return false
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true           //IM:3
	for j := 2; j <= n; j++ { //IM:4
		if p[j-1] == '*' { // im ''
			dp[0][j] = dp[0][j-2] //im : -2
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i-1][j] //(不匹配（0个），匹配)
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[m][n]
}

/*

dp[i][j]表示两个指针， 求dp[m][n]
base: 1. n==0 ... m=0 n=1 false
      2.dp[0][j] = dp[0][j-2]   if p[j-1] == '*'
     *** 3.dp[0][0] = true

**保证每次出现字符 * 时，前面都匹配到有效的字符**

1.s[i-1]==p[j-1] || p[j-1] == '.': dp[i][j] = dp[i-1][j-1]
2.p[j-1] == '*' :
 2.1: s[i-1]==p[j-2] || p[j-2] == '.' (匹配0个或多个*): dp[i][j]= dp[i][j-2] || dp[i-1][j]
 2.3: else: dp[i][j] = dp[i][j-2]
3.else: dp[i][j] = false
*/