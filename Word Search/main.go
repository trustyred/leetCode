package main

import "fmt"

func main() {
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	result := exist(board,"SEE")
	fmt.Println(result)
}
func exist(board [][]byte, word string) bool {
	rowLen := len(board)
	colLen := len(board[0])
	wordByteList := []byte(word)
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if backtrack(i,j,wordByteList,board,0){
				return true
			}
		}
	}
	return false

}

func backtrack(currentI int, currentJ int, remainWordByteList []byte, board [][]byte, num int) bool {
	if currentI >= len(board)||currentI<0||currentJ>=len(board[0])||currentJ<0{
		return false
	}
	if board[currentI][currentJ] != remainWordByteList[num] {
		return false
	}
	if len(remainWordByteList)-1 == num {
		return true
	}
	state := false
	board[currentI][currentJ] ^= 255
	state = backtrack(currentI+1, currentJ, remainWordByteList, board, num+1)||backtrack(currentI, currentJ+1, remainWordByteList, board, num+1) ||backtrack(currentI-1, currentJ, remainWordByteList, board, num+1) ||backtrack(currentI, currentJ-1, remainWordByteList, board, num+1)
	board[currentI][currentJ] ^= 255
	return state
}
