package main

func validPalindrome(s string) bool {

	i, j := 0, len(s)-1

	for i < j {
		if !isValid(s[i]) {
			i++
			continue
		}

		if !isValid(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true

}

func isValid(t byte) bool {
	if (t >= 'a' && t <= 'z') || (t >= 'A' && t <= 'Z') {
		return true
	}

	return false
}
