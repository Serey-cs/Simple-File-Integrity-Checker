package main

func decryptRailFence(cipher string, key int) string {
	if key <= 1 {
		return cipher
	}

	n := len(cipher)
	matrix := make([][]rune, key)
	for i := range matrix {
		matrix[i] = make([]rune, n)
	}

	// Mark zigzag positions
	row, dir := 0, 1
	for col := 0; col < n; col++ {
		matrix[row][col] = '*'

		if row == 0 {
			dir = 1
		} else if row == key-1 {
			dir = -1
		}
		row += dir
	}

	// Fill the cipher text
	index := 0
	for i := 0; i < key; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '*' && index < n {
				matrix[i][j] = rune(cipher[index])
				index++
			}
		}
	}

	// Read plaintext
	result := ""
	row, dir = 0, 1
	for col := 0; col < n; col++ {
		result += string(matrix[row][col])

		if row == 0 {
			dir = 1
		} else if row == key-1 {
			dir = -1
		}
		row += dir
	}

	return result
}
