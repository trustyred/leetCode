package main

import "fmt"

func main() {
	result := spiralOrder([][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}})
	fmt.Println(result)
}

func spiralOrder(matrix [][]int) []int {
	rBegin := 0
	cBegin := 0
	if len(matrix) == 0{
		return []int{}
	}
	rEnd := len(matrix) - 1
	cEnd := len(matrix[0]) - 1
	var result []int
	var i int
	for rBegin<=rEnd && cBegin<=cEnd{
		for i=cBegin;i<=cEnd;i++{
			result = append(result, matrix[rBegin][i])
		}
		rBegin ++
		for i=rBegin;i<=rEnd;i++{
			result = append(result, matrix[i][cEnd])
		}
		cEnd --
		if  rBegin <= rEnd{
			for i=cEnd;i>=cBegin;i--{
				result = append(result, matrix[rEnd][i])
			}
		}

		rEnd --
		if cBegin <= cEnd{
			for i=rEnd;i>=rBegin;i--{
				result = append(result, matrix[i][cBegin])
			}
		}

		cBegin ++
	}

	return result

}
