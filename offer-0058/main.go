package main

// 暴力解法
func reverseLeftWords(s string, n int) string {
	chars := []byte(s)
	for i := 0; i < n; i++ {
		temp, j := chars[0], 0
		for j < len(chars)-1 {
			chars[j], j = chars[j+1], j+1
		}
		chars[j] = temp
	}

	return string(chars)
}

func reverseLeftWords2(s string, n int) string {
	return s[n:] + s[:n]
}
