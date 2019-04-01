package main

import "fmt"

func main() {
	result := climbStairs(5)
	fmt.Println(result)

}

func climbStairs(n int) int {
	dp := make([]int,n+2)
	dp[1] = 1
	dp[2] = 2
	for i:=3;i<=n;i++{
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}