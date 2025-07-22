package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	fmt.Println(closeStrings("test", "stee"))
}

func closeStrings(word1 string, word2 string) bool {
	str1 := make(map[rune]int)
	for _, char := range word1 {
		val, ok := str1[char]
		if ok {
			str1[char] = val + 1
		} else {
			str1[char] = 1
		}
	}

	str2 := make(map[rune]int)
	for _, char := range word2 {
		val, ok := str2[char]
		if ok {
			str2[char] = val + 1
		} else {
			str2[char] = 1
		}
	}

	var key1, key2 []rune
	var val1, val2 []int

	for key, val := range str1 {
		key1 = append(key1, key)
		val1 = append(val1, val)
	}

	for key, val := range str2 {
		key2 = append(key2, key)
		val2 = append(val2, val)
	}
	sort.Ints(val1)
	sort.Ints(val2)

	return SortString(string(key1)) == SortString(string(key2)) && reflect.DeepEqual(val1, val2)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
