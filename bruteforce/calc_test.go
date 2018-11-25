package bruteforce

import (
	"fmt"
)

func ExampleCalc() {
	rowNum := int64(3)
	colNum := int64(3)

	board := [][]bool{
		{false, true, false},
		{true, false, true},
		{false, true, false},
	}
	data := Calc(board, rowNum, colNum)
	fmt.Println(data.Cnt, data.Pattern)

	board = [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}
	data = Calc(board, rowNum, colNum)
	fmt.Println(data.Cnt, data.Pattern)
	// Output:
	// 4 &[[false true false] [true false true] [false true false]]
	// 0 &[[false false false] [false false false] [false false false]]
}
