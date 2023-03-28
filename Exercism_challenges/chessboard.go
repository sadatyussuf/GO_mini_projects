package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int { 
    c := 0
	for _,V := range cb[file]{
        if V{
            c++
        }
    }
     return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank <= 0 || rank > 8 {
		return 0
	}
	c := 0
	for _,V := range cb{
       if V[rank-1]{
           c++
       }
    }
     return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	c := len(cb)
    i := len(cb["A"])
     return c*i
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    c := 0
	for _,V := range cb{
       for _,i := range V { 
           if i{
               c++
             }
       }   
    }
return c
}
