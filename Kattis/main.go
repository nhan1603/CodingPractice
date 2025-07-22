package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var list [][]int
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 4; i++ {
		line, _ := reader.ReadString('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		data := strings.Split(line, " ")

		var temp []int
		for indx := range data {
			val, err := strconv.Atoi(data[indx])
			if err != nil {
				val, _ = strconv.Atoi(data[indx][:len(data[indx])-1])
			}
			temp = append(temp, val)
		}
		list = append(list, temp)
	}
	var n int
	fmt.Scanln(&n)

	for _, val := range list {
		str := fmt.Sprint(val)
		fmt.Println(str[1 : len(str)-1])
	}

	switch n {
	// two pointer
	case 0:
		// left
		for _, row := range list {
			i := 0
			j := i + 1
			for i < 3 {
				for row[j] == 0 && j < 3 {
					j++
				}
				if row[i] == row[j] {
					row[i] = row[i] + row[j]
					for j < 3 {
						row[j] = row[j+1]
						j++
					}
					row[3] = 0
				}
				i++
				j = i + 1
			}
		}
	case 1:
		// up
	case 2:
		// right
		for _, row := range list {
			i := 3
			j := i - 1
			for i > 0 {
				for j > 0 && row[j] == 0 {
					j--
				}
				if row[i] == row[j] {
					row[i] = row[i] + row[j]
					for j > 0 {
						row[j] = row[j-1]
						j--
					}
					row[0] = 0
				}
				i--
				j = i - 1
			}
		}
	case 3:
		// down
	}

	for _, val := range list {
		str := fmt.Sprint(val)
		fmt.Println(str[1 : len(str)-1])
	}

}
