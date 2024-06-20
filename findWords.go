package main

func findWords(words []string) []string {
	rows := [26]int{
		2, 3, 3, 2, 1, 2, 2, 2, 1, 2, 2, 2, 3,
		3, 1, 1, 1, 1, 2, 1, 1, 3, 1, 3, 1, 3,
	}

	var c int

	for _, w := range words {
		l := len(w)

		var invalid bool

		for i := 0; i < l-1; i++ {
			if rows[toIndex(w[i])] != rows[toIndex(w[i+1])] {
				invalid = true
				break
			}
		}

		if invalid {
			continue
		}

		words[c] = w
		c++
	}

	return words[:c]
}

func toIndex(r byte) byte {
	if r >= 'A' && r <= 'Z' {
		return r - 'A'
	}
	return r - 'a'
}
