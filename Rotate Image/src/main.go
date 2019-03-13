package main

import "fmt"

func main() {
	matrix := [][]int{{}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int)  {
	length := len(matrix)
	for i:=0;i<length;i++{
		for j:=i;j<length-i-1;j++{
			tmpI := j
			tmpJ := length - 1 -i
			for k:=0;k<3;k++{
				tmp := matrix[i][j]
				matrix[i][j] = matrix[tmpI][tmpJ]
				matrix[tmpI][tmpJ] = tmp
				save := tmpI
				tmpI = tmpJ
				tmpJ = length -1 - save
			}
		}
	}
}