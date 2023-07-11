package main

func main() {

}

func longestPalindrome(words []string) int {
	length := len(words)
	if length == 1 {
		if words[0] == reverse(words[0]) {
			return len(words[0])
		} else {
			return 0
		}
	} else {

		return 0
	}
}

func reverse(s string) string {
	var result string
	length := len(s) - 1
	for i := length; i >= 0; i++ {
		result += string(s[i])
	}
	return result
}
