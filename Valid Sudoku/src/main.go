package main

import (
	"fmt"
)
type Set struct {
	Map map[interface{}]int
}

func main() {
	m := map[int]int{}
	m[1] = 2
	_,ok := m[1]
	if ok{
		fmt.Println("have")
	}else{
		fmt.Println("no")
	}
}
//进行编码操作
//行数据编码成()line_number
//列数据编码成column_number()
//3*3方块内数据编码成(x)(y)
func isValidSudoku(board [][]byte) bool {

	var sudokuMap map[string]Set

	for i:= 0;i<9;i++{
		for j:=0 ;j<9;j++{
			lineIndex := fmt.Sprintf("()%d",i)
			colIndex := fmt.Sprintf("%d()",j)
			localIndex := fmt.Sprintf("(%d)(%d)",i/3,j/3)

			set1,ok1 := sudokuMap[lineIndex]
			set2,ok2 := sudokuMap[colIndex]
			set3,ok3 := sudokuMap[localIndex]
			if ok1{
				set1.Add()
			}
		}
	}
	return true
}

func (set *Set) Add(elem interface{}) bool{
	_ , ok := set.Map[elem]
	if ok{
		return false
	}else{
		set.Map[elem] = 1
		return true
	}
}
