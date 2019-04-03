package main

import "fmt"

func main() {
	matrix := [][]int{{1,1,1},{1,0,2},{1,2,1}}
	setZeroes(matrix)
	fmt.Println(matrix)

}
func setZeroes(matrix [][]int)  {
	row0 := 1
	col0 := 1

	for i:=0;i<len(matrix);i++{
		if matrix[i][0] == 0{
			col0 = 0
		}
	}
	for i:=0;i<len(matrix[0]);i++{
		if matrix[0][i] == 0{
			row0 = 0
		}
	}


	for i:=1;i<len(matrix);i++{
		for j:=1;j<len(matrix[0]);j++{
			if matrix[i][j] == 0{
				matrix[i][0] =0
				matrix[0][j] =0
			}
		}
	}
	//fmt.Println(matrix)
	// 看每行起点，如果每行第一个为0那就清空这一行，但不判断第一行
	for i:=1;i<len(matrix);i++{
		if matrix[i][0] == 0{
			for j:=0;j<len(matrix[0]);j++{
				matrix[i][j] = 0
			}
		}
	}
	//fmt.Println(matrix)
	//看没列起点，如果为0，就把这一列设置为0,但不判断第一列
	for i:=1;i<len(matrix[0]);i++{
		if matrix[0][i] == 0{
			for j:=0;j<len(matrix);j++{
				matrix[j][i] = 0
			}
		}
	}
	if row0 == 0 {
		for i:=0;i<len(matrix[0]);i++{
			matrix[0][i] = 0
		}
	}
	if col0 == 0 {
		for i:=0;i<len(matrix);i++{
			matrix[i][0] = 0
		}
	}

}
