package main

import "fmt"

func main() {
	result := uniquePaths(2,3)
	fmt.Println(result)
}
func uniquePaths(m int, n int) int {
	dp := make([][]int,n+1)
	for i:=0;i<n+1;i++{
		dp[i] = make([]int,m+1)
	}

	for i:=1;i<=n;i++{
		for j:=1;j<=m;j++{
			if i==1 || j==1{
				dp[i][j] = 1
			}else{
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}
func uniquePathsBacktrack(m int, n int) int {
	pathNum :=0
	backtrack(1,1,m,n,&pathNum)
	return pathNum
}
/* direction = 1 is right direction = 2 is down*/
func backtrack(currentX int, currentY int, m int,n  int,pathNum *int){
	if currentX >m||currentY>n{
		return
	}else if currentX == m && currentY == n{
		*pathNum++
		return
	}
	directionList := []int{1,2}
	for i:=0;i<len(directionList);i++{
		if i==0{
			backtrack(currentX+1,currentY,m,n,pathNum)
		}else if i ==1{
			backtrack(currentX,currentY+1,m,n,pathNum)
		}
	}
}