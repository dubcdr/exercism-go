package chessboard

import "fmt"

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0
	for _, v := range cb[rank] {
		if v {
			count += 1
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0
	if file < 0 || file > 8 {
		return 0
	}
	fileIndex := file - 1
	for k := range cb {
		if fileIndex < 0 || fileIndex > 7 {
			fmt.Println("whats going on")
		}
		temp := cb[k][fileIndex]

		if temp {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for i := range cb {
		for range cb[i] {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for i := range cb {
		for _, v := range cb[i] {
			if v {
				count++
			}
		}
	}
	return count
}
