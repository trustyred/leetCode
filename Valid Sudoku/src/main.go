package main

import (
	"fmt"
)
type Set struct {
	Map map[interface{}]int
}

func main() {
	fmt.Printf("%d\n",'.')
	sudo := [][]byte{
		{'.','8','7','6','5','4','3','2','1'},
		{'2','.','.','.','.','.','.','.','.'},
		{'3','.','.','.','.','.','.','.','.'},
		{'4','.','.','.','.','.','.','.','.'},
		{'5','.','.','.','.','.','.','.','.'},
		{'6','.','.','.','.','.','.','.','.'},
		{'7','.','.','.','.','.','.','.','.'},
		{'8','.','.','.','.','.','.','.','.'},
		{'9','.','.','.','.','.','.','.','.'},
	}
	//sudo := [][]byte{
	//	{'8','3','.','.','7','.','.','.','.'},
	//	{'6','.','.','1','9','5','.','.','.'},
	//	{'.','9','8','.','.','.','.','6','.'},
	//	{'8','.','.','.','6','.','.','.','3'},
	//	{'4','.','.','8','.','3','.','.','1'},
	//	{'7','.','.','.','2','.','.','.','6'},
	//	{'.','6','.','.','.','.','2','8','.'},
	//	{'.','.','.','4','1','9','.','.','5'},
	//	{'.','.','.','.','8','.','.','7','9'},
	//}
	result := isValidSudoku(sudo)
	fmt.Println(result)
}
//进行编码操作
//行数据编码成()line_number
//列数据编码成column_number()
//3*3方块内数据编码成(x)(y)
func isValidSudoku(board [][]byte) bool {

	sudokuMap := make(map[string]*Set)

	for i:= 0;i<9;i++{
		for j:=0 ;j<9;j++{
			lineIndex := fmt.Sprintf("()%d",i)
			colIndex := fmt.Sprintf("%d()",j)
			localIndex := fmt.Sprintf("(%d)(%d)",i/3,j/3)

			set1,ok1 := sudokuMap[lineIndex]
			set2,ok2 := sudokuMap[colIndex]
			set3,ok3 := sudokuMap[localIndex]
			if ok1{

				ret := set1.Add(board[i][j])
				if !ret && board[i][j]!='.'{
					return false
				}
			}else{
				var s *Set
				s = New()
				sudokuMap[lineIndex] = s
				s.Add(board[i][j])

			}
			if ok2 {

				ret := set2.Add(board[i][j])
				if !ret && board[i][j]!='.'{
					return false
				}
			}else{
				var s *Set
				s = New()
				sudokuMap[colIndex] = s
				s.Add(board[i][j])
			}
			if ok3 {
				ret := set3.Add(board[i][j])
				if !ret && board[i][j]!='.'{
					return false
				}
			}else{
				var s *Set
				s = New()
				sudokuMap[localIndex] = s
				s.Add(board[i][j])

			}

		}
	}
	//fmt.Println(sudokuMap)
	//flagCount := 0
	//for _,v:= range sudokuMap{
	//	if len(v.Map) == 1{
	//		flagCount ++
	//	}
	//}
	//if flagCount == 27||flagCount==0{
	//	return true
	//}else{
	//	return false
	//}
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

func New() *Set{
	return &Set{
		Map: map[interface{}]int{},
	}
}