package boardstuff

/**
 * Board used to simulate a Chess Board that contain Queen as much as Size at Queens
 * represent as a solution that would be collected to MainBoard.
 */
type Board struct {
	Size   int
	Queens []Queen
}

/**
 * QueenIsFull is a Board function to check is Queens are already full
 * @return true if Queens have length same as the Size, false if its not the same amount.
 */
func (b *Board) QueenIsFull() bool {
	return b.Size == len(b.Queens)
}

/**
 * AddQueen is a Board function to add more Queen to the Queens
 * homeWrecker should not be added to the Queens if the length of it exceed Size
 * @param homewrecker is the Queen that would be added
 */
func (b *Board) AddQueen(homewrecker Queen) {
	if !b.QueenIsFull() {
		b.Queens = append(b.Queens, homewrecker)
	}
}

/**
 * AbleToDate is a Board function to check if a Queen could be added to the Queens Board
 * @param homewrecker is the Queen that would be checked with the current Queens
 * @return true if homewrecker fits the Queens Board, false if there will be madness if the homewrecker added to the Queens Board.
 */
func (b *Board) AbleToDate(homewrecker Queen) bool {
	for _, queen := range b.Queens {
		if homewrecker.MadnessPotential(queen) {
			return false
		}
	}
	return true
}

/**
 * QueensAreAllTheSame is a Board function to check if the Board's Queens are the same with the dopple Board's Queens
 * @param doppleBoard is the Board to be compared
 * @return true if dopple Board's Queens are the doppleganger of Board's Queens, false if one or more of them are not a doppleganger.
 */
func (b *Board) QueensAreAllTheSame(doppleBoard Board) bool {
	myQueen := b.Queens
	queens := doppleBoard.Queens
	if len(myQueen) != len(queens) {
		return false
	}
	for i, queen := range queens {
		if myQueen[i].X != queen.X || myQueen[i].Y != queen.Y {
			return false
		}
	}
	return true
}
