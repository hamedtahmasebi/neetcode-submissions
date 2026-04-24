func isPalindrome(s string) bool {
	stack := make([]byte, 0, len(s))

	// in order to simply the code, if array length is odd, we will remove the middle of the array
	reg := regexp.MustCompile(`(?i)[^a-zA-Z0-9]`)
	str := strings.ToLower(reg.ReplaceAllString(s, ""))

	if len(str)%2 != 0 {
		mid := len(str) / 2
		str = strings.Join([]string{str[0:mid], str[mid+1:]}, "")
	}

	mid := len(str) / 2
	for i := 0; i < mid; i++ {
		stack = append(stack, str[i])
	}
	for i := mid; i < len(str); i++ {
		if str[i] != stack[mid-(i-mid)-1] {
			return false
		}
	}
	return true
}
