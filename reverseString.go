package main

func reverseString(s []byte) {

	z := make([]byte, len(s))
	j := 0

	for i := len(s) - 1; i >= 0; i-- {
		z[j] = s[i]
		j++
	}

	copy(s, z)

}

func reverseString2(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

}
