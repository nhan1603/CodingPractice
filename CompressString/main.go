package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
}

func compress(chars []byte) int {
	var tempChar byte
	var count int
	var buffer bytes.Buffer
	for _, char := range chars {
		if char != tempChar {
			if count > 1 {
				buffer.WriteString(strconv.Itoa(count))
			}
			buffer.WriteByte(char)
			tempChar = char
			count = 1
		} else {
			count = count + 1
		}
	}
	if count > 1 {
		buffer.WriteString(strconv.Itoa(count))
	}
	compressed := buffer.Bytes()
	copy(chars, compressed)
	return len(compressed)
}
