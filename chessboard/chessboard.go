package chessboard

type Rank []bool

type Chessboard map[string]Rank

func CountInRank(cb Chessboard, rank string) int {
	count := 0

	for _, val := range cb[rank] {
		if val {
			count++
		}
	}

	return count
}

func CountInFile(cb Chessboard, file int) int {
	count := 0

	if file < 1 || file > len(cb) {
		return count
	}

	for _, rank := range cb {
		if rank[file - 1] {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0

	for _, rank := range cb {
		for range rank {
			count++
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
		count := 0

	for _, rank := range cb {
		for _, isOccupied := range rank {
			if isOccupied {
				count++
			}
		}
	}

	return count
}
