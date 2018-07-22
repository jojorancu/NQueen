package main

import (
	"math/rand"
	"os"
	"testing"
	"time"

	. "github.com/jojorancu/NQueen/pkg/boardstuff"
)

/**
 * TestMain at main_test.go are used to initialize test cases
 */
func TestMain(m *testing.M) {
	//Setup

	code := m.Run()

	os.Exit(code)
}

/**
 * TestMain_AddSolution are testing for MainBoard to add solution to the slice
 */
func TestMain_AddSolution(t *testing.T) {
	sampleMainBoard := MainBoard{}
	n := 8
	sampleBoard := Board{Size: n}

	q1 := Queen{X: 6, Y: 0}
	q2 := Queen{X: 4, Y: 1}
	q3 := Queen{X: 2, Y: 2}
	q4 := Queen{X: 0, Y: 3}
	q5 := Queen{X: 5, Y: 4}
	q6 := Queen{X: 7, Y: 5}
	q7 := Queen{X: 1, Y: 6}
	q8 := Queen{X: 3, Y: 7}
	sampleBoard.AddQueen(q1)
	sampleBoard.AddQueen(q2)
	sampleBoard.AddQueen(q3)
	sampleBoard.AddQueen(q4)
	sampleBoard.AddQueen(q5)
	sampleBoard.AddQueen(q6)
	sampleBoard.AddQueen(q7)
	sampleBoard.AddQueen(q8)

	if len(sampleMainBoard.Solution) != 0 {
		t.Errorf("SampleMainBoard Solution should be nil or empty. Got Solution length is %d", len(sampleMainBoard.Solution))
	}

	sampleMainBoard.AddSolution(sampleBoard)

	if len(sampleMainBoard.Solution) == 0 {
		t.Errorf("SampleMainBoard Solution should not be nil or empty. Got Solution length is %d", len(sampleMainBoard.Solution))
	} else {
		firstSolution := sampleMainBoard.Solution[0]
		randomIndex := randomNumber(0, n)
		randomQueen := firstSolution.Queens[randomIndex]
		randomDefaultQueen := sampleBoard.Queens[randomIndex]
		if randomQueen.X != randomDefaultQueen.X || randomQueen.Y != randomDefaultQueen.Y {
			t.Errorf("Randomized Queen should be equal by X and Y position.\n Got Random Queen (Solution) %v and Random Default Queen %v", randomQueen, randomDefaultQueen)
		}
	}
}

/**
 * randomNumber is used to return a random number by 3 parameters (2 are input (min,max) and 1 is seamless (timestamp UNIX))
 * @param min is an integer at lowest random number that will generataed
 * @param max is an integer at highest random number that will generated
 * @return random number in range of min and max by time UNIX as the seed
 */
func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

/**
 * TestMain_SolveGame are testing result of SolveTheGame function that should be right at some Board and also the Queens
 */
func TestMain_SolveGame(t *testing.T) {
	testMainBoard := MainBoard{}
	n := 8
	SolveTheGame(n, &testMainBoard)

	if testMainBoard.SumOfPossiblities() == 0 {
		t.Errorf("Solution should be filled and solved. Got length of Solution %d", testMainBoard.SumOfPossiblities())
	}

	doppleBoard := Board{}

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

	if testMainBoard.Solution[87].QueensAreAllTheSame(doppleBoard) && testMainBoard.SumOfPossiblities() != 92 {
		t.Errorf("There are some error at the solution.")
	}
}
