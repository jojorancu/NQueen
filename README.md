[![Build Status](https://travis-ci.org/jojorancu/NQueen.svg?branch=master)](https://travis-ci.org/jojorancu/NQueen)
# NQueen - Problem

Basically this is a way of N-Queen problem could be solved by using **`DFS-Backtracking Algorithm`** and of course recursively.

Author added mix of humor to the code e.g. `firstLady` until `eigthBaby`is a named variable for testing purposes that points to the **first queen** until the **last queen** to be added (since currently the N is 8 for my testing code).

Another example is `Board` struct has `AbleToDate` function like this:
```go
func (b *Board) AbleToDate(homewrecker Queen) bool {
	for _, queen := range b.Queens {
		if homewrecker.MadnessPotential(queen) {
			return false
		}
	}
	return true
}
```
This function check whether the `homeWrecker` as a `Queen` could be able to fill up the list `Queens` that `Board` has or not.

I know right, you might found another one function that `Queen` has. `MadnessPotential` function contain all the constraints **OR** you could say the ***requirements*** that fulfill to date **`The King`**.
And... ***Ba Dum Tss*** here are the ***requirements*** :
* There should not be Queens on the same **row**.
* And Queens are not allowed on the same **column** too.
* Queens are also not allowed on the same **diagonal**.

*Voila!* Here is the function:
```go
func (q1 *Queen) MadnessPotential(q2 Queen) bool {
	return q1.X == q2.X || q1.Y == q2.Y || math.Abs(float64(q1.X-q2.X)) == math.Abs(float64(q1.Y-q2.Y))
}
```

## Usage
This repository doesn't have any dependency to any library. Create new directory:
```sh
cd /to/your/dev/dir
mkdir nameyouwanted
cd nameyouwanted
```
Using `git clone` to pull the repository:
```sh
git clone https://github.com/jojorancu/NQueen.git
```
There will be **`NQueen`** directory contains the code. 

## Test
Wanted to test ? Change the directory to the **`NQueen`** directory:
```sh
cd NQueen/
```

Use this syntax at the root directory:
```sh
go test -v ./...
```

## Author
[@jojorancu](https://twitter.com/jojorancu "jojorancu / CJSparrow / Jonathan Surya Laksana")
