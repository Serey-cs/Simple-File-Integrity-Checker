package main

func encryptRailFence(text string, key int) string {
	if key <= 1 {
		return text
	}

	rails := make([][]rune, key)
	row, dir := 0, 1

	for _, ch := range text {
		rails[row] = append(rails[row], ch)

		if row == 0 {
			dir = 1
		} else if row == key-1 {
			dir = -1
		}
		row += dir
	}

	result := ""
	for _, rail := range rails {
		result += string(rail)
	}
	return result
}
