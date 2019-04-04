package main

import "fmt"

func main() {

	matrix := [][]int{{}}
	result := searchMatrix(matrix,5)
	fmt.Println(result)
}
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 ||len(matrix[0]) == 0{
		return false
	}
	rowLen := len(matrix)
	colLen := len(matrix[0])
	i := 0
	j := rowLen*colLen - 1
	for i<j {
		mid := (i+j)/2
		if target < matrix[mid/colLen][mid%colLen]{
			j = mid
		}else if target > matrix[mid/colLen][mid%colLen]{
			i++
		}else{
			return true
		}
	}
	if target == matrix[(rowLen*colLen - 1)/colLen][(rowLen*colLen - 1)%colLen]{
		return true
	}
	return false
}
