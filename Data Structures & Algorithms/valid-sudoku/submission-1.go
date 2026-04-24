func isValidSudoku(board [][]byte) bool {
	boxes := make([][]byte, 9)
	// row validation
	for r, row := range board {
		rexistMap := make(map[byte]bool)
		for c, cell := range row {
			boxes[getBoxIdx(r, c)] = append(boxes[getBoxIdx(r, c)], cell)
			if cell == '.' {
				continue
			}
			_, exists := rexistMap[cell]
			if exists {
				return false
			}
			rexistMap[cell] = true
		}
	}

	// col validation
	for c := 0; c < len(board); c++ {
		existMap := make(map[byte]bool)
		for r := 0; r < len(board); r++ {
			cell := board[r][c]
			if cell == '.' {
				continue
			}
			_, exists := existMap[cell]
			if exists {
				return false
			}
			existMap[cell] = true
		}
	}

	// box validation
	for _, b := range boxes {
		existMap := make(map[byte]bool)
		for _, cell := range b {
			if cell == '.' {
				continue
			}
			_, exists := existMap[cell]
			if exists {
				return false
			}
			existMap[cell] = true
		}
	}

	return true
}

func getBoxIdx(row, col int) int {
	return (row/3)*3 + (col / 3)
}