// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

func isPalindrome(s string) bool {
	str, rev := "", ""
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			str = str + string(ch)
			rev = string(ch) + rev
		}
	}

	if strings.ToLower(str) != strings.ToLower(rev) {
		return false
	}
	return true
}