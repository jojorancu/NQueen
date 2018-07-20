package boardstuff

import "math"

/**
 * Queen used to be simulate a Queen's Chess Board game
 * currently Queen only have a position information in linear way (X and Y)
 * X points to column
 * Y points to row
 */
type Queen struct {
	X, Y int
}

/**
 * MadnessPotential is a Queen function to check if the Queen would be madness if its meet another Queen in the same way of direction (constraints)
 * @param q2 is the another Queen that compared
 * @return true if there will be madness with the another Queen, false if they wont meet the constraints
 */
func (q1 *Queen) MadnessPotential(q2 Queen) bool {
	return q1.X == q2.X || q1.Y == q2.Y || math.Abs(float64(q1.X-q2.X)) == math.Abs(float64(q1.Y-q2.Y))
}
