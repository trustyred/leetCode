package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := numDecodings("226")

	fmt.Println(result)
}
func numDecodings(s string) int {
	dp := make([]int,len(s)+1)

	dp[len(s)] = 1
	if s[len(s)-1] =='0'{
		dp[len(s)-1] = 0
	}else{
		dp[len(s)-1] = 1
	}
	for i:=len(s)-2;i>=0;i--{
		if s[i] == '0'{
			continue
		}else{
			num,err := strconv.Atoi(s[i:i+2])
			if err==nil && num<=26{
				dp[i] = dp[i+1]+dp[i+2]
			}else if err==nil && num>26{
				dp[i] = dp[i+1]
			}
		}
	}
	return dp[0]
}
