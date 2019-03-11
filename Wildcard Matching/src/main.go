package main

func main() {

}

func isMatch(s string, p string) bool {
	dp := [len(p)][len(s)]bool{}
	for i:=0;i<len(p);i++{
		for j:=0;j<len(s);j++{
			if j==0{
				if i==0 && (s[j]==p[i] || p[i]=='?'||p[i]=='*'){
					dp[i][j] = true
				}else if i==0{
					dp[i][j] = false
				}
				if i>0{
					if  dp[i-1][j] && p[i] == '*'{
						dp[i][j] = true
 					}else if  dp[i-1][j] && p[i-1] == '*'{
 						dp[i][j] = true
					}else if
						dp[i][j]
					}
				}

			}
		}
	}
}