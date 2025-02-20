# Go Map Access Panic

This repository demonstrates a common error in Go when working with maps: accessing a key that does not exist.  This can lead to runtime panics and unexpected program termination.

The `bug.go` file shows the problematic code.  The `bugSolution.go` file demonstrates how to safely handle this situation.

## How to reproduce

1. Clone the repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go`.

You'll observe a runtime panic because the code attempts to access a non-existent key in an uninitialized map.

## Solution

The `bugSolution.go` file provides a solution by checking if the key exists before accessing it.  Alternative approaches include using the comma ok idiom for concise error handling or using the `golang.org/x/exp/maps` package's safe access functions (if applicable).
