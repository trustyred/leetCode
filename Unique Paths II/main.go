package main

import (
	"fmt"
)

func main() {
	result :=uniquePathsWithObstacles([][]int{{1,0}})
	fmt.Println(result)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0{
		return 0
	}
	dp := make([][]int,len(obstacleGrid)+1)
	for i:=0;i<len(obstacleGrid)+1;i++{
		dp[i] = make([]int,len(obstacleGrid[0])+1)
	}
	if obstacleGrid[0][0] == 1{
		dp[1][1] = 0
	}else{
		dp[1][1] = 1
	}
	for i:=1;i<=len(obstacleGrid);i++{
		for j:=1;j<=len(obstacleGrid[0]);j++{
			if i==1&&j==1{
				continue
			}
			if obstacleGrid[i-1][j-1] == 1{
				dp[i][j] = 0
			}else{
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)][len(obstacleGrid[0])]
}