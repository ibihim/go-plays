package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	left, right *node
	empty       bool
}

func randBool() bool {
	return rand.Intn(2) == 1
}

func createDecreasingProbability(depth int) func() bool {
	// Probability should decrease every 2nd call. So left and right
	// branch have equal probabilities.
	decreaseProbability := false
	// The higher it gets, the lower the probability gets for true
	decreasor := 0

	return func() bool {
		// result needs to be smaller than depth - decreasor.
		// It should be true on the first run and then decrease with every
		// additional tier by 1/depth.
		result := rand.Intn(depth) < depth-decreasor

		// If decrease probability is true, decreasor will be increased
		if decreaseProbability {
			decreasor++
		}
		// Change decrease probability state to opposite.
		decreaseProbability = !decreaseProbability

		return result
	}
}

// createTree creates a tree with a probability of 0.5 for every branch.
func createTree(randBool func() bool) *node {
	trunk := node{}

	// Create a left branch with a probability of 0.5
	if randBool() {
		trunk.left = createTree(randBool)
	}

	// Create a right branch with a probability of 0.5
	if randBool() {
		trunk.right = createTree(randBool)
	}

	return &trunk
}

// printTree prints tree. Indent should start with empty string("")
func printTree(root *node, indent string) {
	fmt.Printf("%s|\n", indent)

	// Prints recursively, if there is a left branch
	if root.left != nil {
		fmt.Printf("%sͰ Left\n", indent)
		printTree(root.left, indent+"\t")
	}

	// Prints recursively, if there is a right branch
	if root.right != nil {
		fmt.Printf("%sͰ Right\n", indent)
		printTree(root.right, indent+"\t")
	}
}
