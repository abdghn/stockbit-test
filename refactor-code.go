package main

import (
	"fmt"
	"regexp"
)

func findFirstStringInBracket(str string) string {
	if len(str) == 0 {
		return ""
	}

	re := regexp.MustCompile(`\((.*?)\)`)
	match := re.FindStringSubmatch(str)
	if len(match) < 1 {
	return ""
	}

	strings := match[1]
	result := []rune(strings)
	return string(result[0:1])
}

	func main () {
		fmt.Print(findFirstStringInBracket("t(est)"))
	}