package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
}

func equalPairs(grid [][]int) int {
	size := len(grid)
	columns := make(map[int][]int, size)

	for _, row := range grid {
		for indx, val := range row {
			columns[indx] = append(columns[indx], val)
		}
	}

	var result int

	for _, row := range grid {
		for _, column := range columns {
			if reflect.DeepEqual(row, column) {
				result++
			}
		}
	}

	return result
}

func equalPairs2(grid [][]int) int {
	cols := make(map[string]int)
	rows := make(map[string]int)
	n := len(grid) // n x n matrix
	for i := 0; i < n; i++ {
		var row, col strings.Builder
		for j := 0; j < n; j++ {
			row.WriteString(fmt.Sprintf("%d;", grid[i][j]))
			col.WriteString(fmt.Sprintf("%d;", grid[j][i]))
		}
		cols[col.String()]++
		rows[row.String()]++
	}
	var res int
	for col, count1 := range cols {
		if count2, ok := rows[col]; ok {
			// We are not finding unique pairs; If we have 2 rows and 2 columns; we can make
			// 2 * 2 = 4 pairs
			res += count1 * count2
		}
	}
	return res
}
