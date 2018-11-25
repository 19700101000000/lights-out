package main

import (
	"fmt"
	"github.com/19700101000000/lights-out/bruteforce"
	"math/rand"
	"time"
)

func main() {
	var row, col int64
	var board [][]bool

	row = 3
	col = 3

	rand.Seed(time.Now().UnixNano())
	for y := int64(0); y < row; y += 1 {
		var row []bool
		for x := int64(0); x < col; x += 1 {
			var col bool
			switch rand.Intn(2) {
			case 1:
				col = true
			default:
				col = false
			}
			row = append(row, col)
		}
		board = append(board, row)
	}
	fmt.Println(board)
	result := bruteforce.Calc(board, row, col)
	if result.Cnt < 0 {
		fmt.Println("fail")
	} else {
		fmt.Println("%d::", result.Cnt)
		for _, row := range *result.Pattern {
			for _, col := range row {
				if col {
					fmt.Printf("o")
				} else {
					fmt.Printf("x")
				}
			}
			fmt.Println()
		}
	}
}
