package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    var sum int
        if v, ok := cb[file]; ok {
       		for i := 0; i < 8; i++ {
                if v[i] {
                    sum++
                }
            }
        }
	return sum
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank <= 0 || rank > 8 {
        return 0
    }
    var sum int
	for _, v := range cb {
        if v[rank-1] {
            sum++
        }
    }
    return sum
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var sum int
    for _, v := range cb {
       for range v {
           sum++
       }
    }
	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var sum int
    for _, v := range cb {
        for i := 0; i < 8; i++ {
            if v[i] {
                sum++
            }
        }
    }
	return sum
}
