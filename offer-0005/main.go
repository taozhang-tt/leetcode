package main

import (
	"strings"
)

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func replaceSpace2(s string) string {
	ret, cnt := make([]rune, 0, len(s)*3), 0
	for _, char := range []rune(s) {
		if char == ' ' {
			ret = append(ret, '%', '2', '0')
			cnt++
		} else {
			ret = append(ret, char)
		}
	}

	return string(ret[:len(s)+cnt*2])
}
