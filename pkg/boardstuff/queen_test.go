package boardstuff_test

import (
	"os"
	"testing"

	. "github.com/jojorancu/NQueen/pkg/boardstuff"
)

func TestMain(m *testing.M) {
	// setup
	// should initialize all before run the test
	// firstLady := Point{x: 0, y: 0}

	code := m.Run()
	// shutdown

	os.Exit(code)
}

func TestQueen_MadnessPotential(t *testing.T) {
	firstLady := Queen{X: 0, Y: 0}
	secondQueen := Queen{X: 5, Y: 4}
	if firstLady.MadnessPotential(secondQueen) {
		t.Errorf("First Lady %v should not have madness potential with the Second Queen %v", firstLady, secondQueen)
	}

	thirdWife := Queen{X: 2, Y: 1}
	if !secondQueen.MadnessPotential(thirdWife) {
		t.Errorf("Second Queen %v should have madness potential with the Second Queen %v", secondQueen, thirdWife)
	}
}
