package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	dummyDict = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
)

func Anagram() {
	list := make(map[string][]string)
	result := make([][]string , 0)

	for _, word := range dummyDict {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		wordsArr := []string{}
		for _, w := range words {
			wordsArr = append(wordsArr, w)
		}
		result = append(result, wordsArr)
	}
	fmt.Print(result)
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	Anagram()
}
