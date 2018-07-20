package boardstuff_test

import (
	. "nqueen/pkg/boardstuff"
	"testing"
)

/**
 * TestBoard_AddQueen are testing a Board that adding a Queen to the Queens.
 */
func TestBoard_AddQueen(t *testing.T) {
	theBoard := Board{Size: 8}
	firstLady := Queen{X: 0, Y: 0}

	theBoard.AddQueen(firstLady)

	theFirstLady := theBoard.Queens[0]
	if theFirstLady.X != firstLady.X || theFirstLady.Y != firstLady.Y {
		t.Error("First Lady should be added to the board")
	}

	secondQueen := Queen{X: 5, Y: 4}
	theBoard.AddQueen(secondQueen)

	theSecondQueen := theBoard.Queens[1]
	if theSecondQueen.X != secondQueen.X || theSecondQueen.Y != secondQueen.Y {
		t.Error("Second Queen should be added to the board")
	}
}

/**
 * TestBoard_AddQueenAfterFull are testing adding Queen if the Queens are already full.
 */
func TestBoard_AddQueenAfterFull(t *testing.T) {
	theBoard := Board{Size: 8}
	firstLady := Queen{X: 0, Y: 1}
	secondQueen := Queen{X: 2, Y: 2}
	thirdWife := Queen{X: 1, Y: 3}
	fourthVeyonce := Queen{X: 2, Y: 4}
	fifthChick := Queen{X: 0, Y: 5}
	sixthGirl := Queen{X: 7, Y: 6}
	seventhBeauty := Queen{X: 2, Y: 7}
	eighthBaby := Queen{X: 4, Y: 0}
	theBoard.AddQueen(firstLady)
	theBoard.AddQueen(secondQueen)
	theBoard.AddQueen(thirdWife)
	theBoard.AddQueen(fourthVeyonce)
	theBoard.AddQueen(fifthChick)
	theBoard.AddQueen(sixthGirl)
	theBoard.AddQueen(seventhBeauty)
	theBoard.AddQueen(eighthBaby)
	if theBoard.Queens[7].X != eighthBaby.X || theBoard.Queens[7].Y != eighthBaby.Y {
		t.Errorf("EigthBaby %v should be the same as Queens placed at index:7. Got %v", eighthBaby, theBoard.Queens[7])
	}

	homeWrecker := Queen{X: 5, Y: 2}
	theBoard.AddQueen(homeWrecker)
	listOfQueens := theBoard.Queens
	if len(listOfQueens) != theBoard.Size {
		t.Errorf("HomeWrecker %v should not be added to the list of Queens.\nLast of the list of Queens %v, expected EightBaby %v", homeWrecker, listOfQueens[len(listOfQueens)-1], eighthBaby)
	}
}

/**
 * TestBoard_DateAllQueen are testing a Board to date all the Queen.
 */
func TestBoard_DateAllQueen(t *testing.T) {
	theBoard := Board{Size: 8}
	firstLady := Queen{X: 0, Y: 0}
	secondQueen := Queen{X: 5, Y: 4}
	thirdWife := Queen{X: 2, Y: 1}

	if !theBoard.AbleToDate(firstLady) {
		t.Errorf("First Lady %v should be able to date.", firstLady)
	}
	theBoard.AddQueen(firstLady)
	if !theBoard.AbleToDate(secondQueen) {
		t.Errorf("Second Queen %v should be able to date", secondQueen)
	}
	theBoard.AddQueen(secondQueen)
	if theBoard.AbleToDate(thirdWife) {
		t.Errorf("Third Wife %v should not be able to date", thirdWife)
	}

}

/**
 * TestBoard_AddQueenTillFull are testing Board to add Queen until its Size (full)
 */
func TestBoard_AddQueenTillFull(t *testing.T) {
	theBoard := Board{Size: 8}
	firstLady := Queen{X: 0, Y: 0}
	secondQueen := Queen{X: 5, Y: 4}
	thirdWife := Queen{X: 2, Y: 1}

	theBoard.AddQueen(firstLady)
	theBoard.AddQueen(secondQueen)
	theBoard.AddQueen(thirdWife)
	theBoard.AddQueen(firstLady)
	theBoard.AddQueen(secondQueen)
	theBoard.AddQueen(thirdWife)
	theBoard.AddQueen(firstLady)
	if theBoard.QueenIsFull() {
		t.Errorf("Queens should not be full yet, Queens size: %d", len(theBoard.Queens))
	}

	theBoard.AddQueen(secondQueen)
	if !theBoard.QueenIsFull() {
		t.Errorf("Queens should be full, Queens size: %d", len(theBoard.Queens))
	}
}

/**
 * TestBoard_QueensAreAllTheSame are testing a Board's Queens compared to another Board that has the same list of Queens (Doppleganger)
 */
func TestBoard_QueensAreAllTheSame(t *testing.T) {

	solutionBoard := Board{Size: 8}

	q01 := Queen{X: 6, Y: 0}
	q02 := Queen{X: 4, Y: 1}
	q03 := Queen{X: 2, Y: 2}
	q04 := Queen{X: 0, Y: 3}
	q05 := Queen{X: 5, Y: 4}
	q06 := Queen{X: 7, Y: 5}
	q07 := Queen{X: 1, Y: 6}
	q08 := Queen{X: 3, Y: 7}
	solutionBoard.AddQueen(q01)
	solutionBoard.AddQueen(q02)
	solutionBoard.AddQueen(q03)
	solutionBoard.AddQueen(q04)
	solutionBoard.AddQueen(q05)
	solutionBoard.AddQueen(q06)
	solutionBoard.AddQueen(q07)
	solutionBoard.AddQueen(q08)

	emptyBoard := Board{Size: 8}

	doppleBoard := Board{Size: 8}

	q1 := Queen{X: 6, Y: 0}
	q2 := Queen{X: 4, Y: 1}
	q3 := Queen{X: 2, Y: 2}
	q4 := Queen{X: 0, Y: 3}
	q5 := Queen{X: 5, Y: 4}
	q6 := Queen{X: 7, Y: 5}
	q7 := Queen{X: 1, Y: 6}
	q8 := Queen{X: 3, Y: 7}
	doppleBoard.AddQueen(q1)
	doppleBoard.AddQueen(q2)
	doppleBoard.AddQueen(q3)
	doppleBoard.AddQueen(q4)
	doppleBoard.AddQueen(q5)
	doppleBoard.AddQueen(q6)
	doppleBoard.AddQueen(q7)
	doppleBoard.AddQueen(q8)

	if solutionBoard.QueensAreAllTheSame(emptyBoard) {
		t.Errorf("Queens at solution board should not be the same as empty board has.")
	}
	if !solutionBoard.QueensAreAllTheSame(doppleBoard) {
		t.Errorf("Queens at solution board should be the same as the dopple board has.")
	}
}
