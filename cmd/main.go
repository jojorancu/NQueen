package main

import (
	"fmt"
	"log"
	. "nqueen/pkg/boardstuff"
	"sync"
	"time"
)

/**
 * MainBoard used to store list of Board that have been solved.
 * Solution contain slice of Board that points to ALL possible solution.
 */
type MainBoard struct {
	Solution []Board
}

/**
 * AddSolution is used to add another solution to the MainBoard.Solution.
 * @param solution as a Board to be added
 */
func (m *MainBoard) AddSolution(solution Board) {
	m.Solution = append(m.Solution, solution)
}

/**
 * WaitGroup are used to Wait() ALL the concorrency Go until its Done().
 */
var wg sync.WaitGroup

/**
 * main() function is where the program started (you know, C are using main class right ?).
 */
func main() {
	// Using Concurrency for every N
	start := time.Now()
	theMainBoard := MainBoard{}
	n := 8
	SolveTheGameConcurrently(n, &theMainBoard)
	theMainBoard.PrintAllPossibilities()
	elapsed := time.Since(start)
	log.Printf("Took %s to complete.", elapsed)

	// Using No Concurrency
	// start := time.Now()
	// theMainBoard := MainBoard{}
	// n := 8
	// SolveTheGame(n, &theMainBoard)
	// elapsed := time.Since(start)
	// log.Printf("Took %s to complete.", elapsed)
}

/**
 * PrintAllPossibilities are used to print ALL possibilities not unique
 */
func (m *MainBoard) PrintAllPossibilities() {
	for _, theBoard := range m.Solution {
		fmt.Println(theBoard.Queens)
	}
	fmt.Println(m.SumOfPossiblities())
}

/**
 * SumOfPossibilities are used to sum how much possibilities as amount of solution
 * @return length of MainBoard.Solution as integer
 */
func (m *MainBoard) SumOfPossiblities() int {
	return len(m.Solution)
}

/**
 * SolveTheGame solve without concurrency and print ALL the possiblities
 * @param n represent N-queen as integer
 * @param mainBoard represent all solution Board as MainBoard
 */
func SolveTheGame(n int, mainBoard *MainBoard) {
	for i := 0; i < n; i++ {
		firstLady := Queen{X: i, Y: 0}
		currentBoard := Board{Size: n}
		RecSearchPossibleCheats(firstLady, currentBoard, mainBoard)
	}
}

/**
 * SolveTheGameConcurrently solve with concurrency for every N and print ALL the possiblities
 * @param n represent N-queen as integer
 * @param mainBoard represent all solution Board as MainBoard
 */
func SolveTheGameConcurrently(n int, mainBoard *MainBoard) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			firstLady := Queen{X: i, Y: 0}
			currentBoard := Board{Size: n}
			RecSearchPossibleCheats(firstLady, currentBoard, mainBoard)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/**
 * RecSearchPossibleCheats search recursively for ALL possibilities for cheating the Queens
 * @param theQueen represent current Queen that being checked to be added to the list (theBoard.Queens)
 * @param theBoard represent current Board that contain list (Queens) as solution
 * @param mainBoard represent current MainBoard that contain ALL Board as Solution
 */
func RecSearchPossibleCheats(theQueen Queen, theBoard Board, mainBoard *MainBoard) {
	if theBoard.AbleToDate(theQueen) {
		theBoard.AddQueen(theQueen)
		if theBoard.QueenIsFull() {
			allQueenSet := make([]Queen, theBoard.Size)
			for i, queen := range theBoard.Queens {
				allQueenSet[i] = queen
			}
			solutionBoard := Board{Size: theBoard.Size, Queens: allQueenSet}
			mainBoard.AddSolution(solutionBoard)
		} else {
			for i := 0; i < theBoard.Size; i++ {
				for j := theQueen.Y; j < theBoard.Size; j++ {
					nextQueen := Queen{X: i, Y: j}
					RecSearchPossibleCheats(nextQueen, theBoard, mainBoard)
				}
			}
		}
	}
}
