package main

import (
	"fmt"
	"math"
)

func main() {
	var grid [][]int
	grid = [][]int{{1},{1},{4}}

	result := minPathSum(grid)
	fmt.Println(result)


}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make([][]int,len(grid)+1)
	for i:=0;i<len(grid)+1;i++{
		dp[i] = make([]int,len(grid[0])+1)
	}
	max := 4294967295
	for i:=0;i<len(dp);i++{
		for j:=0;j<len(dp[0]);j++{
			if i==0||j==0{
				dp[i][j] = max
			}else if i==1 && j==1{
				dp[i][j] = grid[0][0]
			}else{
				dp[i][j] = grid[i-1][j-1] + int(math.Min(float64(dp[i-1][j]),float64(dp[i][j-1])))
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]

}