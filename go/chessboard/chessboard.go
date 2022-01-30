package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [8]bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	total := 0
	for _, occupied := range cb[rank] {
		if occupied {
			total++
		}
	}

	return total
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	total := 0

	if file < 1 || file > 8 {
		return 0
	}

	for _, ranks := range cb {
		if ranks[file-1] {
			total++
		}
	}

	return total
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return 64
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	total := 0
	for rank := range cb {
		total += CountInRank(cb, rank)
	}

	return total
}
