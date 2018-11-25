package bruteforce

import (
	"math"
	"sort"
)

type data struct {
	Cnt     int64
	Pattern *[][]bool
}

func Calc(board [][]bool, rowNum, colNum int64) data {
	result := data{
		Cnt: -1,
	}

	/* cols */
	total := rowNum * colNum
	/* pattern */
	times := int64(math.Pow(2, float64(total)))
	dataStream := make(chan data, times)

	results := []data{}
	for i := int64(0); i < times; i += 1 {
		go force(board, rowNum, colNum, i, dataStream)
	}
	for i := int64(0); i < times; i += 1 {
		select {
		case data := <-dataStream:
			if data.Cnt < 0 {
				continue
			}
			results = append(results, data)
		}
	}

	if len(results) > 0 {
		sort.Slice(results, func(i, j int) bool {
			return results[i].Cnt < results[j].Cnt
		})
		result = results[0]
	}
	return result
}
func newPattern(rowNum, colNum, seed int64) *[][]bool {
	pattern := make([][]bool, rowNum)
	for y := int64(0); y < rowNum; y += 1 {
		row := make([]bool, colNum)
		for x := int64(0); x < colNum; x += 1 {
			switch seed & 1 {
			case 1:
				row[x] = true
			default:
				row[x] = false
			}
			seed = seed >> 1
		}
		pattern[y] = row
	}
	return &pattern
}

func force(origin [][]bool, rowNum, colNum, seed int64, stream chan<- data) {
	board := make([][]bool, len(origin))
	for i := range origin {
		board[i] = make([]bool, len(origin[i]))
		copy(board[i], origin[i])
	}
	pattern := newPattern(rowNum, colNum, seed)

	for y := int64(0); y < rowNum; y += 1 {
		for x := int64(0); x < colNum; x += 1 {
			if (*pattern)[y][x] {
				change(&board, x, y, colNum, rowNum)
			}
		}
	}

	var cnt int64
	for y := int64(0); y < rowNum; y += 1 {
		for x := int64(0); x < colNum; x += 1 {
			if board[y][x] {
				stream <- data{
					Cnt:     -1,
					Pattern: pattern,
				}
				return
			}
			if (*pattern)[y][x] {
				cnt += 1
			}
		}
	}
	stream <- data{
		Cnt:     cnt,
		Pattern: pattern,
	}
	return
}
func change(board *[][]bool, x, y, xNum, yNum int64) {
	(*board)[y][x] = !(*board)[y][x]
	if x-1 >= 0 {
		(*board)[y][x-1] = !(*board)[y][x-1]
	}
	if x+1 < xNum {
		(*board)[y][x+1] = !(*board)[y][x+1]
	}
	if y-1 >= 0 {
		(*board)[y-1][x] = !(*board)[y-1][x]
	}
	if y+1 < yNum {
		(*board)[y+1][x] = !(*board)[y+1][x]
	}
}
