func isPalindrome(x int) bool {
	intString := strconv.Itoa(x)
	for i := 0; i < len(intString); i++ {
		j := len(intString) - 1 - i
		if intString[i] != intString[j] {
			return false
		}
	}
	return true
}