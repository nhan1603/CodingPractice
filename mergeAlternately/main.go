package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeAlternately("oh", "2222"))
}

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder

	var doBackTrack func(word1 string, word2 string)
	doBackTrack = func(word1 string, word2 string) {
		if len(word1) == 0 {
			return
		}
		if len(word1) > 0 {
			sb.WriteByte(word1[0])
		}
		word1 = word1[1:]
		if len(word2) > 0 {
			doBackTrack(word2, word1)
		} else {
			doBackTrack(word1, word2)
		}
	}

	if len(word1) == 0 {
		word1, word2 = word2, word1
	}
	doBackTrack(word1, word2)

	return sb.String()
}
