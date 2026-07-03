package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	result := 0
	for _, v := range cb[file] {
		if v {
			result += 1
		}
	}
	return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	files := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	result := 0
	for _, file := range files {
		if cb[file][rank-1] {
			result += 1
		}
	}
	return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	result := 0
	for _, v := range cb {
		for _ = range v {
			result += 1
		}
	}
	return result
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	result := 0
	for _, files := range cb {
		for _, file := range files {
			if file {
				result += 1
			}
		}
	}
	return result
}
