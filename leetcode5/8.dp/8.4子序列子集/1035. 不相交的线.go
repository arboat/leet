func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ { //IM: <=
		for j := 1; j <= n; j++ { //IM: <=
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
1143.最长公共子序列
dp[i][j]
dp[i][j] = dp[i-1][j-1]+1 : nums1[i-1] == nums2[j-1]
         = max(dp[i-1][j], dp[i][j-1])
init :0
*/