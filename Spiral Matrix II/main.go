package main

import "fmt"

func main() {
	result := generateMatrix(0)
	fmt.Println(result)
}
func generateMatrix(n int) [][]int {

	direction := 0
	//0 is right 1 is down 2 is left 3 is up
	row := 0
	col := 0
	result := make([][]int,n)
	for i:=0;i<len(result);i++{
		result[i] = make([]int,n)
	}

	for i:=1;i<=n*n;i++{
		result[row][col] = i
		if direction == 0{
			col ++
			if col>=n || result[row][col] !=0{
				col --
				row ++
				direction = 1
			}
		}else if direction == 1{
			row++
			if row >=n || result[row][col] !=0 {
				row --
				col --
				direction = 2
			}
		}else if direction == 2 {
			col--
			if col <0 || result[row][col] !=0{
				col++
				row--
				direction =3
			}
		}else if direction == 3{
			row --
			if row <0 || result[row][col]!=0{
				row ++
				col ++
				direction = 0
			}
		}
	}
	return  result
}